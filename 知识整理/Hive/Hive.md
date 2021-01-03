[TOC]

# 1.Hive 架构
* 用户接口：命令行模式（CLI），客户端模式（JDBC），WebUI模式；
    * 在cli启动的同时会启动hive的副本；
    * 启动client模式需要指出hive server所在的节点，并在该节点启动hive server。
* hive的元数据存储在关系型数据库中，如mysql，derby；
    * hive的元数据包括表的名字，表的列，分区和属性，表的数据所在目录。
* 解释器、编译器、优化器完成HQL查询语句从词法分析、语法分析、编译、优化以及查询计划的生成；
    * 生成的查询计划存储在HDFS中，并在随后有MapReduce调用执行。
* Hive的数据存储在HDFS中，大部分的查询、计算由MapReduce完成（包含\*的查询，比如 ```select * from tbl``` 不会生成MapRedcue任务）；
* 编译器将一个Hive SQL转换操作符，操作符是Hive的最小的处理单元，每个操作符代表HDFS的一个操作或者一道MapReduce作业。

# 2.Hive 安装
* 安装mysql
```bash
• node01（MySQL Server端）：
$> yum install mysql-server # 安装mysql
$> service mysqld start # 启动mysql服务
$> chkconfig mysqld on # 开机自启
$> mysql # 进入mysql shell界面
    mysql> grant all privileges on *.* to 'root'@'%' identified by '123' with grant option; # 添加mysql权限
    mysql> use mysql;
    mysql> delete from user where host != '%'; # 删除其他权限
    mysql> flush privileges; # 刷新权限列表
    mysql> quit
$> mysql -u root -p
```
* hive单用户模式搭建
```bash
• node02（Hive CLI端）
$> tar -zxvf apache-hive-1.2.1-bin.tar.gz -C /opt/sxt
$> mv apache-hive-1.2.1-bin hive-1.2.1
$> vim /etc/profile
    export HIVE_HOME=/opt/sxt/apache-hive-1.2.1-bin
    export PATH=$PATH:$HIVE_HOME/bin
$> source /etc/profile
$> cd /opt/sxt/hive/conf
$> mv hive-default.xml.template hive-site.xml # 修改hive默认配置文件
$> vim hive-site.xml
    "shift+:",".,$-1d"
    <property>
        # 配置hdfs中默认表数据存放路径
        <name>hive.metastore.warehouse.dir</name> 
        <value>/user/hive_remote/warehouse</value>
    </property>
    <property>
        # 单用户模式本地管理
        <name>hive.metastore.local</name> 
        <value>true</value>
    </property>
    <property>
        # mysql的url
        <name>javax.jdo.option.ConnectionURL</name>
        <value>jdbc:mysql://node01:3306/hive_remote?createDatabaseIfNotExist=true</value>
    </property>
    <property>
        # mysql驱动
        <name>javax.jdo.option.ConnectionDriverName</name>
        <value>com.mysql.jdbc.Driver</value>
    </property>
    <property>
        # 用户名
        <name>javax.jdo.option.ConnectionUserName</name>
        <value>root</value>
    </property>
    <property>
        # 密码
        <name>javax.jdo.option.ConnectionPassword</name>
        <value>123</value>
    </property>
$> cp ~/mysql-connector-java-5.1.32-bin.jar /opt/sxt/apache-hive-1.2.1-bin/lib # 拷贝mysql驱动jar包
$> rm -f /opt/sxt/hadoop-2.6.5/share/hadoop/yarn/lib/jline-0.9.94.jar
$> cp /opt/sxt/apache-hive-1.2.1-bin/lib/jline-2.12.jar /usr/local/hadoop/share/hadoop/yarn/lib # 更新jar包版本
$> hive
```
* **hive多用户模式搭建**

**node02：**
```bash
$> scp -r /opt/sxt/apache-hive-1.2.1-bin node03:`pwd`
$> scp -r /opt/sxt/apache-hive-1.2.1-bin node04:`pwd`
```
**node03（MetaStore Server端，主要负责与MySQL的元数据交互）:**
```xml
$> vim /etc/profile
    export HIVE_HOME=/opt/sxt/apache-hive-1.2.1-bin
    export PATH=$PATH:$HIVE_HOME/bin
$> source /etc/profile
$> cd /opt/sxt/apache-hive-1.2.1-bin/conf
$> vim hive-site.xml
    <property>  
        <name>hive.metastore.warehouse.dir</name>  
        <value>/user/hive/warehouse</value>  
    </property>  
    <property>  
        <name>javax.jdo.option.ConnectionURL</name>  
        <value>jdbc:mysql://node01:3306/hive?createDatabaseIfNotExist=true</value>  
    </property>  
    <property>  
        <name>javax.jdo.option.ConnectionDriverName</name>  
        <value>com.mysql.jdbc.Driver</value>  
    </property>  
    <property>  
        <name>javax.jdo.option.ConnectionUserName</name>  
        <value>root</value>  
    </property>  
    <property>  
        <name>javax.jdo.option.ConnectionPassword</name>  
        <value>123</value>  
    </property>
$> hive --service metastore   # 启动hive metastore服务（阻塞式运行）
```
**node04（Hive CLI端，主要负责与HDFS交互）:**
```xml
$> vim /etc/profile
    export HIVE_HOME=/opt/sxt/apache-hive-1.2.1-bin
    export PATH=$PATH:$HIVE_HOME/bin
$> source /etc/profile
$> cd /opt/sxt/apache-hive-1.2.1-bin/conf
$> vim hive-site.xml
    <property>  
        <name>hive.metastore.warehouse.dir</name>  
        <value>/user/hive/warehouse</value>
    </property>  
    <property>  
        <name>hive.metastore.local</name>  
        <value>false</value>  
    </property>  
    <property>  
        <name>hive.metastore.uris</name>   # 配置hive服务的地址
        <value>thrift://node03:9083</value>  
    </property>
$> rm -f /opt/sxt/hadoop-2.6.5/share/hadoop/yarn/lib/jline-0.9.94.jar
$> cp /opt/sxt/apache-hive-1.2.1-bin/lib/jline-2.12.jar /usr/local/hadoop/share/hadoop/yarn/lib
$> hive
```

