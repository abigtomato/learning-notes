[TOC]

# 1.ES理论知识

## 1.1.集群概念

* Cluster：集群，有多个节点，其中有一个为主节点，这个主节点是可以通过选举产生的，主从节点是相对于集群内部来说。es的一个概念就是去中心化，字面上理解就是无中心节点，这是对于集群外部来说的，因为从外部来看es集群，在逻辑上是个整体，与任意节点的通信和与整个es集群通信是等价的。

* Shards：索引分片，es可以把一个完整的索引分成多个分片，好处是可以把一个大的索引拆分成多个小的，然后分布到不同的节点上，构成分布式存储。分片的数量只能在索引创建前指定，并且索引创建后不能更改。当进行水平扩容时，需要重新设置分片数量，重新导入数据。

* Replicas：索引副本，es可以设置多个索引的副本，副本的作用一是提高系统的容错性，当某个节点某个分片损坏或丢失时可以从副本中恢复；二是提高es的查询效率，es会自动对搜索请求进行负载均衡。 

* Recovery：数据恢复或数据重分布，es集群在有节点加入或退出时会根据机器的负载对索引分片进行重新分配，挂掉的节点重新启动时也会进行数据恢复。 

* River：es的数据源，也是其它存储方式（如：数据库）同步数据到es的一个方法。它是以插件方式存在的服务，通过读取river中的数据并把它在es中建立索引。

* Gateway：es索引快照的存储方式，es默认是先把索引存放到内存中，当内存满了再持久化到本地硬盘。当es集群关闭再重新启动时就会从gateway中读取索引备份数据。es支持多种类型的gateway，有本地文件系统（默认），分布式文件系统，Hadoop的HDFS和Amazon的S3云存储服务。

* Discovery.zen：代表es的自动发现节点机制，es是一个基于p2p的系统，它先通过广播寻找存在的节点，再通过广播协议来进行节点之间的通信，同时也支持点对点的交互。 

* Transport：代表es内部节点或集群与客户端的交互方式，默认内部是使用tcp协议进行交互，同时它支持http协议（json格式），Thrift，Servlet，Memcached，ZeroMQ等的传输协议（通过插件方式集成）。

## 1.2.倒排索引

对存储的数据分词抽取出其中的各个词条，以词条为key，对应数据的出现位置为value。搜索时，对关键字分词，通过词条匹配倒排索引，获取词条在原始数据中出现的位置，以位置作为条件搜索。

* 数据表：

|商品主键|商品名|商品描述|
|:-:|:-:|:-:|
|1|荣耀10|更贵的手机|
|2|荣耀8|相对便宜的手机|
|3|iphone11|要卖肾买的手机|

* 倒排索引表：

|词条（key）|数据（value）|
|:-:|:-:|
|手机|1, 2, 3|
|便宜|2|
|卖肾|3|
|相对|2|
|荣耀|1, 2|
|iphone|3|

## 1.3.存储概念
* Index：索引，包含若干相似结构的document数据，如：客户索引，订单索引，商品索引等，一个index包含多个document，代表一类相似的或相同的document，如：订单索引中存放了所有的订单数据。
* Type：类型，每个索引中都可以有一个Type，Type是Index中的一个逻辑分类，同一个Type中的Document都有相同的field。示例：订单索引，不同状态的订单包含不同的内容，如：未支付订单（自动取消时间）和已支付订单（支付时间）、已发货订单（发货时间、物流信息）等都有不同的内容。

* Document：文档是ES中的最小数据单元，一个Document就是一条数据，一般使用json数据结构表示，每个Index下的Type中都可以存储多个Document，一个Document中有多个field，field就是数据字段。

## 1.4.ES和关系型数据库的对比

|      ES       | 数据库系统  |
| :-----------: | :---------: |
|  Index 索引   | Database 库 |
|   Type 类型   |  Table 表   |
| Document 文档 |   Row 行    |
|  Field 字段   |  Column 列  |



# 2.基本 RESTful API 操作

## 2.1.查看集群健康状态
```http
GET _cat/health?v
```

status 状态：
* **green**：每个索引的primary shard和replica shard都是active的；
* **yellow**：每个索引的primary shard都是active的，但部分的replica shard不是active的；
* **red**：不是所有的索引都是primary shard都是active状态的。

## 2.2.查看索引的shard信息

```http
GET _cat/shards?v
```

## 2.3.设置磁盘限制
ES默认当磁盘空间不足15%时，会禁止分配replica shard，可以动态调整ES对磁盘空间的要求限制

