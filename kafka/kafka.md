## 1.Kafka 架构
* 概念：kafka是一个高吞吐的分部式消息队列（目前新版官方定义：分布式流平台）
* 消息队列常用场景：系统之间的解耦合，峰值压力缓冲，异步通信。

* Kafka 架构：
    * Producer消息生产者：向消息队列尾部生产数据；
    * Consumer消息消费者：从消息队列头部消费数据；
    * Broker中间人：Kafka集群的Server，每个节点称作一个broker，负责存储数据并处理数据的读写请求；
    * Topic主题：存在多个partition消息队列分布式存在集群中。

* Kafka的消息存储和生产消费模型：
    * 一个partition只对应一个broker，一个broker可以对应多个partition，比如，topic有6个partition，集群存在两个broker，那每个broker就管理3个partition；
    * 消息直接写入磁盘文件，并不是存储在内存中；
    * 根据时间策略（默认一周）删除队列中的消息，而不是消费完就删除，在kafka里面没有消费完这么个概念，只有过期这一概念；
    * 生产者自己决定往哪个partition写消息，可以是轮询的负载均衡，也可以是基于hash的partition策略；
    * kafka里面的消息是由topic来组织的，一个队列就是一个topic，为了并行每个topic又分为很多个partition，数据在每个partition里面是强有序的，相当于有序的队列，其中每个消息都有个offset序号，队列从前面消费后面生产；
    * partition可以想象为一个文件，当新生产数据就会对partition做append追加；
    * 消费者自己维护消费数据的offset序号；
    * 每个消费者都有对应的group，group中是队列的消费模型，各个消费者消费不同的partition，一个消息在group内只消费一次，各个group各自独立消费，互不影响。

* Kafka 特点：
    * FIFO的生产者消费者模型：partition内部是FIFO模式，partition之间不是严格的FIFO模式，如果topic只存在一个partition，那么就是严格的FIFO；
    * 高性能：单节点支持上千客户端；
    * 持久性：直接append到磁盘里面去，这样的好处是直接持久化，数据丢失，第二个好处是顺序的写，消费数据也是顺序的读，所以持久化的同时还能保证顺序读写；
    * 分布式：分布式的数据副本，就是同一份数据可以存到多个不同的broker上面去，当一份磁盘数据坏掉的时候，数据存在副本不会丢失；
    * 灵活：消息长时间持久化+Client维护消费状态。

## 2.Kafka 集群搭建
```bash
node01, node02, node03：
$> zkServer.sh start

node01：
$> cd ~
$> tar -zxvf ./kafka_2.10-0.8.2.2.tgz -C /opt/sxt/
$> vi /etc/profile
export KAFKA_HOME=/opt/sxt/kafka_2.10-0.8.2.2
export PATH=$PATH:$KAFKA_HOME/bin
$> source /etc/profile
$> cd /opt/sxt/kafka_2.10-0.8.2.2/config
$> vi server.properties 
broker.id=0 # 配置集群server的id
log.dirs=/data/kafka # kafka集群文件存储路径
zookeeper.connect=node02:2181,node03:2181,node04:2181 # 配置zookeeper集群的信息
$> cd ../../
$> scp -r /opt/sxt/kafka_2.10-0.8.2.2 node03:`pwd`
$> scp -r /opt/sxt/kafka_2.10-0.8.2.2 node04:`pwd`

node02：
$> vi /etc/profile
export KAFKA_HOME=/opt/sxt/kafka_2.10-0.8.2.2
export PATH=$PATH:$KAFKA_HOME/bin
$> source /etc/profile
$> cd /opt/sxt/kafka-0.8.2.2/config
$> vi server.properties
broker.id=1

node03：
$> vi /etc/profile
export KAFKA_HOME=/opt/sxt/kafka_2.10-0.8.2.2
export PATH=$PATH:$KAFKA_HOME/bin
$> source /etc/profile
$> cd /opt/sxt/kafka-0.8.2.2/config
$> vi server.properties
broker.id=2

node01, node02, node03：
$> kafka-server-start.sh -daemon /opt/sxt/kafka_2.10-0.8.2.2/config/server.properties
```

## 3.Kafka 基本操作
* 创建消息队列：
```bash
node01：
$> kafka-topics.sh --create --zookeeper node02:2181,node03:2181,node04:2181 --topic 20180417 --partitions 3 --replication-factor 3 
# create表示创建，zookeeper设置zk集群信息，topic指定会话名称，partition指定分区名称，replication-factor指定副本个数
```
* 查看所有消息队列：
```bash
node01：
$> kafka-topics.sh --list --zookeeper node02:2181,node03:2181,node04:2181
```
* 往消息队列中生产数据：
```bash
node02：
$> kafka-console-producer.sh --topic 20180417 --broker-list node02:9092,node03:9092,node04:9092
```
* 从消息队列中消费数据：
```bash
node03：
# 从消费者建立后开始消费（基于zookeeper保存的消费偏移量）
$> kafka-console-consumer.sh --topic 20180417 --zookeeper node02:2181,node03:2181,node04:2181 
# 添加权限，从头开始进行消费
$> kafka-console-consumer.sh --topic 20180417 --zookeeper node02:2181,node03:2181,node04:2181 --from-beginning 
```
* 通过linux管道命令将文本中的数据批量生产：
```bash
node02：
$> cat ~/NASA_access_log_Aug95 | kafka-console-producer.sh --topic 20180130 --broker-list node01:9092,node02:9092,node03:9092
```
* 查看消息队列的详细信息：
```bash
node01：
$> kafka-topics.sh --describe --zookeeper node02:2181,node03:2181,node04:2181 --topic 20180417
```