# 3.Hive SQL
* **创建库**
```bash
CREATE (DATABASE|SCHEMA) [if not exists] database_name # 大写代表关键字，（）代表选择项，[]代表可选项，小写代表用户填写

hive> create database if not exists demo;
hive> show databases;
hive> use demo;
```
* **删除库**
```
hive> drop database if exists demo;
```
* **创建内部表**
``` bash
hive> show tables;
hive> create table psn0 (
        id int,
        name string,
        likes array<string>, # 数组类型
        address map<string, string> # 键值对类型
    )
hive> row format delimited # 设置数据的导入规则
hive> fields terminated by ',' # 字段划分依据
hive> collection items terminated by '-' # 数组内元素划分依据
hive> map keys terminated by ':' # map键值对划分依据
hive> lines terminated by '\n'; # 数据行划分依据
```
* **导入数据到表中**
```bash
hive> load data local inpath '/root/data_01' into table psn0; # 存在local关键字表示是linux文件系统的路径
hive> load data inpath '/input' into table psn0; # 没有local关键字表示是hdfs路径
```
* **创建外部表**
```bash
hive> create external table psn1 ( # 外部表是数据的引用，删除后不会影响数据
        id int,
        name string,
        likes array<string>,
        address map<string, string>
    )
row format delimited 
fields terminated by ','
collection items terminated by '-'
map keys terminated by ':'
location '/psn1';
```
* **创建分区表**
```bash
hive> create external table psn2 (
        id int,
        name string,
        likes array<string>,
        address map<string, string>
    )
hive> partitioned by(sex string, age int) # 设置分区字段，设置的字段数据中也要存在，会生成子目录按照分区规则划分数据
hive> row format delimited 
hive> fields terminated by ','
hive> collection items terminated by '-'
hive> map keys terminated by ':';
hive> load data local inpath '/root/data_02' into table psn2 partition(sex='man', age=10); # 指定分区字段的值
```
* **添加分区**
```bash
hive> alter table psn2 add partition(sex='boy', age=10);
```
* **删除分区**
```bash
hive> alter table psn2 drop partition(sex='boy', age=10);
```
* **复制并建新表**
```bash
hive> create table psn3 like psn2;
```
* **查询结果建新表**
```bash
hive> create table psn4 as select id, name, likes from psn2;
```
* **分析结果插入结果表**
```bash
hive> create table psn5(res int);
    from psn0
    insert into table psn5
    select count(*);
```
* **查询语句**
```bash
--查询组合字段 select distinct top(<top_specification>) <select_list> --连表 from <left_table><join_type> join <right_table> on <on_predicate> <left_table><apply_type> apply <right_table_expression> as <alias> <left_table> pivot (<pivot_specification>) as <alias> <left_table> unpivot (<unpivot_specification>) as <alias> --查询条件 where <where_pridicate> --分组
group by <group_by_specification> --分组条件
having <having_predicate> --排序 order by <sort col> (desc|) [limit num]
```

# 4.Hive SerDe
* **正则限定序列化反序列化规则**
```bash
create table logtbl ( 
    host string,
    identity string,
    t_user string,
    time string,
    request string,
    referer string,
    agent string
)
row format serde 'org.apache.hadoop.hive.serde2.RegexSerDe' # 调用序列化规则实现类
with serdeproperties ( # 向类传入参数
    "input.regex" = "([^]*) ([^]*) ([^]*) \\[(.*)\\] \"(.*)\" (-|[0-9]*) (-|[0-9]*)" # 设置数据导入的正则限定规则
)
stored as textfile;
load data local inpath '/root/log' into table logtbl; # 导入的数据会按照设置的正则进行清洗（注：导入不符合规则的数据不会报错，在查询时会报错，这是因为hive不是写时检查，而是读时检查机制）
select * from logtbl;
```

# 5.Hive Beeline
**node03：**
```bash
$> hiveserver2 # 在服务端启动hiveserver2服务
```
**node04：**
```bash
方式一：
    $> beeline -u jdbc:hive2://node03:10000 root # 进入hive beeline模式的命令交互模式
    0: jdbc:hive2://node03:10000> !quit
方式二：
    $> beeline
    beeline> !connect jdbc:hive2://node03:10000 root # 手动指定连接的服务节点，hive不对用户名密码做校验
```