```http
PUT _cluster/settings
{
	"transient": {
    	"cluster.routing.allocation.disk.watermark.low": "95%",
    	"cluster.routing.allocation.disk.watermark.high": "5GB"
	}
}
```
**注**：配置磁盘空间限制的时候，要求low必须比high大，可以使用百分比或GB的方式设置，且ES要求low至少满足磁盘95%的容量；

此处配置的百分比都是磁盘的使用百分比，如85%，代表磁盘使用了85%后如何限制，配置的GB绝对值都是剩余空间多少：

* low：对磁盘空闲容量的最低限制，默认85%

* high：对磁盘空闲容量的最高限制，默认90%

**如**：low为50GB，high为10GB，则当磁盘空闲容量不足50GB时停止分配replica shard，当磁盘空闲容量不足10GB时，停止分配shard，并将应该在当前结点中分配的shard分配到其他结点中；ES中默认的限制是：如果磁盘空间不足15%的时候，不分配replica shard，如果磁盘空间不足5%的时候，不再分配任何的primary shard。

## 2.4.查看索引信息
```http
GET _cat/indices?v
```

## 2.5.新增 index
在ES中默认创建索引的时候，会分配5个primary shard，并为每个primary shard分配一个 replica shard。

```http
PUT /test_index
{
  	"settings":{
    	"number_of_shards" : 2,		// 指定该索引分片数量
    	"number_of_replicas" : 1	// 指定每个分片的副本数量
  	}
}
```
## 2.6.修改 index

es中对shard的分布是有要求的，有其内置的特殊算法。es尽可能保证primary shard平均分布在多个节点上，replica shard会保证不和他备份的那个primary shard分配在同一个节点上。

**注**：索引一旦创建，primary shard数量不可变化，但可以改变replica shard数量。 

```http
PUT /test_index/_settings
{
	"number_of_replicas" : 2
}
```
## 2.7.删除 index
```http
DELETE /test_index [other_index, ...]
```

## 2.8.新增 document
es有自动识别机制，如果增加的document对应的index不存在，则自动创建；如果index存在，type不存在，也会自动创建；如果index和type都存在，则使用现有的。

### 2.8.1.PUT 语法
此操作为手工指定id的Document新增方式。

```http
PUT /test_index/test_type/1
{
	"name": "test_doc_01",
   	"remark": "first test elastic search",
   	"order_no": 1
}
```
```json
{
    "_index": "test_index",	// document所属index
    "_type": "test_type",	// document所属type
    "_id": "1",				// 指定的id
    "_version": 1,			// document版本，版本从1开始递增，每次写操作都会+1
    "result": "created",	// 本次操作的类型（created创建，updated修改，deleted删除）
    "_shards": {			// 分片信息
        "total": 2,			// 分片数量只提示primary shard
        "successful": 1,	// 数据document一定只存放在index中的某一个primary shard中
        "failed": 0
	},
    "_seq_no": 0,           // 执行的序列号
    "_primary_term": 1      // 词条比对
}
```

### 2.8.2 POST 语法
此操作为ES自动生成id的新增Document方式。

**注**：在ES中，一个index中的所有type类型的Document是存储在一起的，如果index中的不同的type之间的field差别太大，也会影响到磁盘的存储结构和存储空间的占用。

```http
POST /test_index/test_type
{
   	"name": "test_doc_02",
   	"remark": "first test elastic search",
   	"order_no": 2
}
```

如：index中有type1和type2两个不同的类型：
* type1中的document结构为：`{"_id": "1", "f1": "v1", "f2": "v2"}`；
* type2中的document结构为：`{"_id": "2", "f3": "v3", "f4": "v4"}`；
* 那么ES存储时的统一存储方式是：`{"_id": "1", "f1": "v1", "f2": "v2", "f3": "", "f4": ""}, {"_id": "2", "f1": "", "f2": "", "f3": "v3", "f4": "v4"}`；
* 建议每个index中存储的document结构不要有太大的差别，尽量控制在总计字段数据的10%以内。

## 2.9 查询 document

### 2.9.1.GET 查询
```http
GET /test_index/test_type/1
```
```json
{
    "_index": "test_index",
    "_type": "test_type",
    "_id": "1",
    "_version": 1,
    "found": true,
    // 查询结果
    "_source": {
		"name": "test_doc_01",
		"remark": "first test elastic search",
		"order_no": 1
	}
}
```

