[TOC]

# 1.Zookeeper 集群概念：
* 集群的特点：
    - 原子性：准确的反馈成功或失败，没有中间状态。
    - 最终一致性：
        - 每个server都有统一的数据视图，client的写请求总会由leade处理（如果通过follower提交的请求会转发到leader），leader会先在集群中通信确认接下来的操作是否可以执行，等到半数以上的server都响应后再进行对所有server的数据同步；
        - 如果发生网络延迟丢包的问题没有响应leader导致节点数据没有更新，这时client如果向没更新的server请求数据的话，该server会向更新过的server同步数据，完毕后再响应client；
        - 集群数据同步阶段，leader会在本地内存中为每个server准备消息队列，各个server同步出队的数据，在集群中过半的server完成同步后集群确认同步完成，出现网络延迟或其他问题未同步完成的server在接下来也能够继续一个个的同步出队的数据从而达到集群数据的最终一致性。
    - 高可用性：leader如果发生故障，zk集群会通过选举机制重新选出leader。
    - 防网络分区（脑裂）：在client提交请求时如果集群中存在节点宕机的情况，那么会采取过半原则，如果集群半数以上的节点都能响应client的请求，那么则表示集群响应了请求。

* 集群状态：
    * 选举模式：
        * 无主模型选举leader的状态。
    * 广播模式：
        * leader向所有server同步数据的状态。
* Server状态：
    * leading：当前server为被选举出来的leader；
    * looking：当前server不知道leader是谁，正处于搜寻状态；
    * following：leader已经选举出来，当前server与之同步。
* 主从分工：
    * 领导者（leader）：负责处理clinet提交的增删改的请求，负责进行投票的发起和决议；
    * 学习者（learner）：
        * 跟随者（follower）用于接受客户端查询的请求并响应结果，如果是写请求则会转发请求到leader，也在选举主节点过程中参与投票；
        * 观察者（observer）接受客户端连接，将写请求转发给leader，但observer不参加投票过程，只同步leader的状态，observer的目的是为了扩展系统，提高读取速度。
    * 客户端（client）：请求发起方。
* Session会话：
    * clinet与集群server建立TCP连接后，leader会在client操作的节点上创建session，同时该session会被放到leader的消息队列中，让所有server都同步；
    * 如果clinet连接的Server出现问题，session会经过一段时间自动销毁，如果没有超过timeout，client可以连接其他server，其他server也会根据session获知client的状态。
* Znode数据模型：
    * 目录结构：层次的，目录型结构，便于管理逻辑关系。
    * Znode信息：包含最大1MB的数据信息，记录了Zxid等元数据信息。
    * 短暂模式（ephemeral）：短暂znode的客户端会话结束时，zookeeper会将该短暂znode删除，短暂znode不可以有子节点。
    * 持久模式（persistent）：持久znode不依赖于客户端会话，只有当客户端明确要删除该持久znode时才会被删除。
    * 序列化处理：如果存在同名的znode同时要求创建，序列化处理后保证都能创建成功并保持唯一性。
* Watcher事件监听机制：
    * Watcher是ZooKeeper的核心功能，Watcher可以监控目录节点以及子目录的数据变化，一旦这些状态发生变化，服务器就会通知所有设置在这个目录节点上的Watcher，从而每个客户端都能很快知道它所关注的目录节点状态是否发生变化，从而做出相应的反应。
    * 可以设置观察的操作：exists，getChildren，getData；
    * 可以触发观察的操作：create，delete，setData。

# 2.搭建 Zookeeper 集群：
```
master01：
[root@master01 ~]# tar -zxvf zookeeper-3.4.13.tar.gz -C /usr/local
[root@master01 ~]# vim /etc/profile
```
```
[root@master01 ~]# source /etc/profile
[root@master01 ~]# scp /etc/profile root@slave01:/etc/
profile                                                                  100% 2021     1.4MB/s   00:00    
[root@master01 ~]# scp /etc/profile root@slave02:/etc/
profile                                                                  100% 2021     1.3MB/s   00:00    
[root@master01 ~]# cd /usr/local/zookeeper-3.4.13/conf/
[root@master01 conf]# cp zoo_sample.cfg zoo.cfg
[root@master01 conf]# vim zoo.cfg
dataDir=/usr/local/zookeeper-3.4.13/zk
server.1=master01:2888:3888
server.2=slave01:2888:3888
server.3=slave02:2888:3888
```
```
[root@master01 conf]# mkdir -p /usr/local/zookeeper-3.4.13/zk
[root@master01 conf]# echo 1 > /usr/local/zookeeper-3.4.13/zk/myid
[root@master01 conf]# cd ../../
[root@master01 local]# scp -r zookeeper-3.4.13/ root@slave01:`pwd`
[root@master01 local]# scp -r zookeeper-3.4.13/ root@slave02:`pwd`
slave01：
[root@slave01 ~]# source /etc/profile
[root@slave01 ~]# echo 2 > /usr/local/zookeeper-3.4.13/zk/myid 
slave02：
[root@slave02 ~]# source /etc/profile
[root@slave02 ~]# echo 3 > /usr/local/zookeeper-3.4.13/zk/myid
master01，slave01，slave02：
#> zkServer.sh start
ZooKeeper JMX enabled by default
Using config: /usr/local/zookeeper-3.4.13/bin/../conf/zoo.cfg
Starting zookeeper ... STARTED
#> jps
1920 QuorumPeerMain
```