# 6.Hive JDBC
* 开发所需jar包：```/opt/sxt/hive/lib```
* 在node3节点启动hiveserver2服务：```$> hiveserver2```
* jdbc连接服务器：```jdbc:hive2://node03:10000/sxt```

# 7.Hive函数
* **内置运算符**

1. **关系运算符**

| 运算符  | 类型 | 说明 |
| --- | --- | --- |
| A = B | 所有原始类型 | 如果A与B相等，返回TRUE，否则返回FALSE |
| A == B | 无 | 失败，因为无效的语法。 SQL使用”=”，不使用”==”。 |
| A <> B | 所有原始类型 | 如果A不等于B返回TRUE,否则返回FALSE。如果A或B值为”NULL”，结果返回”NULL”。 |
| A < B | 所有原始类型 | 如果A小于B返回TRUE,否则返回FALSE。如果A或B值为”NULL”，结果返回”NULL”。 |
| A <= B | 所有原始类型 | 如果A小于等于B返回TRUE,否则返回FALSE。如果A或B值为”NULL”，结果返回”NULL”。 |
| A > B | 所有原始类型 | 如果A大于B返回TRUE,否则返回FALSE。如果A或B值为”NULL”，结果返回”NULL”。 |
| A >= B | 所有原始类型 | 如果A大于等于B返回TRUE,否则返回FALSE。如果A或B值为”NULL”，结果返回”NULL”。 |
| A IS NULL | 所有类型 | 如果A值为”NULL”，返回TRUE,否则返回FALSE |
| A IS NOT NULL | 所有类型 | 如果A值不为”NULL”，返回TRUE,否则返回FALSE |
| A LIKE B | 字符串 | 如果A或B值为”NULL”，结果返回”NULL”。字符串A与B通过sql进行匹配，如果相符返回TRUE，不符返回FALSE。B字符串中 的”_”代表任一字符，”%”则代表多个任意字符。例如： (‘foobar’ like ‘foo’)返回FALSE，（ ‘foobar’ like ‘foo_ _ \_’或者 ‘foobar’ like ‘foo%’)则返回TURE  |
| A RLIKE B | 字符串 | 如果A或B值为”NULL”，结果返回”NULL”。字符串A与B通过java进行匹配，如果相符返回TRUE，不符返回FALSE。例如：（ ‘foobar’ rlike ‘foo’）返回FALSE，（’foobar’ rlike ‘^f.\*r$’ ）返回TRUE。 |
| A REGEXP B | 字符串 | 与RLIKE相同。 |

2. **算术运算符**

| 运算符 | 类型 | 说明 |
| --- | --- | --- |
| A + B | 所有数字类型 | A和B相加。结果的与操作数值有共同类型。例如每一个整数是一个浮点数，浮点数包含整数。所以，一个浮点数和一个整数相加结果也是一个浮点数。|
| A – B | 所有数字类型 | A和B相减。结果的与操作数值有共同类型。|
| A * B | 所有数字类型 | A和B相乘，结果的与操作数值有共同类型。需要说明的是，如果乘法造成溢出，将选择更高的类型。|
| A / B | 所有数字类型 | A和B相除，结果是一个double（双精度）类型的结果。|
| A % B | 所有数字类型 | A除以B余数与操作数值有共同类型。|
| A & B | 所有数字类型 | 运算符查看两个参数的二进制表示法的值，并执行按位"与"操作。两个表达式的一位均为1时，则结果的该位为1。否则，结果的该位为0。|
| A\|B | 所有数字类型 | 运算符查看两个参数的二进制表示法的值，并执行按位"或"操作。只要任一表达式的一位为 1，则结果的该位为1。否则，结果的该位为0。|
| A ^ B | 所有数字类型 | 运算符查看两个参数的二进制表示法的值，并执行按位"异或"操作。当且仅当只有一个表达式的某位上为1时，结果的该位才为1。否则结果的该位为0。|
| ~A | 所有数字类型 | 对一个表达式执行按位"非"（取反）。 |

3. **逻辑运算符**

| 运算符 | 类型 | 说明 |
| --- | --- | --- |
| A AND B | 布尔值 | A和B同时正确时,返回TRUE,否则FALSE。如果A或B值为NULL，返回NULL。 |
| A && B | 布尔值 | 与"A AND B"相同 |
| A OR B | 布尔值 | A或B正确,或两者同时正确返返回TRUE,否则FALSE。如果A和B值同时为NULL，返回NULL。 |
| A | B | 布尔值 | 布尔值	与”A OR B”相同 |
| NOT A | 布尔值 | 布尔值	如果A为NULL或错误的时候返回TURE，否则返回FALSE。 |
| ! A | 布尔值 | 与”NOT A”相同 |

4. **复杂类型函数**

| 函数 | 类型 | 说明 |
| --- | --- | --- |
| map | (key1, value1, key2, value2, …) | 通过指定的键/值对，创建一个map。 |
| struct | (val1, val2, val3, …) | 通过指定的字段值，创建一个结构。结构字段名称将COL1，COL2，… |
| array | (val1, val2, …) | 通过指定的元素，创建一个数组。 |

5. **对复杂类型函数操作**