### 2.9.2.GET_mget 批量查询
批量查询可以提高查询效率，推荐使用（相对于单数据查询来说） 

```http
GET /_mget
{
    "docs": [
        {
            "_index": "test_index",
            "_type": "test_type",
            "_id": "1"
        },
        {
            "_index": "test_index",
            "_type": "test_type",
            "_id": "2"
        }
    ]
}
```
```http
GET /test_index/_mget
{
    "docs": [
        {
            "_type": "test_type",
            "_id": "1"
        },
        {
            "_type": "test_type",
            "_id": "2"
        }
    ]
}

```
```http
GET /test_index/test_type/_mget
{
    "docs": [
        {
        	"_id": "1"
        },
        {
        	"_id": "2"
        }
    ]
}
```

## 2.10.修改 document

### 2.10.1.全量替换
要求新数据的字段信息和原数据的字段信息一致，也就是必须包括Document中的所有field才行，本操作相当于覆盖操作。全量替换的过程中，ES不会真的修改Document中的数据，而是标记ES中原有的Document为deleted状态，再创建一个新的Document来存储数据，当ES中的数据量过大时，ES后台回收deleted状态的Document。

```http
PUT /test_index/test_type/1
{
   "name": "new_test_doc_01",
   "remark": "first test elastic search",
   "order_no": 1
}
```
```JSON
{
    "_index": "test_index",
    "_type": "test_type",
    "_id": "1",
    "_version": 2,
    "result": "updated",
    "_shards": {
        "total": 2,
        "successful": 1,
        "failed": 0
    },
    "_seq_no": 1,
    "_primary_term": 1
}
```

### 2.10.2.PUT 语法强制新增
如果使用PUT语法对同一个Document执行多次操作，是一种全量替换操作。如果需要ES辅助检查PUT的Document是否已存在，可以使用强制新增语法。使用强制新增语法时，如果Document的id在ES中已存在，则会报错。

```http
PUT /test_index/test_type/1/_create
{
   "name": "new_test_doc_01",
   "remark": "first test elastic search",
   "order_no": 1
}
```
```http
PUT /test_index/test_type/1?op_type=create
{
   "name": "new_test_doc_01",
   "remark": "first test elastic search",
   "order_no": 1
}
```

### 2.10.3.partial update 更新 document
只更新某Document中的部分字段。这种更新方式也是标记原有数据为deleted状态，创建一个新的Document数据，将新的字段和未更新的原有字段组成这个新的Document并创建。对比全量替换而言，只是操作上的方便，在底层执行上几乎没有区别。

```http
POST /test_index/test_type/1/_update
{
    "doc": {
    	"name": "test_doc_01_for_update"
    }
}
```
```JSON
{
    "_index": "test_index",
    "_type": "test_type",
    "_id": "1",
    "_version": 5,
    "result": "updated",
    "_shards": {
        "total": 2,
        "successful": 1,
        "failed": 0
    },
    "_seq_no": 2,
    "_primary_term": 1
}
```

## 2.11.删除 document
ES中执行删除操作时，会先标记Document为deleted状态，而不是直接物理删除。当ES存储空间不足或工作空闲时，才会执行物理删除操作。标记为deleted状态的数据不会被查询搜索到。ES中删除index，也是标记，后续才会执行物理删除，所有的标记动作都是为了NRT的实现（近实时）。

```http
DELETE /test_index/test_type/1
```
```json
{
    "_index": "test_index",
    "_type": "my_type",
    "_id": "1",
    "_version": 6,
    "result": "deleted",
    "_shards": {
        "total": 2,
        "successful": 1,
        "failed": 0
    },
    "_seq_no": 5,
    "_primary_term": 1
}
```

## 3.12.\_bulk 语法批量增删改
create：强制创建，相当于 PUT /index_name/type_name/id/_create

index：普通的PUT操作，相当于创建Document或全量替换

update ：更新操作（partial update），相当于 POST /index_name/type_name/id/\_update

delete：删除操作

**案例**：

