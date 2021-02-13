



# Hive

## 架构

* 用户接口：命令行模式（CLI），客户端模式（JDBC），WebUI模式；
  * 在cli启动的同时会启动hive的副本；
  * 启动client模式需要指出hive server所在的节点，并在该节点启动hive server。
* hive的元数据存储在关系型数据库中，如mysql，derby；
  * hive的元数据包括表的名字，表的列，分区和属性，表的数据所在目录。
* 解释器、编译器、优化器完成HQL查询语句从词法分析、语法分析、编译、优化以及查询计划的生成；
  * 生成的查询计划存储在HDFS中，并在随后有MapReduce调用执行。
* Hive的数据存储在HDFS中，大部分的查询、计算由MapReduce完成（包含\*的查询，比如 ```select * from tbl``` 不会生成MapRedcue任务）；
* 编译器将一个Hive SQL转换操作符，操作符是Hive的最小的处理单元，每个操作符代表HDFS的一个操作或者一道MapReduce作业。



## WordCount案例

**分析目标**：统计所有单词出现的次数

**数据格式展示**：

* **wc文件**

```
"License" shall mean the terms and conditions for use, reproduction, and distribution as defined by Sections 1 through 9 of this document.
"Licensor" shall mean the copyright owner or entity authorized by the copyright owner that is granting the License.
"Legal Entity" shall mean the union of the acting entity and all other entities that control, are controlled by, or are under common control with that entity. For the purposes of this definition, "control" means (i) the power, direct or indirect, to cause the direction or management of such entity, whether by contract or otherwise, or (ii) ownership of fifty percent (50%) or more of the outstanding shares, or (iii) beneficial ownership of such entity.
```

* **hive分析阶段**

```bash
hive> use sxt;
hive> create table wc(line string);
hive> create table wc_result(word string, ct int);
hive> from (select explode(split(line, ' ')) as word from wc) as tmp
    > insert overwrite table wc_result
    > select tmp.word, count(tmp.word) as ct
    > group by word
    > sort by ct desc limit 10;
```

* **看分析结果**

```bash
hive> select * from wc_result;
the	1162
to	496
of	458
```



## 基站掉话率案例

**分析目标：** 找出掉线率最高的前10基站

**数据格式展示：**

* **cdr_summ_imei_cell_info.csv文件**

| record_time（通话时间） | imei（基站编号） | cell（手机编号） | ph_num      | call_num | drop_num（掉话时间/秒） | duration（通话持续时间/秒） | drop_rate | net_type | erl  |
| ----------------------- | ---------------- | ---------------- | ----------- | -------- | ----------------------- | --------------------------- | --------- | -------- | ---- |
| 2011-07-13              | 00:00:00+08      | 356966           | 29448-37062 | 0        | 0                       | 0                           | 0         | 0        | 0    |
| 2011-07-13              | 00:00:00+08      | 352024           | 29448-51331 | 0        | 0                       | 0                           | 0         | 0        | 0    |
| 2011-07-13              | 00:00:00+08      | 353736           | 29448-51331 | 0        | 0                       | 0                           | 0         | 0        | 0    |

* **hive分析阶段**

```bash
hive> use sxt;
hive> create table cell_monitor( # 创建数据表
        record_time string,
        imei string,
        cell string,
        ph_num int,
        call_num int,
        drop_num int,
        duration int,
        drop_rate DOUBLE,
        net_type string,
        erl string
      )
      row format delimited
      fields terminated by ','
      lines terminated by '\n'
      stored as textfile;
hive> create table cell_drop_monitor( # 创建结果表
        imei string,
        total_drop_num int,
        total_call_num int,
        d_rate DOUBLE
      )
      row format delimited
      fields terminated by '\t'
      stored as textfile;
hive> load data local inpath '/roor/cdr_summ_imei_cell_info.csv' into table cell_monitor; # 本地数据导入
hive> from cell_monitor cm # 指定数据表进行分析
      insert overwrite table cell_drop_monitor # 分析结果写入结果表
      select cm.imei, sum(cm.drop_num), sum(cm.duration), sum(cm.drop_num)/sum(cm.duration) d_rate # 基站掉话率=掉话时间/总通话时间
      group by cm.imei # 对基站id进行分组
      sort by d_rate desc limit 10; # 对计算结果进行降序排序，取前10条记录
```

* **查看分析结果**

```bash
hive> select * from cell_drop_monitor;
639876	1	734	0.0013623978201634877
356436	1	1028	9.727626459143969E-4
351760	1	1232	8.116883116883117E-4
```