| 函数 | 类型 | 说明 |
| --- | --- | --- |
| A[n] | A是一个数组，n为int型 | 返回数组A的第n个元素，第一个元素的索引为0。如果A数组为['foo','bar']，则A[0]返回’foo’和A[1]返回”bar”。 |
| M[key] | M是Map<K, V>，关键K型 | 返回关键值对应的值，例如mapM为 \{‘f’ -> ‘foo’, ‘b’ -> ‘bar’, ‘all’ -> ‘foobar’\}，则M['all'] 返回’foobar’。|
| S.x | S为struct | 返回结构x字符串在结构S中的存储位置。如 foobar \{int foo, int bar\} foobar.foo的领域中存储的整数。|

* 内置函数（UDF）：

1. **数学函数**

| 返回类型 | 函数 | 说明 |
| --- | --- | --- |
| BIGINT | round(double a) | 四舍五入 |
| DOUBLE | round(double a, int d) | 小数部分d位之后数字四舍五入，例如round(21.263,2),返回21.26 |
| BIGINT | floor(double a) | 对给定数据进行向下舍入最接近的整数。例如floor(21.2),返回21。 |
| BIGINT | ceil(double a), ceiling(double a) | 将参数向上舍入为最接近的整数。例如ceil(21.2)，返回23。|
| double | rand(), rand(int seed) | 返回大于或等于0且小于1的平均分布随机数（依重新计算而变）|
| double | exp(double a) | 返回e的n次方 |
| double | ln(double a) | 返回给定数值的自然对数 |
| double | log10(double a) | 返回给定数值的以10为底自然对数 |
| double | log2(double a) | 返回给定数值的以2为底自然对数 |
| double | log(double base, double a) | 返回给定底数及指数返回自然对数 |
| double | pow(double a, double p) power(double a, double p) | 返回某数的乘幂 |
| double | sqrt(double a) | 返回数值的平方根 |
| string | bin(BIGINT a) | 返回二进制格式 |
| string | hex(BIGINT a) hex(string a) | 将整数或字符转换为十六进制格式 |
| string | unhex(string a) | 十六进制字符转换由数字表示的字符。 |
| string | conv(BIGINT num, int from_base, int to_base) | 将 指定数值，由原来的度量体系转换为指定的试题体系。例如CONV(‘a’,16,2),返回。参考：’1010′ http://dev.mysql.com/doc/refman/5.0/en/mathematical-functions.html#function_conv  |
| double | abs(double a) | 取绝对值 |
| int double | pmod(int a, int b) pmod(double a, double b) | 返回a除b的余数的绝对值 |
| double | sin(double a) | 返回给定角度的正弦值 |
| double | asin(double a) | 返回x的反正弦，即是X。如果X是在-1到1的正弦值，返回NULL。 |
| double | cos(double a) | 返回余弦 |
| double | acos(double a) | 返回X的反余弦，即余弦是X，，如果-1<= A <= 1，否则返回null. |
| int double | positive(int a) positive(double a) | 返回A的值，例如positive(2)，返回2。 |
| int double | negative(int a) negative(double a) | 返回A的相反数，例如negative(2),返回-2。|

2. **收集函数**

| 返回类型 | 函数 | 说明 |
| --- | --- | --- |
| int | size(Map<K.V>) | 返回的map类型的元素的数量 |
| int | size(Array<T>) | 返回数组类型的元素数量 |

3. **类型转换函数**

| 返回类型 | 函数 | 说明 |
| --- | --- | --- |
| 指定"type" | cast(expr as <type>) | 类型转换。例如将字符”1″转换为整数: cast(’1′ as bigint)，如果转换失败返回NULL。 |

4. **日期函数**

| 返回类型 | 函数 | 说明 |
| --- | --- | --- |
| string | from_unixtime(bigint unixtime[, string format]) | UNIX_TIMESTAMP参数表示返回一个值’YYYY- MM – DD HH：MM：SS’或YYYYMMDDHHMMSS.uuuuuu格式，这取决于是否是在一个字符串或数字语境中使用的功能。该值表示在当前的时区。|
| bigint | unix_timestamp() | 如果不带参数的调用，返回一个Unix时间戳（从’1970- 01 – 0100:00:00′到现在的UTC秒数）为无符号整数。|
| bigint | unix_timestamp(string date) | 指定日期参数调用UNIX_TIMESTAMP（），它返回参数值’1970- 01 – 0100:00:00′到指定日期的秒数。|
| bigint | unix_timestamp(string date, string pattern) | 指定时间输入格式，返回到1970年秒数：unix_timestamp(’2009-03-20′, ‘yyyy-MM-dd’) = 1237532400 |
| string | to_date(string timestamp) | 返回时间中的年月日： to_date(“1970-01-01 00:00:00″) = “1970-01-01″ |
| string | to_dates(string date) | 给定一个日期date，返回一个天数（0年以来的天数）|
| int | year(string date) | 返回指定时间的年份，范围在1000到9999，或为”零”日期的0。 |
| int | month(string date) | 返回指定时间的月份，范围为1至12月，或0一个月的一部分，如’0000-00-00′或’2008-00-00′的日期。|
| int | day(string date) dayofmonth(date) | 返回指定时间的日期 |
| int | hour(string date) | 返回指定时间的小时，范围为0到23。|
| int | minute(string date) | 返回指定时间的分钟，范围为0到59。|
| int | second(string date) | 返回指定时间的秒，范围为0到59。|
| int | weekofyear(string date) | 返回指定日期所在一年中的星期号，范围为0到53。|
| int | datediff(string enddate, string startdate) | 两个时间参数的日期之差。|
| int | date_add(string startdate, int days) | 给定时间，在此基础上加上指定的时间段。|
| int | date_sub(string startdate, int days) | 给定时间，在此基础上减去指定的时间段。|