```JSON
POST /_bulk
{
    "create": {
        "_index": "test_index",
        "_type": "my_type",
        "_id": "1"
    }
}
{"field_name": "field value"}
```
```JSON
POST /_bulk
{"index": {"_index": "test_index", "_type": "my_type", "_id": "2"}}
{"field_name": "field value 2"}
```
```JSON
POST /bulk
{"update": {"_index": "test_index", "_type": "my_type", "_id": 2", "_retry_on_conflict": 3}}
{"doc": {"field_name": "partial update field value"}}
```
```JSON
POST /_bulk
{"delete": {"_index": "test_index", "_type": "my_type", "_id": "2"}}
```
* 可以一次性执行增删改的所有功能：
```JSON
POST /_bulk
{"create": {"_index": "test_index", "_type": "my_type", "_id": 10}}
{"name": "hadoop"}
{"index": {"_index": "test_index", "_type": "my_type", "_id": 20}}
{"name": "spark"}
{"update": {"_index": "test_index", "_type": "my_type", "_id": 20, "_retry_on_conflict": 3}}
{"doc": {"name": "spark mllib"}}
{"delete": {"_index": "test_index", "_type": "my_type", "_id": 2}}
```

> **注**：bulk语法中要求一个完整的json串不能有换行，不同的json串必须使用换行分隔。多个操作中，如果有错误情况，不会影响到其他的操作，只会在批量操作返回结果中标记失败。bulk语法批量操作时，bulk request会一次性加载到内存中，如果请求数据量太大，性能反而下降（内存压力过高），需要反复尝试一个最佳的bulk request size。一般从1000~5000条数据开始尝试，逐渐增加，如果查看bulk request size的话，一般是5~15MB之间。

> **解释**：bulk语法要求json格式是为了对内存的方便管理，和尽可能降低内存的压力。如果json格式没有特殊的限制，ES在解释bulk请求时，需要对任意格式的json进行解释处理，需要对bulk请求数据做json对象与json array对象的转化，那么内存的占用量至少翻倍，当请求量过大的时候，对内存的压力会直线上升，且需要jvm gc进程对垃圾数据做频繁回收，影响ES效率。
生产环境中，bulk api常用，都是使用java代码实现循环操作，一般一次bulk请求，执行一种操作。如：批量新增10000条数据等。

## 3.13 Document Routing 路由机制
> * ES使用一个路由算法管理Document，这种算法决定了Document存放在哪一个Primary Shard中。

> * **算法为**：**Primary Shard = Hash(routing) % Number_Of_Primary_Shards**
其中的routing默认为Document中的元数据_id，也可以手工指定routing的值，指定方式为：PUT /index_name/type_name/id?routing=xxx。手工指定routing在海量数据中非常有用，通过手工指定的routing，使ES将相关联的Document存储在同一个shard中，方便后期进行应用级别的负载均衡并可以提高数据检索的效率。如：存储电商中的商品，使用商品类型的编号作为routing，ES会把同一个类型的商品Document数据，存在同一个shard中。查询的时候，同一个类型的商品，在一个shard上查询，效率最高。

> * 如果是写操作，计算routing结果后，决定本次写操作定位到哪一个Primary Shard分片上，Primary Shard 分片写成功后，自动同步到对应Replica Shard上。

> * 如果是读操作，计算routing结果后，决定本次读操作定位到哪一个Primary Shard或其对应的Replica Shard上，实现读负载均衡，Replica Shard数量越多，并发读能力越强。

```JSON
# 手动指定routing参数
PUT /test_index/my_type/100?routing=1
{
  "100_field_name_01": "value_01",
  "100_field_name_02": "value_02"
}

PUT /test_index/my_type/200?routing=2
{
  "200_field_name_01": "value_01",
  "200_field_name_02": "value_02"
}

PUT /test_index/my_type/300?routing=3
{
  "300_field_name_01": "value_01",
  "300_field_name_02": "value_02"
}
```
## 3.14 Document 查询原理
客户端发起执行查询操作的请求，查询操作都由Primary Shard和Replica Shard共同处理，此操作请求到节点2（请求发送到的节点随机），这个节点称为协调节点（coordinate node）；

协调节点通过路由算法，计算出本次查询的Document所在的Shard，假设本次查询的Document所在shard为 Shard 0，协调节点计算后，会将操作请求转发到节点1或节点3，至于分配请求到节点1还是节点3是通过随机算法或负载均衡算法计算的，ES会保证当请求量足够大的时候，Primary Shard和Replica Shard处理的查询请求数是均等的（不绝对一致）；

节点1或节点3中的Primary Shard 0或Replica Shard 0在处理请求后，会将查询结果返回给协调节点（节点2）；

协调节点得到查询结果后，再将查询结果返回给客户端。

## 3.15 Document 增删改原理
客户端发起执行增删改操作的请求，所有的增删改操作都由Primary Shard直接处理，Replica Shard只被动的备份数据，此操作请求到节点2 (请求发送到的节点随机)，这个节点称为协调节点(coordinate node)；

