# 1.Centos7 安装 Docker

* 安装 gcc 依赖

  ```bash
  > yum -y install gcc
  > yum -y install gcc-c++
  ```

* 删除旧版本的 Docker

  ```bash
  > yum remove docker \
          docker-client \
          docker-client-latest \
          docker-common \
          docker-latest \
          docker-latest-logrotate \
          docker-logrotate \
          docker-selinux \
          docker-engine-selinux \
          docker-engine
  ```

* 安装所需的软件包

  ```bash
  > yum install -y yum-utils device-mapper-persistent-data lvm2
  ```

* 更新 yum 源

  ```bash
  > yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
  > yum makecache fast
  ```

* 安装 Docker-ce

  ```bash
  > yum -y install docker-ce
  > systemctl start docker
  ```

* 验证安装

  ```bash
  > docker version
  > docker info
  ```

* 添加阿里云镜像加速

  ```bash
  > mkdir -p /etc/docker
  > vim /etc/docker/daemon.json
  {
      "registry-mirrors": ["https://vob2wv9t.mirror.aliyuncs.com"]
  }
  ```

* 启动，重启和查看docker进程

  ```bash
  > systemctl daemon-reload
  > systemctl restart docker
  > ps -ef | grep docker*
  ```



# 2.安装常见应用示例

## 2.1 安装 MySQL

```bash
> docker pull mysql:5.7
> docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -d mysql:5.7
> docker exec -it mysql env LANG=C.UTF-8 /bin/bash
    > mysql -h localhost -u root -p
    > Enter password: 123456
    mysql> select host,user,plugin,authentication_string from mysql.user;
    mysql> use mysql;
    mysql> alter user 'root'@'%' identified with mysql_native_password by '123456';
    mysql> flush privileges;    
    mysql> select host,user,plugin,authentication_string from mysql.user;
```

## 2.2 安装 Redis

```bash
> docker pull redis
> docker run --name redis -p 6379:6379 -v $PWD/data:/data -d redis redis-server --appendonly yes
> docker exec -it redis redis-cli
```

## 2.3 安装 ElasticSearch，Kibana

* 安装elasticsearch

  ```bash
  > docker pull elasticsearch:6.4.0
  > docker run --name elasticsearch -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" -d elasticsearch:6.4.0
  > curl http://192.168.121.100:9200
  ```

* 解决跨域问题

  ```bash
  > docker exec -it elasticsearch /bin/bash
      > cd /usr/share/elasticsearch/config
      > vi elasticsearch.yml
          - http.cors.enabled: true
          - http.cors.allow-origin: "*"
      > exit
  > docker restart elasticsearch
  ```

* 安装ik分词器

  ```bash
  > docker exec -it elasticsearch /bin/bash
      > cd /usr/share/elasticsearch/plugins
      > elasticsearch-plugin install https://github.com/medcl/elasticsearch-analysis-ik/releases/download/v6.4.0/elasticsearch-analysis-ik-6.4.0.zip
      > exit
  > docker restart elasticsearch
  ```

* 安装kibana

  ```bash
  > docker pull kibana
  > docker run --name kibana5.6.11 --link=elasticsearch  -p 5601:5601 -d kibana
  > curl http://192.168.121.100:5601
  ```

## 2.4 安装 RabbitMQ

```bash
> docker pull rabbitmq:3.7.7-management
> docker run --name rabbitmq -p 5672:5672 -p 15672:15672 \
    -v $PWD/data:/var/lib/rabbitmq --hostname myRabbit \ 
    -e "RABBITMQ_DEFAULT_VHOST=/leyou" \
    -e "RABBITMQ_DEFAULT_USER=leyou" \ 
    -e "RABBITMQ_DEFAULT_PASS=leyou" \
    -d rabbitmq:3.4.4-management

> http://192.168.121.100:15672
    - username: guest 
    - password: guest
```

## 2.5 安装 kafka，zookeeper

```bash
> docker pull wurstmeister/zookeeper
> docker pull wurstmeister/kafka:2.11-0.11.0.3

> docker run --name zookeeper -p 2181:2181 -t -d wurstmeister/zookeeper
> docker run --name kafka --publish 9092:9092 --link zookeeper \
    --env KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 \
    --env KAFKA_ADVERTISED_HOST_NAME=192.168.121.100 \
    --env KAFKA_ADVERTISED_PORT=9092 \
    --volume /etc/localtime:/etc/localtime 
    -d wurstmeister/kafka:latest
    
> docker exec -it kafka /bin/bash
    > cd /opt/kafka_2.11-0.11.0.3/bin/
    > ./kafka-topics.sh --create --zookeeper zookeeper:2181 --replication-factor 1 --partitions 8 --topic test
    > ./kafka-console-producer.sh --broker-list 192.168.121.100:9092 --topic test

> docker exec -it kafka /bin/bash
    > cd /opt/kafka_2.11-0.11.0.3/bin/
    > ./kafka-console-consumer.sh --bootstrap-server 192.168.121.100:9092 --topic test --from-beginning
```

## 2.6 安装 alibaba nacos

```bash
> docker pull nacos/nacos-server
> docker run --env MODE=standalone --name nacos -d -p 8848:8848 nacos/nacos-server
```



# 3.开启认证的远程端口2376

