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