协调节点通过路由算法，计算出本次操作的Document所在的shard，假设本次操作的Document所在的shard为 Primary Shard 0，协调节点计算后，会将操作请求转发到节点1；

节点1中的Primary Shard 0在处理请求后，会将数据的变化同步到对应的Replica Shard 0中，也就是发送一个同步数据的请求到节点3中；

Replica Shard 0在同步数据后，会响应通知同步成功，也就是响应给Primary Shard 0（节点1）；

Primary Shard 0（节点1）接收到Replica Shard 0的同步成功响应后，会响应请求者，本次操作完成，也就是响应给协调节点（节点2）；

协调节点返回响应给客户端，通知操作结果。

## 3.16 Document 搜索

### 3.16.1 Query String Search
> * search的参数类似于http请求头中的字符串参数提供搜索条件的。

* **语法**：
```HTTP
GET [/index_name/type_name/]_search[?parameter_name=parameter_value&...]
```
* **例**：全数据搜索（也就是没有搜索条件）
```HTTP
GET /test_index/my_type/_search
```
```JSON
{
  "took": 8,            # 执行的时长。单位毫秒。
  "timed_out": false,   # 是否超时
  "_shards": {          # shard 相关数据
    "total": 5,         # 总计多少个shard
    "successful": 5,    # 成功返回结果的shard数量
    "skipped": 0,
    "failed": 0	
  },
  "hits": {             # 搜索结果相关数据，
    "total": 3,         # 总计多少数据，符合搜索条件的数据数量。
    "max_score": 1,     # 最大相关度分数。和搜索条件的匹配度。
    "hits": [           # 具体的搜索结果
      {
        "_index": "test_index",   # 索引名称
        "_type": "my_type",       # 类型名称
        "_id": "2",               # id值
        "_score": 1,              # 匹配度分数，本条数据匹配度分数
        "_source": {              # 具体的数据内容，源
          "name": "test_doc_02",
          "remark": "second test elastic search",
          "order_no": 2
        }
      }
    ]
  }
}
```
* **例**：搜索remark字段关键字为test的document，搜索结果按照order_no字段降序排序
```HTTP
GET /test_index/my_type/_search?q=remark:test&sort=order_no:desc
```
```json
{
  "took": 17,
  "timed_out": false,
  "_shards": {
    "total": 5,
    "successful": 5,
    "skipped": 0,
    "failed": 0
  },
  "hits": {
    "total": 3,
    "max_score": null,
    "hits": [
      {
        "_index": "test_index",
        "_type": "my_type",
        "_id": "3",
        "_score": null,
        "_source": {
          "name": "test_doc_03",
          "remark": "third test elastic search",
          "order_no": 3
        },
        "sort": [
          3
        ]
      },
      {
        "_index": "test_index",
        "_type": "my_type",
        "_id": "2",
        "_score": null,
        "_source": {
          "name": "test_doc_02",
          "remark": "second test elastic search",
          "order_no": 2
        },
        "sort": [
          2
        ]
      },
      {
        "_index": "test_index",
        "_type": "my_type",
        "_id": "1",
        "_score": null,
        "_source": {
          "name": "test_doc_01",
          "remark": "first test elastic search",
          "order_no": 1
        },
        "sort": [
          1
        ]
      }
    ]
  }
}
```
> **注**：此搜索操作一般只用在快速检索数据使用，如果查询条件复杂，很难构建Query String。

### 3.16.2 Query DSL
* 查询所有数据：
```JSON
GET /test_index/my_type/_search
{
  "query": {
    "match_all": {}
  }
}
```
* 条件查询 + 排序：
```JSON
GET /test_index/my_type/_search
{
  "query": {            # 查询条件
    "match": {          # 模糊匹配，包含匹配
      "remark": "test"  # [字段: 关键字]
    }
  },
  "sort": [             # 排序条件
    {
      "order_no": {     # 排序字段
        "order": "asc"  # 升序排序
      }
    }
  ]
}
```
* 分页查询：
```JSON
GET /test_index/my_type/_search
{
  "query": {
    "match_all": {}
  },
  "from": 0,    # 从第几条数据开始查询，从0开始计数
  "size": 2,    # 查询多少数据
  "sort": [
    {
      "order_no": "desc"  # 降序排序
    }
  ]
}
```
* 部分字段查询：
```JSON
GET /test_index/my_type/_search
{
  "query": {
    "match": {
      "name": "test_doc_03"
    }
  },
  "sort": [
    {
      "order_no": {
        "order": "desc"
      }
    }
  ],
  "_source": ["name", "remark"],  # 字段限定
  "from": 0,
  "size": 20
}
```
> **注**：此搜索操作适合构建复杂查询条件，生产环境常用。

