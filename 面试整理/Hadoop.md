# Hadoop

## HDFS存储模型

* 文件线性切割成block块，偏移量offset的单位是byte；
* block块分散存储在集群的各个节点中；
* 单一文件切分的block块大小是一致的，默认大小为128mb，可以设置block大小；
* block存在副本，副本分散在不同的节点中，默认副本数是3，可以设置副本数（注：副本数不要设置超过节点数）；
* 已上传的文件block副本数可以调整，block大小不能变化；
* hdfs只支持一次写入多次读取，同一时刻只有一个写入者；
* 文件可以append追加数据，hdfs会新增block块存储新数据，并创建副本；
* block副本放置策略：
  * 第一个副本：
    * 集群内提交：
      * 放置在上传文件的datanode中。
    * 集群外提交：
      * 随机挑选一台磁盘不太慢，cpu不太忙的节点。
  * 第二个副本：
    * 放置在与第一个副本不同机架的节点上。
  * 第三个副本：
    * 放置在与第二个副本相同机架的节点上。
  * 更多副本：
    * 随机挑选节点



## HDFS架构模型

**HDFS Client**：

* client与namenode交互元数据信息；
* client与datanode交互文件block数据。

**NameNode（NN）**：

* 保存文件的元数据（如文件大小，时间，block列表，分片位置信息，副本位置信息）；
* 基于内存存储元数据信息，不会和磁盘发生交换；
* 持久化：
  * fsimage：元数据存储到磁盘的文件名为fsimage（内存的快照），fsimage只在集群第一次启动时创建空的文件；
  * editslog：记录了对元数据的操作日志，每隔一段时间与fsimage合并（执行日志记录的操作），生成新的fsimage。

**DataNode（DN）**：

* 保存文件的block数据在磁盘上，同时存储block的元数据文件（MD5校验是否损坏）；
* datanode会向namenode上报心跳数据（3秒一次），提交block列表；
  * 如果namenode10分钟没有收到datanode的心跳，则判定次DN挂掉，从其他DN复制副本到新DN保持副本数。

**SecondaryNameNode（SNN）**：

* 帮助namenode合并fsimage和editslog（避免namenode磁盘IO消费资源）；
* SNN执行合并的时机：
  * 根据配置文件的时间间隔配置项fs.checkpoint.period，默认3600秒；
  * 根据配置文件设置editslog大小配置项fs.checkpoint.size规定，edits文件的最大默认值为64mb。



## HDFS读/写流程

**写流程**：

* client端将要写入的文件进行切分，block大小128mb；
* 与NN交互获取第一个block副本存放的DN列表；
* 将切分后的block再次切分为小文件，小文件大小为64kb；
* client根据从NN获取的DN列表，与其中DN交互，将小文件进行流式传输；
* 第一个DN接收到文件后流式传输副本到下一个DN，以pipeline的方式依次类推直到所有存放副本的DN都将副本写入完毕为止；
* block传输结束后：
  * DN向NN汇报block的信息，NN进行元数据的存储；
  * DN向client汇报写入完成；
  * client向NN汇报写入完成。
* client获取下一个block存放的DN列表，反复执行流式传输，直到文件的block全部写入完毕；
* 最终client汇报完成；
* NN会在写流程更新文件状态。

**读流程**：

* 与NN建立通信，获取一部分block副本的位置列表；
* 线性的从DN获取block，最终合并为一个文件；
* 在block副本列表中按距离择优选择DN。



## HA高可用集群

**HDFS 2.x**

* 解决单点故障：
  * HDFS HA（高可用）：通过主备NN解决，Active NN发生故障，切换到Standby NN。
* 解决内存受限：
  * HDFS Federation（联邦）：水平扩展支持多个NN，所有NN共享DN的存储资源，每个NN分管一部分的目录树结构（保存元数据）。

**HadoopHA 架构：**

* client与NN Active交互元数据，与DN交互block块，但不与NN Standby做交互；
* 所有的DN会同时向两个NN汇报block位置信息；
* NN Active会将元数据写入JournalNode集群（NN之间数据共享），JNN集群过半的节点返回成功消息则代表NN写入成功；
* NN Standby会读取JNN中的元数据，和NN Active保持数据同步；
* 两台NN节点中都存在Zookeeper Failover Controller（ZK的客户端进程）进程，ZKFC进程会与NN和ZK集群两端通信，与NN通信的进程监控NN的健康状态，这两个进程会在ZK集群的目录树结构中争抢创建文件的权利，当某个ZKFC进程成功创建文件，那这个进程管理的NN就是NN Active；
* 当NN Active挂掉，ZKFC进程接收不到心跳，会立即将ZK目录树节点上的文件删除产生事件触发回调，ZKFC Standby进程监听该事件（等待），一但发生事件，ZK将回调ZKFC Standby进程，在ZK集群中创建文件，并将NN Standby提升为NN Active，此时由此节点为client提供服务；
* 当ZKFC Active挂掉，ZK集群的session机制会启动，此时ZKFC Active与ZK集群的socket通信会断开，ZK集群会进行倒计时，计时完毕会产生事件回调，ZKFC Standby创建文件并提升NN Standby为NN Active（两个ZKFC进程还存在隐藏的与对方NN的通信，在提升自己管理的NN为主时会先尝试将对方的NN降级）。



## MapReduce原理

**MapTask**：

* Input Split：原始数据通过split逻辑进行分割；
* Map：多个map按照逻辑对分割后的所有块并行计算，结果会映射成（k，v）格式，并对处理的数据进行分区；
* Buffer In Memory：map处理后的数据会先写入内存缓冲，累加到100MB时会溢写到磁盘；
* Sort：写入磁盘会落地成小文件，小文件内部按照快速排序对相同key的数据分组；
* Merge：所有小文件会通过归并排序合并成一个文件，并行计算的map task阶段会产生多个文件，文件由reduce处理。

**ReduceTask**：

* Merge：reduce task会从多个map task拉取文件，会将一定数量的文件通过归并算法合并；
* Merge：合并后的文件会以归并算法传入reduce的逻辑进行处理；
* Reduce：会按照reduce的逻辑对数据进行处理；
* Output：将计算后的结果输出。



## Yarn资源调度集群

**Yarn集群架构**：

* yarn集群属于主从架构：
  * Resource Manager：管理集群所有的资源
  * NodeManager：管理本节点的资源，任务，并以心跳的方式向RM汇报
  * container：计算框架中的所有角色都由container表示，代表节点的资源单位；
* client提交job后，RM会挑选一台不太忙的节点启动Applocation Master管理当前job的资源调度；
* AM启动完成会回去向RM汇报，由RM决策job任务移动的目标点（container），NM默认启动线程监控container大小，一旦提交的job任务超出了申请资源的额度，会将job杀死；
* 由AM决定job任务的阶段（如MR的map和reduce阶段）提交到哪块container执行，并且job任务的执行情况还会汇报给AM；
* 如果其他的计算框架提交job，RM会在其他节点启动属于该框架的app master，框架之间的资源调度互相隔离。