```bash
cd ~
mkdir ./ca

#生成ca私钥(使用aes256加密),输入两次密码
openssl genrsa -aes256 -out ca-key.pem 4096
#生成ca证书，填写配置信息,然后依次输入国家是 CN，省例如是Shanghai、市Shanghai、组织名称、组织单位、姓名或服务器名、邮件地址
openssl req -new -x509 -days 365 -key ca-key.pem -sha256 -out ca.pem

#生成server证书私钥文件
openssl genrsa -out server-key.pem 4096
#生成server证书请求文件
openssl req -subj "/CN=192.168.121.100" -sha256 -new -key server-key.pem -out server.csr

# 配置白名单，多个用逗号隔开，例如： IP:192.168.1.111,IP:0.0.0.0，这里需要注意，虽然0.0.0.0可以匹配任意，但是仍然需要配置你的服务器外网ip
echo subjectAltName = IP:192.168.121.100,IP:0.0.0.0 >> extfile.cnf
#把 extendedKeyUsage = serverAuth 键值设置到extfile.cnf文件里，限制扩展只能用在服务器认证
echo extendedKeyUsage = serverAuth >> extfile.cnf

#使用CA证书及CA密钥以及上面的server证书请求文件进行签发，生成server自签证书
openssl x509 -req -days 365 -sha256 -in server.csr -CA ca.pem -CAkey ca-key.pem \-CAcreateserial -out server-cert.pem -extfile extfile.cnf

#生成client证书RSA私钥文件
openssl genrsa -out key.pem 4096
#生成client证书请求文件
openssl req -subj '/CN=client' -new -key key.pem -out client.csr

#继续设置证书扩展属性
echo extendedKeyUsage = clientAuth >> extfile.cnf

#生成client自签证书（根据上面的client私钥文件、client证书请求文件生成）
openssl x509 -req -days 365 -sha256 -in client.csr -CA ca.pem -CAkey ca-key.pem \-CAcreateserial -out cert.pem -extfile extfile.cnf

#删除生成的临时文件
rm -rf client.csr server.csr

#修改证书为只读权限保证证书安全
chmod -v 0400 ca-key.pem key.pem server-key.pem
chmod -v 0444 ca.pem server-cert.pem cert.pem

#复制服务端需要用到的证书到docker配置目录下便于识别使用
cp server-cert.pem ca.pem server-key.pem /etc/docker/

#修改docker配置
vim /lib/systemd/system/docker.service
ExecStart=/usr/bin/dockerd-current \
	--tlsverify \
	--tlscacert=/etc/docker/ca.pem \
	--tlscert=/etc/docker/server-cert.pem \
	--tlskey=/etc/docker/server-key.pem \
	-H tcp://0.0.0.0:2376 \
	-H unix:///var/run/docker.sock \

# 开放防火墙的2376的端口
firewall-cmd --zone=public --add-port=2376/tcp --permanent
firewall-cmd --reload

#重载服务并重启docker
systemctl daemon-reload && systemctl restart docker

#查看是否存在2376端口
yum install net-tools
netstat -tunlp

#保存证书客户端文件到本地,这里用的是sz命令，ftp也可以只要能放到本地客户端即可
cd /etc/docker
sz ca.pem cert.pem key.pem

#测试一下证书是否配置成功，如果成功，会输出证书相关信息，如果有fail，请检查证书
docker --tlsverify --tlscacert=ca.pem --tlscert=cert.pem --tlskey=key.pem -H=192.168.121.100:2376 version
```

```bash
#创建 Docker TLS 证书
#!/bin/bash

#相关配置信息
SERVER="192.168.33.76"
PASSWORD="pass123456"
COUNTRY="CN"
STATE="广州省"
CITY="广州市"
ORGANIZATION="公司名称"
ORGANIZATIONAL_UNIT="Dev"
EMAIL="492376344@qq.com"

###开始生成文件###
echo "开始生成文件"

#切换到生产密钥的目录
cd /etc/docker   
#生成ca私钥(使用aes256加密)
openssl genrsa -aes256 -passout pass:$PASSWORD  -out ca-key.pem 2048
#生成ca证书，填写配置信息
openssl req -new -x509 -passin "pass:$PASSWORD" -days 3650 -key ca-key.pem -sha256 -out ca.pem -subj "/C=$COUNTRY/ST=$STATE/L=$CITY/O=$ORGANIZATION/OU=$ORGANIZATIONAL_UNIT/CN=$SERVER/emailAddress=$EMAIL"

#生成server证书私钥文件
openssl genrsa -out server-key.pem 2048
#生成server证书请求文件
openssl req -subj "/CN=$SERVER" -new -key server-key.pem -out server.csr
#使用CA证书及CA密钥以及上面的server证书请求文件进行签发，生成server自签证书
openssl x509 -req -days 3650 -in server.csr -CA ca.pem -CAkey ca-key.pem -passin "pass:$PASSWORD" -CAcreateserial  -out server-cert.pem

#生成client证书RSA私钥文件
openssl genrsa -out key.pem 2048
#生成client证书请求文件
openssl req -subj '/CN=client' -new -key key.pem -out client.csr

sh -c 'echo "extendedKeyUsage=clientAuth" > extfile.cnf'
#生成client自签证书（根据上面的client私钥文件、client证书请求文件生成）
openssl x509 -req -days 3650 -in client.csr -CA ca.pem -CAkey ca-key.pem  -passin "pass:$PASSWORD" -CAcreateserial -out cert.pem  -extfile extfile.cnf

#更改密钥权限
chmod 0400 ca-key.pem key.pem server-key.pem
#更改密钥权限
chmod 0444 ca.pem server-cert.pem cert.pem
#删除无用文件
rm client.csr server.csr

echo "生成文件完成"
###生成结束###
```