### 3.16.3 Query Filter
> * 过滤查询，此操作实际上就是Query DSL的补充语法。过滤的时候，不进行任何的匹配分数计算，相对于query来说，filter相对效率较高。Query要计算搜索匹配相关度分数。Query更加适合复杂的条件搜索。

* 不使用filter语法，order_no字段需要计算相关度分数：
```JSON
GET /test_index/my_type/_search
{
  "query": {
    "bool": {         # 多条件搜索，内部的若干条件，只要有一个条件符合即可
      "must": [       # 内部若干条件，必须都匹配成功才有结果
        {
          "match": {  # 字段中必须包含才有结果
            "remark": "elastic"
          }
        },
        {
          "range": {        # 字段的数据必须在某一范围存在才有结果
            "order_no": {
              "gte": 1,     # gt, gte -> 大于，大于等于
              "lte": 4      # lt, lte -> 小于，小于等于
            }
          }
        }
      ],
      "must_not": [         # 内部若干条件，必须都不匹配才有结果
        {
          "match": {        # 此处条件的意思就是name字段不包含mapreduce关键字
            "name": "mapreduce"
          }
        }
      ]
    }
  }
}
```
* 使用filter语法，不需要计算任何的相关度分数：
```JSON
GET /test_index/my_type/_search
{
  "query": {
    "bool": {
      "must": [
        {
          "match": {
            "remark": "elastic"
          }
        }
      ],
      "filter": [     # 在已有的搜索结果中进行过滤，满足内部条件的返回
        {
          "range": {
            "order_no": {
              "gte": 1,
              "lte": 4
            }
          }
        }
      ]
    }
  }
}
```

### 3.16.4 Full-Text Search
> * 全文检索：要求查询条件拆分后的任意词条与具体数据匹配就算搜索结果，查询结果的显示顺序默认与匹配度分数相关。

* **例**：搜索remark字段中包含elastic或search的document（多关键字包含匹配）
```JSON
GET /test_index/my_type/_search
{
  "query": {
    "match": {
      "remark": "elastic search"
    }
  }
}
```

### 3.16.5 Phrase Search
> * 短语检索精确匹配，要求查询条件必须和具体数据完全匹配才算搜索结果。

* **例**：搜索remark字段内容是 "first test elastic search" 的document
```JSON
GET /test_index/my_type/_search
{
  "query": {
    "match_phrase": {
      "remark": "first test elastic search"
    }
  }
}
```

### 3.16.6 Highlight Display
> * 高亮显示，高亮的不是搜索条件而是显示逻辑，在搜索的时候，经常需要对查询条件实现高亮显示。

* **例**：
```JSON
GET /test_index/my_type/_search
{
  "query": {
    "match": {
      "remark": "elastic search"
    }
  },
  "highlight": {
    "fields": {
      "remark": {
        "number_of_fragments": 2,   # 显示多少条件
        "fragment_size": 2          # 高亮显示条件的个数
      }
    }
  }
}
```

## 3.17 聚合搜索
> ES提供的aggs语法，使用DSL搜索的语法，实现聚合数据的统计，查询。

### 3.17.1 准备数据源
```JSON
PUT /products_index/phone_type/1
{
   "name": "IPHONE 8",
   "remark": "64G",
   "price": 548800,
   "producer": "APPLE",
   "tags": [ "64G", "red color", "Nano SIM" ]
}
PUT /products_index/phone_type/2
{
   "name": "IPHONE 8",
   "remark": "64G",
   "price": 548800,
   "producer": "APPLE",
   "tags": [ "64G", "golden color", "Nano SIM" ]
}
PUT /products_index/phone_type/3
{
   "name": "IPHONE 8 PLUS",
   "remark": "128G",
   "price": 748800,
   "producer": "APPLE",
   "tags": [ "128G", "red color", "Nano SIM" ]
}
PUT /products_index/phone_type/4
{
   "name": "IPHONE 8 PLUS",
   "remark": "256G",
   "price": 888800,
   "producer": "APPLE",
   "tags": [ "256G", "golden color", "Nano SIM" ]
}
```