5. **条件函数**

| 返回类型 | 函数 | 说明 |
| --- | --- | --- |
| T | if(boolean testCondition, T valueTrue, T valueFalseOrNull) | 判断是否满足条件，如果满足返回一个值，如果不满足则返回另一个值。|
| T | COALESCE(T v1, T v2, …) | 返回一组数据中，第一个不为NULL的值，如果均为NULL,返回NULL。|
| T | CASE a WHEN b THEN c [WHEN d THEN e]* [ELSE f] END | 当a=b时,返回c；当a=d时，返回e，否则返回f。|
| T | CASE WHEN a THEN b [WHEN c THEN d]* [ELSE e] END | 当值为a时返回b,当值为c时返回d。否则返回e。|

6. **字符函数**

| 返回类型 | 函数 | 说明 |
| --- | --- | --- |
| int | length(string A) | 返回字符串的长度 |
| string | reverse(string A) | 返回倒序字符串 |
| string | concat(string A, string B…) | 连接多个字符串，合并为一个字符串，可以接受任意数量的输入字符串 |
| string | concat_ws(string SEP, string A, string B…) | 链接多个字符串，字符串之间以指定的分隔符分开 |
| string | substr(string A, int start) substring(string A, int start) | 从文本字符串中指定的起始位置后的字符 |
| string | substr(string A, int start, int len) substring(string A, int start, int len) | 从文本字符串中指定的位置指定长度的字符 |
| string | upper(string A) ucase(string A) | 将文本字符串转换成字母全部大写形式 |
| string | lower(string A) lcase(string A) | 将文本字符串转换成字母全部小写形式 |
| string | trim(string A) | 删除字符串两端的空格，字符之间的空格保留 |
| string | ltrim(string A) | 删除字符串左边的空格，其他的空格保留 |
| string | rtrim(string A) | 删除字符串右边的空格，其他的空格保留 |
| string | regexp_replace(string A, string B, string C) | 字符串A中的B字符被C字符替代 |
| string | regexp_extract(string subject, string pattern, int index) | 通过下标返回正则表达式指定的部分。regexp_extract(‘foothebar’, ‘foo(.\*?)(bar)’, 2) returns ‘bar.’ |
| string | parse_url(string urlString, string partToExtract [, string keyToExtract]) | 返回URL指定的部分。parse_url(‘http://facebook.com/path1/p.php?k1=v1&k2=v2#Ref1′, ‘HOST’) 返回：’facebook.com’ |
| string | get_json_object(string json_string, string path) | select a.timestamp, get_json_object(a.appevents, ‘$.eventid’), get_json_object(a.appenvets, ‘$.eventname’) from log a; |
| string | space(int n) | 返回指定数量的空格 |
| string | repeat(string str, int n) | 重复N次字符串 |
| int | ascii(string str) | 返回字符串中首字符的数字值 |
| string | lpad(string str, int len, string pad) | 返回指定长度的字符串，给定字符串长度小于指定长度时，由指定字符从左侧填补 |
| string | rpad(string str, int len, string pad) | 返回指定长度的字符串，给定字符串长度小于指定长度时，由指定字符从右侧填补 |
| array | split(string str, string pat) | 将字符串转换为数组 |
| int | find_in_set(string str, string strList) | 返回字符串str第一次在strlist出现的位置。如果任一参数为NULL,返回NULL；如果第一个参数包含逗号，返回0 |
| array<array<string>> | sentences(string str, string lang, string locale) | 将字符串中内容按语句分组，每个单词间以逗号分隔，最后返回数组。 例如sentences(‘Hello there! How are you?’) 返回：( (“Hello”, “there”), (“How”, “are”, “you”)) |
| array<struct<string,double>> | ngrams(array<array<string>>, int N, int K, int pf) | SELECT ngrams(sentences(lower(tweet)), 2, 100 [, 1000]) FROM twitter; |
| array<struct<string,double>> | context_ngrams(array<array<string>>, array<string>, int K, int pf) | SELECT context_ngrams(sentences(lower(tweet)), array(null,null), 100, [, 1000]) FROM twitter; |

* **内置聚合函数（UDAF）**

| 返回类型 | 函数 | 说明 |
| --- | --- | --- |
| bigint | count(\*) , count(expr), count(DISTINCT expr[, expr_., expr_.]) | 返回记录条数 |
| double | sum(col), sum(DISTINCT col) | 求和 |
| double | avg(col), avg(DISTINCT col) | 求平均值 |
| double | min(col) | 返回指定列中最小值 |
| double | max(col) | 返回指定列中最大值 |
| double | var_pop(col) | 返回指定列的方差 |
| double | var_samp(col) | 返回指定列的样本方差 |
| double | stddev_pop(col) | 返回指定列的偏差 |
| double | stddev_samp(col) | 返回指定列的样本偏差 |
| double | covar_pop(col1, col2) | 两列数值协方差 |
| double | covar_samp(col1, col2) | 两列数值样本协方差 |
| double | corr(col1, col2) | 返回两列数值的相关系数 |
| double | percentile(col, p) | 返回数值区域的百分比数值点。0<=P<=1,否则返回NULL,不支持浮点型数值 |
| array<double> | percentile(col, array(p~1,,\ [, p,,2,,]…)) | 返回数值区域的一组百分比值分别对应的数值点。0<=P<=1,否则返回NULL,不支持浮点型数值 |
| double | percentile_approx(col, p[, B]) | Returns an approximate p^th^ percentile of a numeric column (including floating point types) in the group. The B parameter controls approximation accuracy at the cost of memory. Higher values yield better approximations, and the default is 10,000. When the number of distinct values in col is smaller than B, this gives an exact percentile value. |
| array<double> | percentile_approx(col, array(p~1,, [, p,,2_]…) [, B]) | Same as above, but accepts and returns an array of percentile values instead of a single one. |
| array<struct\{‘x’,'y’\}> | histogram_numeric(col, b) | Computes a histogram of a numeric column in the group using b non-uniformly spaced bins. The output is an array of size b of double-valued (x,y) coordinates that represent the bin centers and heights |
| array | collect_set(col) | 返回无重复记录 |

* **内置表生成函数（UDTF）**

| 返回类型 | 函数 | 说明 |
| --- | --- | --- |
| 数组 | explode(array<TYPE> a) | 数组一条记录中有多个参数，将参数拆分，每个参数生成一列 |
|  | json_tuple | get_json_object 语句：select a.timestamp, get_json_object(a.appevents, ‘$.eventid’), get_json_object(a.appenvets, ‘$.eventname’) from log a; json_tuple语句: select a.timestamp, b.* from log a lateral view json_tuple(a.appevent, ‘eventid’, ‘eventname’) b as f1, f2 |

* **自定义UDF函数**
```bash
hive> add jar /root/tuomin.jar; # 添加jar包
hive> create temporary function tm as 'com.sxt.hive.TuoMin'; # 指定jar的入口函数创建hive函数
hive> select id, name, tm(id), tm(name) from psn1; # 测试自定义函数
```

# 8.Hive基站掉话率案例
* **分析目标：** 找出掉线率最高的前10基站
* **数据格式展示：**

**cdr_summ_imei_cell_info.csv文件**

| record_time（通话时间） | imei（基站编号） | cell（手机编号） | ph_num | call_num | drop_num（掉话时间/秒） | duration（通话持续时间/秒） | drop_rate | net_type | erl |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 2011-07-13 | 00:00:00+08 | 356966 | 29448-37062 | 0 | 0 | 0 | 0 | 0 | 0 |
| 2011-07-13 | 00:00:00+08 | 352024 | 29448-51331 | 0 | 0 | 0 | 0 | 0 | 0 |
| 2011-07-13 | 00:00:00+08 | 353736 | 29448-51331 | 0 | 0 | 0 | 0 | 0 | 0 |

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

# 9.Hive WordCount案例

* **分析目标**：统计所有单词出现的次数
* **数据格式展示**：

**wc文件**
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

# 10.Hive参数：
* 参数设置方式：
    * 修改配置文件 ```${HIVE_HOME}/conf/hive-site.xml```
    * 启动hive cli时，通过--hiveconf key=value的方式设置
```bash
$> hive --hiveconf hive.cli.print.header=true
```
    * 进入hive cli后，通过set命令设置
```bash
hive> set hive.cli.print.header=true; # set设置
hive> set hive.cli.print.header # set查看
$> vi ~/.hiverc
    hive.cli.print.header=true; # 每次启动hive cli都会读取文件中的配置项
$> tail 10 ~/.hivehistory # hive历史操作命令集
```

# 11.Hive动态分区：
* **动态分区参数设置**
```bash
hive> set hive.exec.dynamic.partition=true; # 开启动态分区支持
hive> set hive.exec.dynamic.partition.mode=nostrict; 
hive> set hive.exec.max.dynamic.partitions.pernode; # 每一个执行mr的节点上，允许创建动态分区的最大数量
hive> set hive.exec.max.dynamic.partitions; # 每一个执行mr的节点上，允许创建的所有动态分区的最大数量
hive> set hive.exec.max.created.files; # 所有mr job允许创建文件的最大数量
```

* **创建动态分区的表**
```bash
hive> create table psn6( # 分区表指定分区字段
        id int,
        name string,
        likes array<string>,
        address map<string, string>)
    partitioned by (sex string, age int)
    row format delimited
    fields terminated by ','
    collection items terminated by '-'
    map keys terminated by ':';
hive> create table psn7( # 常规表用于数据的导入
        id int,
        name string,
        sex string,
        age int,
        likes array<string>,
        address map<string, string> )
    row format delimited
    fields terminated by ','
    collection items terminated by '-'
    map keys terminated by ':';
hive> load data local inpath '/root/data' into table psn7;
hive> from psn7
     insert overwrite table psn6 partition(sex, age)
     select id, name, likes, address, sex, age
     distribute by sex, age;
```

# 12.Hive分桶
* **分桶参数设置**
```bash
hive> set hive.enforce.bucketing=true; # 开启分桶的支持
注：mr运行时会根据bucket的个数自动分配reduce task的个数
```
* **桶表抽样查询语法**
```bash
select * from bucket_table tablesample(bucket x out of y);
tablesample(bucket x out of y)语法：
    - x：表示从哪个bucket开始抽取数据
    - y：为该表总bucket数的倍数或因子
```
* **例：总bucket数为32时**
	tablesample(bucket 2 out of 16)共抽取2(32/16)个bucket的数据，抽取的是第2个bucket，第2+16=18个bucket的数据
* **例：总bucket数为32时**
	tablesample(bucket 3 out of 256)共抽取1/8(32/256)个bucket的数据，抽取的是第3个bucket的1/8的数据
* **桶表抽样查询**
```bash
hive> create table psn8( # 创建原始数据表
        id int,
        name string,
        age int )
     row format delimited
     fields terminated by ',';
hive> load data local inpath '/root/data_05.txt' into table psn8; # 导入数据到数据表
hive> create table psnbucket( # 创建桶表
        id int,
        name string,
        age int )
     clustered by(age) into 4 buckets # 设置按照age列的哈希对桶数取模来分桶
     row format delimited
     fields terminated by ',';
hive> from psn8 # 插入数据到桶表
     insert into table psnbucket
     select id, name, age;
hive> select id, name, age # 抽样查询数据
     from psnbucket 
     tablesample(bucket 2 out of 4 on age); # 抽样取2号桶的全部数据
```

# 13.Hive Lateral View
* **Lateral View**
```bash
与UDTF函数结合使用，通过UDTF函数将数据拆分成多行，再将多行结果组合成一个支持别名的虚拟表，主要解决在select中使用UDTF做查询过程中，查询只能包含单个UDTF，不能包含其他字段其他UDTF的问题
语法：lateral view udtf(expression) tableAlias as columnAlias(',', columnAlias)
```
* **使用Lateral View**
```bash
统计人员表中公有多少种爱好，多少个城市：
hive> select count(distinct(likeCol)), count(distinct(addressCol_01)) from psn2
    lateral view explode(likes) myTable_01 as likeCol # 使用UDTF后建立虚拟表myTable_01,as指定结果的字段
    lateral view explode(address) myTable_02 as addressCol_01, addressCol_02;
```

# 14.Hive视图
* **创建视图**
```sql
create view view_psn1 as 
select likes[0] from psn1;
```
* **查询视图**
```sql
select * from view_psn1;
```
* **删除视图**
```sql
drop view view_psn1;
```

# 15.Hive索引
* **创建索引**
```bash
hive> create index index_psn1 on table psn1(name) # 指定针对某表的某字段创建索引
    as 'org.apache.hadoop.hive.ql.index.compact.CompactIndexHandler' with deferred rebuild # 指定索引器（索引实现类）
    in table index_psn1_table; # 指定存入索引表
hive> create index index2_psn1 on table psn1(id)
    as 'org.apache.hadoop.hive.ql.index.compact.CompactIndexHandler' with deferred rebuild; # 若不指定索引表，默认生成sxt_psn1_index_psn1_表中
```
* **查询索引**
```sql
show index on psn1;
```
* **重建索引**
```sql
alter index index_psn1 on psn1 rebuild; # 建立索引后必须重建索引才能生效
```
* **删除索引**
```sql
drop index if exists index_psn1 on psn1;
```

# 16.Hive运行方式
* **命令行模式**
```bash
CLI：
    node03：$> hive --service metastore
    node04：$> hive
beeline：
    node03：$> hiveserver2 
    node04：$> beeline -u jdbc:hive2://node03:10000 root
与Linux交互：
    hive> !pwd;
与hdfs交互：
    hive> dfs -ls /user/hive;
```
* **脚本运行方式**
```bash
$> hive -e "select * from psn1 limit 3"
$> hive -e "select * from psn2" > test
$> hive -e "select * from psn3" >> test
$> hive -S -e "select * from psn1"
$> vim hivesql
    select * from psn1;
    select * from psn2 limit 10;
$> hive -f hivesql
$> hive -i hivesql
hive> source /root/hivesql; 
```
* **JDBC方式（hiveserver2）**
```bash
node03：
$> hiveserver2
client： 
jdbc:hive2://node03:10000/sxt
```
* **Web GUI接口**
```bash
下载apache-hive-*-src.tar.gz源码包
dos> cd apache-hive-1.2.1-src/hwi/web # 跳转到源码包的web目录下
dos> jar -cvf hive-hwi.war * # 封装war包
$> mv /root/hive-hwi.war /opt/sxt/hive/lib 
$> cp /usr/java/jdk/lib/tools.jar /opt/sxt/hive/lib
$> vi /opt/sxt/hive/conf/hive-site.xml # 修改hive配置文件
    <property>
        <name>hive.hwi.listen.host</name>
        <value>0.0.0.0</value>
    </property>
    <property>
        <name>hive.hwi.listen.port</name>
        <value>9999</value>
    </property>
    <property>
        <name>hive.hwi.war.file</name>
        <value>lib/hive-hwi.war</value>
    </property>
node03：
    $> hive --service hwi # hive服务端启动hwi服务
http://node03:9999/hwi # 浏览器访问web页面
```

# 17.Hive权限管理
* **基于sql的权限控制**
```xml
Hive - SQL Standards Based Authorization in HiveServer2
$> vim /opt/sxt/hive/conf/hive-site.xml
    <property>
        <name>hive.security.authorization.enabled</name>
        <value>true</value>
    </property>
    <property>
        <name>hive.server2.enable.doAs</name>
        <value>false</value>
    </property>
    <property>
        <name>hive.users.in.admin.role</name>
        <value>root</value>
    </property>
    <property>
        <name>hive.security.authorization.manager</name> 
        <value>org.apache.hadoop.hive.ql.security.authorization.plugin.sqlstd.SQLStdHiveAuthorizerFactory</value>
    </property>
    <property>
        <name>hive.security.authenticator.manager</name>
        <value>org.apache.hadoop.hive.ql.security.SessionStateUserAuthenticator</value>
    </property>
服务端启动hiveserver2,客户端通过beeline进行连接
```
* **角色的添加、删除、查看、设置**
```sql
create role role_name;  -- 创建角色
drop role role_name;  -- 删除角色
set role (role_name|ALL|NONE);  -- 设置角色
show current roles;  -- 查看当前具有的角色
show roles;  -- 查看所有存在的角色
```

# 18.Hive优化
* **显示执行计划**
```bash
hive> explain [extended] select * from psn1;
```

* **本地模式**
```bash
hive> set hive.exec.mode.local.auto=true;
注：hive.exec.mode.local.auto.inputbytes.max # 默认为128M，表示加载文件的最大值，若大于该配置仍会以集群方式来运行
```

* **并行计算**
```bash
hive> set hive.exec.parallel=true;
注：hive.exec.parallel.thread.number # 一次sql计算中允许并行执行的job最大值
```

* **严格模式**
```bash
hive> set hive.mapred.mode=strict;
查询限制：
1.对于分区表，必须添加where对于分区字段的条件过滤
2.order by语句必须包含limit输出限制
3.限制执行笛卡尔积的查询
```

* **hive排序**
```
Order By - 对于查询结果做全排序，只允许有一个reduce处理
Sort By - 对于单个reduce的数据进行排序
Distribute By - 分区排序，经常和Sort By结合使用
Cluster By - 相当于Sort By + Distribute By（Cluster By不能通过asc、desc的方式指定排序规则，可通过 distribute by column sort by column asc|desc 的方式）
```

* **hive join**
```
join计算时，将小表（驱动表）放在join的左边
Map Join：在Map端完成Join
```

* **两种实现方式**
1. SQL方式（在SQL语句中添加MapJoin标记）
```sql
select /*+ mapjoin(smallTable) */ smallTable.key, bigTable.value
from smallTable join bigTable on smallTable.key = bigTable.key;
```
2. 开启自动的MapJoin
```
set hive.auto.convert.join = true;（该参数为true时，Hive自动对左边的表统计量，如果是小表就加入内存，即对小表使用Map join）
hive.mapjoin.smalltable.filesize;（大表小表判断的阈值，如果表的大小小于该值则会被加载到内存中运行）
hive.ignore.mapjoin.hint;（默认值：true；是否忽略mapjoin hint 即mapjoin标记）
hive.auto.convert.join.noconditionaltask;（默认值：true；将普通的join转化为普通的mapjoin时，是否将多个mapjoin转化为一个mapjoin）
hive.auto.convert.join.noconditionaltask.size;（将多个mapjoin转化为一个mapjoin时，其表的最大值）
```
* Map-Side聚合：
```
set hive.map.aggr=true; # 开启在Map端的聚合
hive.groupby.mapaggr.checkinterval; # map端group by执行聚合时处理的多少行数据（默认：100000）
hive.map.aggr.hash.min.reduction; # 进行聚合的最小比例（预先对100000条数据做聚合，若聚合之后的数据量/100000的值大于该配置0.5，则不会聚合）
hive.map.aggr.hash.percentmemory; # map端聚合使用的内存的最大值
hive.map.aggr.hash.force.flush.memory.threshold; # map端做聚合操作是hash表的最大可用内容，大于该值则会触发flush
hive.groupby.skewindata; # 是否对GroupBy产生的数据倾斜做优化，默认为false
```

* 控制Hive中Map以及Reduce的数量：
    * Map数量相关的参数：
        * mapred.max.split.size # 一个split的最大值，即每个map处理文件的最大值
        * mapred.min.split.size.per.node # 一个节点上split的最小值
        * mapred.min.split.size.per.rack # 一个机架上split的最小值
    * Reduce数量相关的参数：
        * mapred.reduce.tasks # 强制指定reduce任务的数量
        * hive.exec.reducers.bytes.per.reducer # 每个reduce任务处理的数据量
        * hive.exec.reducers.max # 每个任务最大的reduce数
* JVM重用：
    * 适用场景：
        * 小文件个数过多
        * task个数过多
```
hive> set mapred.job.reuse.jvm.num.tasks=n; # 设置task插槽个数
注：设置开启之后，task插槽会一直占用资源，不论是否有task运行，直到所有的task即整个job全部执行完成时，才会释放所有的task插槽资源
```