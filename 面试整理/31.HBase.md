# HBase

## 数据模型

| Row Key  | Time Stamp | CF1          | CF2         | CF3         |
| -------- | ---------- | ------------ | ----------- | ----------- |
|          | T6         |              | CF2:q1=val1 | CF3:q3=val3 |
| 11248112 | T3         |              |             |             |
|          | T2         | CF1：q2=val2 |             |             |

**Row Key（行键）**：

* 决定一行数据，相当于主键；
* 写数据时按照字典顺序插入（ASCII排序）；
* 行键只能存储64k的字节数据（越短越提高检索性能）。

**Time Stamp（时间戳）**：

* 列数据的版本号，当对某一列提交新数据时hbase表通过添加数据并标记版本实现update；
* 每个列族都可以设置maxversion，表示版本的最大有效数。

**Column Family（列族）& qualifier（列）**：

* HBase表中的每列都归属于列族（列族必须在表创建时预先定义）；
* 列族存在多个列成员，列族名作为该列族所有列名的前缀，列可以动态添加；
* HBase将列族数据存储在同一目录下，分多个文件保存。

**Cell（单元格）**：

* 由rowkey与列族：列交叉决定；
* 单元格表示列数据，存在版本；
* 内容是未解析的字节数组（字节码）；
* 由 {row key，column =（<family> + <qualifier>），version} 唯一决定。



## 架构模型

**Client（客户端）**：

* 访问HBase的接口；
* 维护cache加快对hbase的访问。

**Zookeeper（分布式协同）**：

* 保证集群中只存在一个HMaster主节点，实现HA（高可用）；
* 监控Region Server的健康状态，出现宕机等情况会实时通知HMaster进行数据迁移；
* 存储所有Region的寻址入口；
* 存储HBase表的元数据信息。

**HMaster（主节点）**：

* 为Region Server从节点分配Region；
* 对Region Server做负载均衡；
* 重新分配宕机的Region Server上的Region；
* 管理用户对表的增删改查。

**HRegion Server（从节点）**：

* 维护Region，处理对Region的IO请求；
* 负责切分在运行过程中达到阈值的Region（等分原则）。

**HRegion（数据区域）**：

* 一段连续的表数据存储区域（Row Key会顺序排列）；
* Region中的数据达到某个阈值就会进行水平拆分（同一行的数据一定会存在同一个Region中）。

**Store（列族）**：

* 多个Store组成Region，1个Store对应1个列族；
* 由1个MemStore组成和0至多个StoreFile组成。

**MemStore（写缓存）**：

* MemStore是Client提交操作进行后Store先写入内存的缓存数据（1个）。

**StoreFile（持久化）**：

* StoreFile是MemStore达到阈值溢写到磁盘（Linux文件系统 or HDFS）的小文件（0或多个）；
* StoreFile的数量到达阈值时系统会进行合并（minor小范围合并，major大范围合并）；
* 当一个Region中的所有StoreFile大小数量达到阈值时，会拆分当前的Region，并由HMaster迁移到相应的从节点；
* Client检索数据会先在MemStore中找，找不到再在StoreFile中找；
* Store以HFile的格式保存在HDFS中。

**HLog（日志文件）**：

* 存储Client提交数据的动作和数据。