> * 将文本类型的field的fielddata设置为true。用于将ES中的倒排索引内容重设一份正排索引，并提供内存存储计算能力；
> * 正排索引：类似数据库中的普通索引，倒排索引不做二次解析，通过分析后的词条信息，根据文档建立索引，如该词条在哪些文档中出现。索引用于内存计算，如：分析，分组，字符串排序等。

```JSON
PUT /products_index/_mapping/phone_type
{
   "properties": {
      "tags": {
         "type": "text",
         "fielddata": true
      }
   }
}
```

### 3.17.2 聚合统计
> * terms：分组计数，按照准词组分组，并统计每组document的数量，类似数据库中的count函数。
> * 聚合搜索：语法的大体结构和DSL搜索语句类似，类似SQL中的 `select count(\*) from table;`

* **例**：
```JSON
GET /products_index/phone_type/_search
{
   "size" : 0,                # 代表显示多少计算源数据Document
   "aggs" : {                 # 开始聚合的标志，类似query，是一个api
      "group_by_tags":{       # 给此聚合加一个自定义的命名
        "terms" : {           # terms是一个聚合api，类似数据库中的聚合函数，解析某字段中的词条，按照词条统计出现次数。如：a字段的值是test field，假设解析后词条为test和field，那么就是根据a字段的解析词条test和field来统计这两个数据在多少个document中存在
          "field" : "tags"
        }
      }
   }
}
```
```json
{
  "took": 3,
  "timed_out": false,
  "_shards": {
    "total": 5,
    "successful": 5,
    "skipped": 0,
    "failed": 0
  },
  "hits": {
    "total": 4,
    "max_score": 0,
    "hits": []
  },
  "aggregations": {
    "group_by_tags": {
      "doc_count_error_upper_bound": 0,
      "sum_other_doc_count": 0,
      "buckets": [
        {
          "key": "color",     # 分析后的词条
          "doc_count": 4      # 在多少个document出现
        },
        {
          "key": "nano",
          "doc_count": 4
        },
        {
          "key": "sim",
          "doc_count": 4
        },
        {
          "key": "64g",
          "doc_count": 2
        },
        {
          "key": "golden",
          "doc_count": 2
        },
        {
          "key": "red",
          "doc_count": 2
        },
        {
          "key": "128g",
          "doc_count": 1
        },
        {
          "key": "256g",
          "doc_count": 1
        }
      ]
    }
  }
}
```

### 3.17.3 增加匹配条件的聚合统计
> * 搜索名称中包含PLUS的Document，并计算每个tag中的Document数量，统计是search中的一部分，一般在DSL query中使用，所以经常和条件搜索配合完成统计。 

* 条件匹配 + 聚合统计：
```JSON
GET /products_index/phone_type/_search
{
  "size": 0, 
  "query": {
    "match": {
      "name": "plus"
    }
  },
  "aggs": {
    "group_by_tags": {
      "terms": {
        "field": "tags"
      }
    }
  }
}
```

### 3.17.4 聚合后实现计算
* 平均值聚合统计 + 条件匹配，计算name字段包含plus的document的price字段的平均值：
```JSON
GET /products_index/phone_type/_search
{
  "size": 0, 
  "query": {
    "match": {
      "name": "plus"
    }
  },
  "aggs": {
    "avg_by_price": {
      "avg": {
        "field": "price"
      }
    }
  }
}
```
* 嵌套聚合，搜索包含plus的document，然后根据tags做词条统计，根据统计结果计算price字段的平均值。聚合是可以嵌套的，内层聚合是依托于外层聚合的结果之上，实现聚合计算。
```JSON
GET products_index/phone_type/_search
{
  "size": 0,
  "query": {
    "match": {
      "name": "plus"
    }
  },
  "aggs": {
    "group_by_name": {
      "terms": {
        "field": "tags"
      },
      "aggs": {
        "avg_by_price": {
          "avg": {
            "field": "price"
          }
        }
      }
    }
  }
}
```

### 3.17.5 聚合的排序
> * 类似SQL -> select … from group by … order by …
> * 聚合aggs中如果使用order排序的话，要求排序字段必须是一个aggs聚合相关的字段。
> * 聚合相关字段：当前聚合的子聚合的自定义命名。如：外部聚合是使用terms实现的聚合，命名为group_by_tags，其内层子聚合是使用avg计算平均值，聚合名称为avg_by_price，那么这个avg_by_price则称为聚合相关字段。

