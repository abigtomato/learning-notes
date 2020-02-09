[TOC]

# 1.Redis单节点安装
* 编译安装redis源码包
```bash
$> tar -zxvf redis-2.8.18.tar.gz
$> yum -y install gcc tcl -y      # 安装编译环境
$> make && make PREFIX=/opt/sxt/redis install  # make命令编译，&&判断是否编译成功，如果成功在PREFIX目录下安装
```
* 添加redis环境变量
```bash
$> vim /etc/profile
    export REDIS_HOME=/opt/sxt/redis
    export PATH=$PATH:$REDIS_HOME/bin
$> source /etc/profile
```
* 安装redis服务
```bash
$> cd cd ~/redis-2.8.18/utils
$> ./install_server.sh
（之后的提示是设置服务占用的端口号，文件的路径，持久化数据的目录，redis的可执行路径，全部默认即可）
```
* 启动redis客户端
```bash
$> redis-cli
```

# 2.Redis数据模型

## 2.1 redis基本操作命令
![b900679c50912a2ca645a58fc6dd2ed2.png](en-resource://database/708:1)

```bash
127.0.0.1:6379> select database     # 选择要操作的库
127.0.0.1:6379> help                # 查看帮助
127.0.0.1:6379> help @<group>
127.0.0.1:6379> ttl key             # 查看过期时间
127.0.0.1:6379> type key            # 查看值的类型
127.0.0.1:6379> object encoding key # 通过编码查看类型
127.0.0.1:6379> keys 模式匹配   # 通过模式匹配查看所有key
    * 任意长度字符
    ? 任意一个字符
    []集合中的任意一个
127.0.0.1:6379> flushall            # 清空库
```

## 2.2 二进制安全的字符串操作

### 2.2.1 基本操作
* 以格式化的方式启动客户端
```bash
127.0.0.1:6379> redis-cli --raw
set key value [EX seconds] [PX milliseconds] [NX|XX]    # 新建字符串
    EX: 设置过期时间，秒 
    PX: 设置过期时间，毫秒
    NX: 键不存在时才能设置 (只能新建)
    XX: 键存在时才能设置 (只能修改)
```
* 设置多个字符串的值
```BASH
127.0.0.1:6379> mset key value [key value ...]
```
* 设置单个字符串的值（key不存在时才能设置成功）
```BASH
127.0.0.1:6379> setnx key value
```
* 设置多个字符串的值（key都不能存在时才能设置成功）
```BASH
127.0.0.1:6379> msetnx key value [key value ...]
```
* 通过键取字符串
```BASH
127.0.0.1:6379> get key
```
* 通过多个键获取对应的字符串
```BASH
127.0.0.1:6379> mget key [key ...]
```
* 返回旧串并设置新串
```BASH
127.0.0.1:6379> getset key value
```
* 查看字符串长度
```BASH
127.0.0.1:6379> strlen key
```
* 追加字符串，key存在就追加，否则新建
```BASH
127.0.0.1:6379> append key value
```
* 获取子串，start和end表示索引区间 (正向索引从0开始，反向索引从-1开始)
```BASH
127.0.0.1:6379> getrange key start end
```
* 覆盖字符串，offset表示偏移量，value表示新覆盖的字符串
```BASH
127.0.0.1:6379> setrange key offset value
```

### 2.2.2 数值计算
* 步长增加1（字符串会被转换为64位有符号整型操作，结果依旧转换为字符串）
```BASH
127.0.0.1:6379> incr key
```
* 步长减少1
```BASH
127.0.0.1:6379> decr key
```
* 步长增加指定数值，decrement表示增加的数值
```BASH
127.0.0.1:6379> incrby key decrement
```
* 步长减少指定数值，decrement表示减少的数值
```BASH
127.0.0.1:6379> decrby key decrement
```
* 步长增加指定浮点数值
```BASH
127.0.0.1:6379> incrbyfloat key decrement
```
* 步长减少指定浮点数值
```BASH
127.0.0.1:6379> decrbyfloat key decrement
```

### 2.2.3 位图bitmap
* 设置底层二进制位上某一位的值，offset表示二进制位的偏移量
```BASH
127.0.0.1:6379> setbit key offset value(0/1)
```
* 获取二进制位上某一位上的值
```BASH
127.0.0.1:6379> getbit key offset
```
* 返回二进制值在指定字节区间上第一次出现的二进制位偏移量，start end表示字节数组的区间，返回的偏移量是整个底层存储二进制位的偏移量
```BASH
127.0.0.1:6379> bitpos key bit [start] [end]
```
* 统计指定字节区间上值为1的个数，start end表示字节数组的区间
```BASH
127.0.0.1:6379> bitcount key [start] [end]
```
* 对一个或多个字节的二进制位求逻辑且，结果保存在destkey中
```BASH
127.0.0.1:6379> bitop and destkey key [key...]
```
* 对一个或多个字节的二进制位求逻辑或
```BASH
127.0.0.1:6379> bitop or destkey key [key...]
```
* 对一个或多个字节的二进制位求逻辑异或
```BASH
127.0.0.1:6379> bitop xor destkey key [key...]
```
* 对一个或多个字节的二进制位求逻辑非
```BASH
127.0.0.1:6379> bitop not destkey key [key...]
```

### 2.2.4 位图的应用示例
* 统计网站在某一时间段内用户的上线天数：
    * 用户ID为key，天作为offset，一年的天数置为365个二进制位(0-364)，用户在某天上线则将某天对应的二进制位置为1，这样存储大约50byte左右。
    * 要统计一段时间内上线天数只要使用 bitcount userID 0 364 命令统计1出现的次数即可。
* 统计某几天网站的上线用户：
    * 天数作为key，网站注册的所有用户数置为二进制位，每个用户对应一位，在当天上线将相应的位置置为1。
    * 要统计指定天数的活跃用户只需要先通过 bitop or result 20180401 20180402 20180403 将几天的二进制位 (也就是上线情况)进行按位或(取出为1的)操作写入结果key-value，然后通过bitcount result 0 -1统计1出现的次数即可。
```BASH
127.0.0.1:6379> setbit 20180401 8 1        # 表示第8号用户在20180401这天登录
127.0.0.1:6379> setbit 20180402 18 1
127.0.0.1:6379> setbit 20180403 24 1
127.0.0.1:6379> bitop or result 20180401 20180402 20180403
127.0.0.1:6379> bitcount result 0 -1
```

## 2.3 有序可重复的List列表

### 2.3.1 栈，双端队列
* 从左边向列表压入数据
```BASH
127.0.0.1:6379> lpush key value [value ...]
```
* 从右边向列表压入数据
```BASH
127.0.0.1:6379> rpush key value [value ...]
```
* 从列表左边弹出数据
```BASH
127.0.0.1:6379> lpop key
```
* 从列表右边弹出数据
```BASH
127.0.0.1:6379> rpop key
```
* 可做为双端队列使用
```BASH
127.0.0.1:6379> lpush rpop|rpush lpop
```
* 可做为栈使用
```BASH
127.0.0.1:6379> lpush lpop|rpush rpop
```
* 从列表右边弹出元素压入另一个列表左边
```BASH
127.0.0.1:6379> rpoplpush source destination
```

### 2.3.2 数组
* 获取指定下标范围元素
```BASH
127.0.0.1:6379> lrange key start end
```
* 获取指定下标元素
```BASH
127.0.0.1:6379> lindex key index
```
* 设置指定下标元素
```BASH
127.0.0.1:6379> lset key index value
```
* 返回列表长度
```BASH
127.0.0.1:6379> llen key
```
* count > 0从表头开始搜索，删除count个与value相同的元素，反之从表尾搜索
```BASH
127.0.0.1:6379> lrem key count value
```
* 去除指定范围外的元素
```BASH
127.0.0.1:6379> ltrim key start end
```

### 2.3.3 链表
* 在列表中某个存在的值前后插入新数据，pivot表示存在的数据，value表示新数据
```BASH
127.0.0.1:6379> linsert key before|after pivot value
```

### 2.3.4 阻塞的消息队列
* 阻塞弹出操作，timeout为超时时间，为0表示永久阻塞，直到有数据可以弹出
```BASH
127.0.0.1:6379> blpop key [key...] timeout
```
* 模拟消费者，直到生产者压入数据才会弹出
```BASH
127.0.0.1:6379> brpop key [key...] timeout
```
* 模拟生产者
```BASH
127.0.0.1:6379> lpush key value
```
* 从列表右边阻塞弹出元素压入到另一个列表的左边
```BASH
127.0.0.1:6379> brpoplpush source destination timeout
```

## 2.4 Hash散列

### 2.4.1 基本操作
* 插入单个字段，field和value是内部键值对
```BASH
127.0.0.1:6379> hset key field value
```
* 插入多个字段
```BASH
127.0.0.1:6379> hmset key field value [field valie ...]
```
* 返回字段个数
```BASH
127.0.0.1:6379> hlen key
```
* 判断字段是否存在
```BASH
127.0.0.1:6379> hexists key field
```
* 返回字段值
```BASH
127.0.0.1:6379> hget key field
```
* 返回多个字段值
```BASH
127.0.0.1:6379> hmget key field [field ...]
```
* 返回所有建值对
```BASH
127.0.0.1:6379> hgetall key
```
* 返回所有字段名
```BASH
127.0.0.1:6379> hkeys key
```
* 返回所有值
```BASH
127.0.0.1:6379> hvals key
```
* 在字段对应得值上进行整数的增量计算
```BASH
127.0.0.1:6379> hincrby key field increment
```
* 删除指定字段
```BASH
127.0.0.1:6379> hdel key field [field ...]
```

### 2.4.2 散列的应用示例
* 微博好友关注
    * 用户ID为key，field为所有好友ID，value为对应关注时间
* 用户维度统计
    * 用户ID为key，不同维度为field，value为对应得统计数

## 2.5 无序不重复的Set集合

### 2.5.1 基本操作
* 添加一个或多个元素
```BASH
127.0.0.1:6379> sadd key member [member ...]
```
* 返回集合包含的所有元素
```BASH
127.0.0.1:6379> smembers key
```
* 检查给定元素是否存在于集合中
```BASH
127.0.0.1:6379> sismember key member
```
* 返回集合中元素的个数
```BASH
127.0.0.1:6379> scard key
```
* 将元素重原集合destination移动到目标集合member
```BASH
127.0.0.1:6379> smove source destination member
```

### 2.5.2 随机抽取
* count为整数且小于集合基数，返回包含count个随机元素的集合；如果count大于等于整个集合基数，返回整个集合；count如果为负数，返回一个count绝对值长度的有重复元素的数组；count为0，返回空；count不指定，随机返回一个元素
```BASH
127.0.0.1:6379> srandmember key [count]
```
* 随机从集合中移除并返回移除的元素
```BASH
127.0.0.1:6379> spop key
```

### 2.5.3 交并差
求差集(从第一个key的集合中去除其他集合和自己的交集部分)
```bash
127.0.0.1:6379> sdiff key [key ...]
```
将差集的结果存储到目标destination中
```bash
127.0.0.1:6379> sdiffstore destination key [key ...]
```
取所有集合并集：
```bash
127.0.0.1:6379> sunion key [key ...]
```
将并集的结果存储到目标destination中
```bash
127.0.0.1:6379> sunionstore destination key [key ...]
```
取集合的交集：
```bash
127.0.0.1:6379> sinter key [key ...]
```
将交集的结果存储到目标destination中
```bash
127.0.0.1:6379> sinterstore destination key [key ...]
```

### 2.5.4 集合的应用示例：
* 微博的共同关注
    * key_01的关注，key_02的关注，使用命令sinter key_01 key_02取交集即可

## 2.6 有序不重复的SotedSet集合

### 2.6.1 基本操作
* 增加一个或多个元素，如果元素存在则使用新的分值
```bash
127.0.0.1:6379> zadd key score member [score member ...]
```
* 移除一个或多个元素，元素不存在则忽略
```bash
127.0.0.1:6379> zrem key member [member ...]
```
* 显示分值
```bash
127.0.0.1:6379> zscore key member
```
* 增加减少分值，increment为负数则减少，分值的增减会导致有序集合顺序的动态变化
```bash
127.0.0.1:6379> zincrby key increment memter
```
* 返回元素排名；zrevrank为逆序
```bash
127.0.0.1:6379> zrank key member
```
* 返回指定索引（排名）区间的元素；zrevrange为逆序
```bash
127.0.0.1:6379> zrange key start stop [withscores]
```
* 返回数值区间内的元素；默认返回[min，max]区间的元素；[offset] 表示跳过多少元素；[(]表示将区间修改为开区间；[-inf|+inf]表示负无穷和正无穷；zrevzrangebyscore key max min为逆序
```bash
127.0.0.1:6379> zrangebyscore key [(][-|+]min [-|+]max[)] [withscores] [limit offset count]
```
* 移除指定索引（排名）区间的元素
```bash
127.0.0.1:6379> zremrangebyrank key start stop
```
* 移除指定分值区间的元素
```bash
127.0.0.1:6379> zremrangebyscore key min max
```
* 返回集合中元素的个数
```bash
127.0.0.1:6379> zcard key
```
* 返回指定分值区间元素的个数
```bash
127.0.0.1:6379> zcount key min max
```
* 并集；numkeys表示参与并操作的key数量；weights选项与key对应，对应key的每个score都要乘以权重；aggregate选项指定并集score的聚合方式（求和，最小，最大）
```bash
127.0.0.1:6379> zunionstore destination numkeys key [key ...] [weights weight] [aggregate sum|min|max] [withscores]
```
* 交集
```bash
127.0.0.1:6379> zintestore destination numkeys key [key ...] [weights weight] [aggregate sum|min|max][withscores]
```

### 2.6.2 有序集合应用示例
* 网易云音乐歌曲的排行榜：
    * 每首歌名做为元素，每首歌对应得播放次数做为分值；
    * zrevrange key start stop 逆序获取最高播放次数的歌曲前n位。
* 新浪微博动态翻页：
    * 每条微博做为元素，发微的时间戳做为分值；
    * zrevrange key start stop 逆序获取最新时间发的微博（如果在翻页时微博出现新的动态，有序集合会动态的重新排序）。

# 3.Redis消息发布订阅

## 3.1 概念
* Redis 发布订阅(pub/sub)是一种消息通信模式：发送者(pub)发送消息，订阅者(sub)接收消息。
![f46118c91f1335065ded5ceb292fc8d7.png](en-resource://database/710:1)
* 当有新消息通过PUBLISH命令发送给频道channel1 时，这个消息就会被发送给订阅它的三个客户端。

## 3.2 命令
* 将message消息发送到指定的channel频道
```bash
127.0.0.1:6379> publish channel message
```
* 订阅一个或多个频道的信息
```bash
127.0.0.1:6379> subscribe channel [channel ... ]
```
* 退订一个或多个频道
```bash
127.0.0.1:6379> unsubscribe [channel [channel ... ]]
```
* 订阅一个或多个符合给定模式的频道
```bash
127.0.0.1:6379> psubscribe pattern [pattern ... ]
```
* 退订所有给定模式的频道
```bash
127.0.0.1:6379> punsubscribe [pattern [pattern ... ]]
```
* 查看订阅与发布系统的状态
```bash
127.0.0.1:6379> pubsub subcommand [argument [argument ...]]
```

# 4.Redis事务

## 4.1 概念
* Redis事务可以看做是一个批量执行命令的脚本，批量的命令在exec命令发送前被放入队列缓存，收到exec命令后进入事务执行，其中的一条命令若执行失败，前面执行成功的命令不会回滚，后面的命令继续执行；
* 在事务执行过程，其他客户端提交的命令请求不会插入到事务执行命令序列中。

## 4.2 命令
* 标记一个事务块的开始
```bash
127.0.0.1:6379> multi
```
* 执行所有事务块中的命令
```bash
127.0.0.1:6379> exec
```
* 取消事务
```bash
127.0.0.1:6379> discard
```
* 监视一个或多个key ，如果在事务执行之前这个或这些key被其他命令所改动，那么事务将被打断
```bash
127.0.0.1:6379> watch key [key ...]
```
* 取消WATCH命令对所有key的监视
```bash
127.0.0.1:6379> unwatch
```

# 5.Redis持久化

## 5.1 RDB（Redis DB）
* redis默认将数据库快照保存在名为dump.rdb的二进制文件中。
* 阻塞方式：![90b951e6b912d75e5eeab13dbffdd19c.png](en-resource://database/712:1)
* client执行save命令，redis server进程执行持久化操作。
* 非阻塞方式：![72c4b5fa202279d38f8925ae36ef303e.png](en-resource://database/714:1)
* clinet执行bgsave命令，server进程fork出新的子进程异步执行持久化操作。
* fork：
    * 子进程内存中存储父进程内存的指向，不必开辟新的内存空间，如果父进程的值被改变，会新开辟空间存储新值，子进程依旧指向旧值
* 自动：
    * 配置文件中的条件满足就执行bgsavesave 60 1000，Redis要满足在60秒内至少有1000个key被改动，会自动保存一次。
* 手动：
    * 由客户端发起的save，bgsave命令。
* save命令：
    * 阻塞Redis服务，无法响应客户端请求 (创建新的dump.rdb替代旧文件)：
```BASH
127.0.0.1:6379> save
```

* bgsave命令：
    * 非阻塞，Redis服务正常接收处理客户端请求；
    * Redis会fork()一个新的子进程来创建RDB文件，子进程处理完后会向父进程发送一个信号，通知它处理完毕；
    * 父进程用新的dump.rdb替代旧文件。
```BASH
127.0.0.1:6379> bgsave
```

## 5.2 AOF（AppendOnlyFile）
![407f75b259355785108466f41f1cd7b1.png](en-resource://database/716:1)

* 写入机制：
    * 不会直接写入磁盘，先将内容放入内存缓冲区，等到缓冲区被填满，或者用户执行fsync和fdatasync调用时才将缓冲区中的数据溢写入磁盘。
* 写入磁盘的策略：
    * appendfsync选项：
        * Always：服务器每写入一个命令，就调用fdatasync将缓冲区的数据写入硬盘；
        * Everysec (default)：服务器每一秒重调用一次fdatasync，将缓冲区中的数据写入硬盘；
        * No：服务器不调用fdatasync，由操作系统决定何时将缓冲区中的数据写入硬盘。
* AOF重写机制：
    * 避免aof文件过大，会合并重复的操作，aof会使用尽可能少的命令来记录。
* 重写过程：
    * fork一个子进程负责重写AOF文件；
    * 子进程会创建一个临时文件写入AOF信息；
    * 父进程会开辟一个内存缓冲区接收新的写命令；
    * 子进程重写完成后，父进程会获得一个信号，将父进程接收到的新的写操作由子进程写入到临时文件中；
    * 新文件替代旧文件。

# 6.Redis主从复制集群
![1f41134387e71fa32f8a4f3c9a58c954.png](en-resource://database/718:1)
![9372ee3410134e1ec1fc7c2729fa6d7c.png](en-resource://database/720:1)

## 6.1 Redis主从复制集群概念
* 一个Redis服务可以有多个该服务的复制品，这个Redis服务称为Master，其他复制品称为Slaves；
* 只要网络连接正常，Master会一直将自己的数据更新同步给Slaves，保持主从数据同步；
* 只有Master可以执行写命令，Slaves只能执行读命令。

## 6.2 单节点多实例搭建伪主从复制集群
```BASH
[root@basic ~]# cd ~
[root@basic ~]# mkdir redis
[root@basic ~]# cd redis
[root@basic ~]# mkdir 6380          // 创建集群角色的目录
[root@basic ~]# mkdir 6381
[root@basic ~]# mkdir 6382
[root@basic ~]# cd 6380
[root@basic ~]# redis-server --port 6380
[root@basic ~]# cd 6381
[root@basic ~]# redis-server --port 6381 --slaveof 127.0.0.1 6380   // 启动slave节点追随主节点
[root@basic ~]# cd 6382
[root@basic ~]# redis-server --port 6382 --slaveof 127.0.0.1 6380    
[root@basic ~]# redis-cli -p 6380   // 启动客户端连接主从节点确认
```

## 6.3 Redis哨兵网络概念
![ad8ff1b0467a08f15220de7bb663cb4d.png](en-resource://database/722:1)

* Sentinel会不断检查Master和Slaves是否正常。
* 每一个Sentinel可以监控任意多个Master和该Master下的Slaves。
* 监控同一个Master的Sentinel会自动连接，组成一个分布式的Sentinel网络，互相通信并交换彼此关于被监视的服务器信息。
* Sentinel会在Master下线后自动执行Failover操作，提升一台Slave为Master，并让其他Slaves重新成为新Master的Slaves。

## 6.4 哨兵网络伪集群搭建 (实现主节点的高可用)
```BASH
$> cd ~
$> mkdir sentinel
$> cd sent
$> vim s1.conf      # 创建哨兵节点的配置文件
    port 26380
    sentinel monitor sxt 127.0.0.1 6380 2
$> vim s2.conf
    port 26381    # 配置哨兵节点的占用端口
    sentinel monitor sxt 127.0.0.1 6380 2        # sentinel moniter表示哨兵启动命令，sxt表示主从集群的逻辑名称，2表示集群一致性过半原则
$> vims3.conf
    port 26382
    sentinel monitor sxt 127.0.0.1 6380 2
$> cd ~/redis-2.8.18/src/
$> cp redis-sentinel /opt/sxt/redis/bin    # 复制哨兵节点的启动命令
$> redis-sentinel ~/sent/s1.conf           # 根据对应的配置文件启动哨兵网络
$> redis-sentinel ~/sent/s2.conf
$> redis-sentinel ~/sent/s3.conf
```

# 7.Redis分布式集群

## 7.1 概念
![540913033a8a35615c8c9ac3d874b62c.png](en-resource://database/724:1)

* 由多个Redis服务器组成的分布式网络服务器集群，每一个Redis服务器称为节点Node，节点之间会互相通信，两两相连（Redis集群无中心节点）。

## 7.2 节点复制
![664db7d880ad8c6c44f6c6c1e74f2b59.png](en-resource://database/726:1)

* Redis集群的每个节点都有两种角色可选，主节点master node和从节点slave node，其中主节点用于存储数据，而从节点则是对应主节点的镜像复制。
* 当用户需要处理更多读请求的时候，添加从节点可以扩展系统的读性能，因为Redis集群重用了单机Redis复制特性的代码，所以集群的复制行为和单机复制特性的行为是完全一样的。

## 7.3 故障转移
![1129ba61a647b791f35aaa8b50cdaa54.png](en-resource://database/728:1)
![430a28c2a4eac6425624fccd7e6c6008.png](en-resource://database/730:1)

* Redis集群的主节点内置了类似Redis Sentinel的节点故障检测和自动故障转移功能，当集群中的某个主节点下线时，集群中的其他在线主节点会注意到这一点，并对已下线的主节点进行故障转移。
* 集群进行故障转移的方法和Redis Sentinel进行故障转移的方法基本一样，不同的是，在集群里面，故障转移是由集群中其他在线的主节点负责进行的，所以集群不必另外使用Redis Sentinel。

## 7.4 集群分片
* 集群将整个数据库分为16384个slot槽位，所有key都属于这些slot中的一个，key的槽位计算公式为`slot_number=crc16(key)%16384`，其中crc16为16位的循环冗余校验和函数。
* 集群中的每个主节点都可以处理0个至16383个槽，当16384个槽都由各自的节点在负责处理时，集群进入上线状态，并开始处理客户端发送的数据命令请求。
* 数据库中所有的数据由集群中各个节点承载一段范围，clinet向某节点请求的数据如果不存在本地，则会指导clinet向存在该数据的节点重新发送请求。

## 7.5 请求转向
![ef170cfedca5e0ed6622797637101a68.png](en-resource://database/732:1)
![d0fcf815b3e15ac9f2027326a31a2713.png](en-resource://database/734:1)

* 由于Redis集群无中心节点，clinet请求会发送给集群中的任意节点。
* 当前主节点只会处理自己负责槽位的命令请求，如果是其它槽位的命令请求，该主节点会返回给客户端一个转向错误。
* 客户端根据错误中包含的地址和端口重新向正确负责的主节点发起命令请求。

## 7.6 单节点多实例搭建伪分布式集群
```BASH
Redis实例1：
[root@basic ~]# cd /opt/sxt/redis
[root@basic ~]# rm -rf bin
[root@basic ~]# cd ~/redis-cluster
[root@basic ~]# tar -zxvf redis-3.0.4.tar.gz
[root@basic ~]# cd redis-3.0.4
[root@basic ~]# make && make PREFIX=/opt/sxt/redis install
[root@basic ~]# yum install ruby rubygems -y
[root@basic ~]# cd ../
[root@basic ~]# gem install --local redis-3.3.0.gem
[root@basic ~]# cd redis-test
[root@basic ~]# cd 7000
[root@basic ~]# redis-server redis.conf

Redis实例2：
[root@basic ~]# cd ~/redis-cluster/redis-test/7001
[root@basic ~]# redis-server redis.conf

Redis实例3：
[root@basic ~]# cd ~/redis-cluster/redis-test/7002
[root@basic ~]# redis-server redis.conf

Redis实例4：
[root@basic ~]# cd ~/redis-cluster/redis-test/7003
[root@basic ~]# redis-server redis.conf

Redis实例5：
[root@basic ~]# cd ~/redis-cluster/redis-test/7004
[root@basic ~]# redis-server redis.conf

Redis实例6：
[root@basic ~]# cd ~/redis-cluster/redis-test/7005
[root@basic ~]# redis-server redis.conf
[root@basic ~]# cd ~/redis-cluster/redis-3.0.4/src
[root@basic ~]# ./redis-trib.rb create --replicas 1 127.0.0.1:7000 127.0.0.1:7001 127.0.0.1:7002 127.0.0.1:7003 127.0.0.1:7004 127.0.0.1:7005         # --replicas 1 表示主从模型为1主1从，共两个节点为一组；redis会自动将后面填写的ip和端口进行/2计算，得出共3组主从模型，那么前3个主机为分布式集群的主节点，后3个为对应主节点镜像复制的从节点
[root@basic ~]# redis-cli -p 7000 -c       # 启动redis client，指定连接7000端口的节点，-c表示集群模式
```