* **例**：计算每个包含tag的Document的price的平均值，并根据求平均后的price字段进行升序排序。
```JSON
GET products_index/phone_type/_search
{
  "size": 0,
  "aggs": {
    "group_by_tags": {
      "terms": {
        "field": "tags",
        "order": {
          "avg_by_price": "asc"
        }
      },
      "aggs": {
        "avg_by_price": {
          "avg": {
            "field": "price"
          }
        }
      }
    }
  }
}
```

### 3.17.6 范围分组并计算
* **例**：使用price的取值范围分组，再计算分组后的每组price的平均值。
```JSON
GET products_index/phone_type/_search
{
  "size": 0, 
  "query": {
    "match_all": {}
  },
  "aggs": {
    "range_by_price": {
      "range": {            # 在一段范围内搜索的条件
        "field": "price",   # 指定分组字段
        "ranges": [         # 按照指定的范围分组
          {
            "from": 500000, # price在500000 ~ 600000以内的分一组
            "to": 600000
          },
          {
            "from": 600001,
            "to": 800000
          },
          {
            "from": 800001,
            "to": 1000000
          }
        ]
      },
      "aggs": {
        "avg_by_price": {
          "avg": {
            "field": "price"
          }
        }
      }
    }
  }
}
```

# 4.常见元数据
> 在ES中，除了定义的index，type，和管理的document外，还有若干的元数据，这些元数据用于记录ES中需要使用的核心数据。在ES中元数据通常使用下划线“_”开头。 

## 4.1 查看数据
* **语法**：
```
GET /index_name/type_name/{id}
```
* **例**：
```HTTP
GET /test_index/my_type/1
```
```JSON
{
  "_index": "test_index",
  "_type": "my_type",
  "_id": "1",
  "_version": 1,
  "found": true,
  "_source": {
    "name": "test_doc_01",
    "remark": "first test elastic search",
    "order_no": 1
  }
}
```
## 4.2 \_index
> * 代表document存放在哪个index中，\_index就是索引的名字。生产环境中，类似的Document存放在一个index中，非类似的Document存放在不同的index中。一个index中包含若干相似的Document。index名称必须是小写的，且不能以下划线'_'，'-'，'+'开头。 

## 4.3 \_type
> * 代表document属于index中的哪个type（类别），就是type的名字。ES6.x版本中，一个index只能定义一个type。结构类似的document保存在一个index中。Type命名要求：字符大小写无要求，不能下划线开头，不能包含逗号。（ES低版本，5.x或更低版本。一般一个索引会划分若干type，逻辑上对index中的document进行细致的划分。在命名上，可以全大写或者全小写，不能下划线开头，不能包含逗号。） 

## 4.4 \_id
> * 代表document的唯一标识。使用index、type和id可以定位唯一的一个document。id可以在新增document时手工指定，也可以由es自动创建。

### 4.4.1 手动指定 id
```JSON
PUT /index_name/type_name/id_value
{
  "field_name" : "field_value"
}
```

> * 使用这种方式，需要考虑是否满足手动指定id的条件。如果数据是从其他数据源中读取并新增到ES中的时候，使用手动指定id。如：数据是从Database中读取并新增到ES中的，那么使用Database中的PK作为ES中的id比较合适。建议，不要把不同表的数据新增到同一个index中，可能有id冲突。 

### 4.4.2 自动生成 id
```JSON
POST /index_name/type_name
{
  "field_name" : "field_value"
}
```

> * 自动生成的ID特点：长度为20的字符串；URL安全（经过base64编码的）；GUID生成策略，支持分布式高并发（在分布式系统中，并发生成ID也不会有重复可能，参考https://baike.baidu.com/item/GUID/3352285?fr=aladdin）。适合用于手工录入的数据。数据没有一个数据源，且未经过任何的管理和存储。这种数据，是没有唯一标识，如果使用手工指定id的方式，容易出现id冲突，导致数据丢失。相对少见。

## 4.5 \_source
> * 就是查询的document中的field值。也就是document的json字符串。此元数据可以定义显示结果（field）。
> * 语法是：GET /index_name/type_name/id_value?\_source=field_name1,field_name2

## 4.6 \_version
> * 代表的是document的版本。在ES中，为document定义了版本信息，document数据每次变化，代表一次版本的变更。版本变更可以避免数据错误问题（并发问题，乐观锁），同时提供ES的搜索效率；
> * 第一次创建Document时，_version版本号为1，默认情况下，后续每次对Document执行修改或删除操作都会对_version数据进行自增；
> * 删除Document也会使_version自增；
> * 当使用PUT命令再次增加同id的Document，_version会继续之前的版本继续自增。