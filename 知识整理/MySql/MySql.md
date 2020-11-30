[TOC]

# 1.数据库操作

## 1.1 显示数据库
```sql
mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| performance_schema |
| sys                |
+--------------------+
4 rows in set (0.00 sec)
```

* MySQL 自带的4个数据库：
    * **information_schema：** 存储了 mysql 服务器管理数据库的信息；
    * **performance_schema：** mysql5.5 新增的表，用来保存数据库服务器性能的参数；
    * **mysql：** mysql 系统数据库，保存的登录用户名，密码，以及每个用户的权限等等；
    * **test：** 给用户学习和测试的数据库。

## 1.2 创建数据库
* 语法：
```sql
create database [if not exists] `数据库名` [charset=字符编码];
```
* 直接创建数据库：
```sql
create database `stu`;
Query OK, 1 row affected (0.05 sec)
```
* 创建时添加判断条件：
```sql
create database if not exists `stu`;
Query OK, 1 row affected, 1 warning (0.00 sec)
```
**注**：创建时为数据库名添加反引号，以确保不受其他条件影响。

* 创建时指定字符编码：
```sql
create database `teacher` charset=utf8;
Query OK, 1 row affected, 1 warning (0.10 sec)
```

## 1.3 删除数据库
* 语法：
```sql
drop database [if exists] `数据库名`;
```
* 直接删除数据库：
```sql
drop database `stu`;
Query OK, 0 rows affected (0.19 sec)
```
* 删除时添加判断条件：
```sql
drop database if exists `teacher`;
Query OK, 0 rows affected (0.00 sec)
```

## 1.4 显示创建数据库时的 SQL 语句
* 语法：
```sql
show create database `数据库名`;
```
* 显示创建数据库的SQL语句：
```sql
create database if not exists `person` charset=utf8;
Query OK, 1 row affected, 1 warning (0.09 sec)

show create database `person`;
+----------+-----------------------------------------------------------------+
| Database | Create Database                                                 |
+----------+-----------------------------------------------------------------+
| person   | CREATE DATABASE `person` /*!40100 DEFAULT CHARACTER SET utf8 */ |
+----------+-----------------------------------------------------------------+
1 row in set (0.00 sec)
```

## 1.5 修改数据库
* 语法：
```sql
alter database `数据库名` charset=字符编码
```
* 修改数据库的字符编码：
```sql
alter database `person` charset=gbk;
Query OK, 1 row affected (0.00 sec)

show create database `person`;
+----------+----------------------------------------------------------------+
| Database | Create Database                                                |
+----------+----------------------------------------------------------------+
| person   | CREATE DATABASE `person` /*!40100 DEFAULT CHARACTER SET gbk */ |
+----------+----------------------------------------------------------------+
1 row in set (0.00 sec)
```

## 1.6 选择数据库
* 语法：
```sql
use `数据库名`
```
* 切换到其他数据库：
```sql
use `student`;
Database changed
```

# 2.表操作

## 2.1 显示所有表
* 语法：
```sql
show tables;
```
* 例：
```sql
show tables;
```

## 2.2 创建数据表
* 语法：
```sql
create table [if not exists] `表名`(
    `字段名` 数据类型 [null|not null] [auto_increment] [primary key] [comment],
    `字段名` 数据类型 [default]...
)engine=存储引擎
```
* 选项参数：

|选项|说明|
|:-:|:-:|
|null\|not null|空\|非空|
|default|默认值|
|auto_increment|自动增长|
|primary key|主键|
|comment|备注|
|engine|数据库引擎（innodb, myisam, memory）|

* 创建简单的表：
```sql
create database `itcast`;
Query OK, 1 row affected (0.05 sec)

use `itcast`;
Database changed

show tables;
Empty set (0.00 sec)

create table `stu`(
    `id` int,
    `name` varchar(30)
);
Query OK, 0 rows affected (0.04 sec)

show tables;
+------------------+
| Tables_in_itcast |
+------------------+
| stu              |
+------------------+
1 row in set (0.00 sec)
```
* 创建复杂的表：
```sql
create table if not exists `teacher`(
    `id` int not null auto_increment primary key comment 'primary key',
    `name` varchar(20) null comment 'name',
    `phone` varchar(20) comment 'phone',
    `address` varchar(100) default 'empty' comment 'address'
)engine=innodb;
Query OK, 0 rows affected (2.01 sec)
```
* 为其他数据库创建表：
```sql
show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| itcast             |
| mysql              |
| performance_schema |
| person             |
| student            |
| sys                |
+--------------------+
7 rows in set (0.00 sec)

use `itcast`;
Database changed

create table person.stu(
    id int,
    name varchar(20)
);
Query OK, 0 rows affected (1.83 sec)
```

## 2.3 显示创建表时的语句
* 语法：
```sql
show create table `表名`
```
* 显示创建表的SQL（将显示结果的两个字段纵向排列）：
```sql
show create table `teacher`\G;
*************************** 1. row ***************************
       Table: teacher
Create Table: CREATE TABLE `teacher` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'primary key id',
  `name` varchar(20) DEFAULT NULL COMMENT 'name',
  `phone` varchar(20) DEFAULT NULL COMMENT 'phone',
  `address` varchar(100) DEFAULT 'empty' COMMENT 'address',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
1 row in set (0.01 sec)
```

## 2.4 查看表结构
* 语法：
```sql
desc[ribe] `表名`
```
* 格式化显示表的结构：
```sql
describe `teacher`;
+---------+--------------+------+-----+---------+----------------+
| Field   | Type         | Null | Key | Default | Extra          |
+---------+--------------+------+-----+---------+----------------+
| id      | int(11)      | NO   | PRI | NULL    | auto_increment |
| name    | varchar(20)  | YES  |     | NULL    |                |
| phone   | varchar(20)  | YES  |     | NULL    |                |
| address | varchar(100) | YES  |     | empty   |                |
+---------+--------------+------+-----+---------+----------------+
4 rows in set (0.00 sec)

desc `teacher`;
+---------+--------------+------+-----+---------+----------------+
| Field   | Type         | Null | Key | Default | Extra          |
+---------+--------------+------+-----+---------+----------------+
| id      | int(11)      | NO   | PRI | NULL    | auto_increment |
| name    | varchar(20)  | YES  |     | NULL    |                |
| phone   | varchar(20)  | YES  |     | NULL    |                |
| address | varchar(100) | YES  |     | empty   |                |
+---------+--------------+------+-----+---------+----------------+
4 rows in set (0.00 sec)
```

## 2.5 删除表
* 语法：
```sql
drop table [if exists] `表1`,`表2`,...
```
* 删除多个表表：
```sql
drop table `stu`,`teacher`;
Query OK, 0 rows affected (0.02 sec)
```
* 删除表时添加判断条件：
```sql
drop table if exists `stu`;
Query OK, 0 rows affected, 1 warning (0.00 sec)
```

## 2.6 修改表
* 语法：
```sql
alter table `表名`
```
* 添加字段：
```sql
# 语法：
alter table `表名` add [column] 字段名 数据类型 [first|after]

# 添加字段到最后：
create table `person`(
    `id` int,
    `name` varchar(20)
)
Query OK, 0 rows affected (0.08 sec)

alter table `person` add age int;
Query OK, 0 rows affected (1.05 sec)
Records: 0  Duplicates: 0  Warnings: 0

describe `person`;
+-------+-------------+------+-----+---------+-------+
| Field | Type        | Null | Key | Default | Extra |
+-------+-------------+------+-----+---------+-------+
| id    | int(11)     | YES  |     | NULL    |       |
| name  | varchar(20) | YES  |     | NULL    |       |
| age   | int(11)     | YES  |     | NULL    |       |
+-------+-------------+------+-----+---------+-------+
3 rows in set (0.00 sec)

# 添加到开头字段：
alter table `person` add sex varchar(2) first;
Query OK, 0 rows affected (0.07 sec)
Records: 0  Duplicates: 0  Warnings: 0

describe `person`;
 +-------+-------------+------+-----+---------+-------+
| Field | Type        | Null | Key | Default | Extra |
+-------+-------------+------+-----+---------+-------+
| sex   | varchar(2)  | YES  |     | NULL    |       |
| id    | int(11)     | YES  |     | NULL    |       |
| name  | varchar(20) | YES  |     | NULL    |       |
| age   | int(11)     | YES  |     | NULL    |       |
+-------+-------------+------+-----+---------+-------+
4 rows in set (0.00 sec)

# 添加到指定字段的后面：
alter table `person` add email varchar(10) after name;
Query OK, 0 rows affected (0.09 sec)
Records: 0  Duplicates: 0  Warnings: 0

describe `person`;
+-------+-------------+------+-----+---------+-------+
| Field | Type        | Null | Key | Default | Extra |
+-------+-------------+------+-----+---------+-------+
| sex   | varchar(2)  | YES  |     | NULL    |       |
| id    | int(11)     | YES  |     | NULL    |       |
| name  | varchar(20) | YES  |     | NULL    |       |
| email | varchar(10) | YES  |     | NULL    |       |
| age   | int(11)     | YES  |     | NULL    |       |
+-------+-------------+------+-----+---------+-------+
5 rows in set (0.00 sec)
```
* 删除字段：
```sql
alter table `person` drop email;
Query OK, 0 rows affected (0.06 sec)
Records: 0  Duplicates: 0  Warnings: 0

describe `person`;
+-------+-------------+------+-----+---------+-------+
| Field | Type        | Null | Key | Default | Extra |
+-------+-------------+------+-----+---------+-------+
| sex   | varchar(2)  | YES  |     | NULL    |       |
| id    | int(11)     | YES  |     | NULL    |       |
| name  | varchar(20) | YES  |     | NULL    |       |
| age   | int(11)     | YES  |     | NULL    |       |
+-------+-------------+------+-----+---------+-------+
5 rows in set (0.00 sec)
```
* 修改字段：
```sql
# 改名改类型：
alter table `person` change name `phone` varchar(20);
Query OK, 0 rows affected (0.38 sec)
Records: 0  Duplicates: 0  Warnings: 0

describe `person`;
+-------+-------------+------+-----+---------+-------+
| Field | Type        | Null | Key | Default | Extra |
+-------+-------------+------+-----+---------+-------+
| sex   | varchar(2)  | YES  |     | NULL    |       |
| id    | int(11)     | YES  |     | NULL    |       |
| phone | varchar(20) | YES  |     | NULL    |       |
| age   | int(11)     | YES  |     | NULL    |       |
+-------+-------------+------+-----+---------+-------+
4 rows in set (0.00 sec)

# 只改类型：
alter table `person` modify `phone` varchar(16);
Query OK, 0 rows affected (0.48 sec)
Records: 0  Duplicates: 0  Warnings: 0

describe `person`;
+-------+-------------+------+-----+---------+-------+
| Field | Type        | Null | Key | Default | Extra |
+-------+-------------+------+-----+---------+-------+
| sex   | varchar(2)  | YES  |     | NULL    |       |
| id    | int(11)     | YES  |     | NULL    |       |
| phone | varchar(16) | YES  |     | NULL    |       |
| age   | int(11)     | YES  |     | NULL    |       |
+-------+-------------+------+-----+---------+-------+
4 rows in set (0.00 sec)

# 修改引擎：
alter table `person` engine=myisam;
Query OK, 0 rows affected (0.42 sec)
Records: 0  Duplicates: 0  Warnings: 0

# 修改表名：
alter table `person` rename to `student`;
Query OK, 0 rows affected (0.13 sec)

show tables;
+------------------+
| Tables_in_itcast |
+------------------+
| student          |
+------------------+
1 row in set (0.00 sec)
```

## 2.7 复制表
* 语法一：
```sql
create table `新表名` select *|`字段`, ... from `旧表`
```
**特点：** 不能复制表的主键，只能够复制数据

* 语法二：
```sql
create table `新表名` like `旧表名`
```
**特点：** 只能复制表的结构，不能复制表的数据

# 3.数据操作
* 生成测试表：
```sql
create table if not exists `stu`(
    `id` int not null auto_increment primary key comment 'primary key',
    `name` varchar(20) null comment 'name',
    `phone` varchar(20) null comment 'phone',
    `address` varchar(100) default 'empty' comment 'address',
    `score` int default 0 comment 'score'
)engine=innodb;
```

## 3.1 插入数据
* 语法：
```sql
insert into `表名`(`字段名`, `字段名`, ...) values(值, 值, ...), (值, 值, ...), ...
```
* 插入一条完整字段的数据：
```sql
insert into `stu`(`id`, `name`, `phone`, `address`, `score`) values(1, 'albert', '123456789', 'mit', 80);
```
* 插入部分段，自增字段忽略：
```sql
insert into `stu`(`name`, `phone`, `address`) values('lily', '131007896', 'bkl');
```
* 忽略字段指定插入数据，自增字段可插入null：
```sql
insert into `stu` values(null, 'king', '4567890', 'usa', 90);
```
* 通过 default 插入默认值：
```sql
insert into `stu` values(null, 'jack', '1237890', default, 70);
```
* 插入多条数据：
```sql
insert into `stu` values(null, 'tom', '123456', default, 80), (null, 'stack', '456789', default, 90);
Query OK, 2 row affected (0.06 sec)
Records: 2  Duplicates: 0  Warnings: 0
```

## 3.2 更新数据
* 语法：
```sql
update `表名` set `字段`=`值` [where 条件]
```
* 修改1号学生的姓名为sony：
```sql
update `stu` set `name`='sony' where `id`=1;
```
* 修改2号学生的电话和地址：
```sql
update `stu` set `phone`='123', `address`='beijing' where `id`=2;
```

## 3.3 删除数据
* 语法：
```sql
delete from `表名` [where 条件]
```
* 删除成绩小于等于70的学生：
```sql
delete from `stu` where `score` <= 70;
```
* 删除表中全部数据：
```sql
delete from `stu`;
```

## 3.4 清空表
* 语法：清空数据，自增主键从0开始重新计数
```sql
truncate table `表名`
```

## 3.5 简单查询
* 语法：
```sql
select *|`字段`, `字段`, ... from `表名`
```
* 查询所有字段的数据：
```sql
select * from `stu`;
```
* 查询指定字段的数据：
```sql
select `id`, `name`, `phone` from `stu`;
```

# 4.数据库扩展

## 4.1 SQL分类
> 1. DDL（data definition language）数据库定义语言 CREATE, ALTER, DROP；
> 2. DML（data manipulation language）数据操纵语言 SELECT, UPDATE, INSERT, DELETE；
> 3. DCL（data control language）数据库控制语言，是用来设置或更改数据库用户或角色权限的语句。

## 4.2 数据表的文件
> 1. 一个数据库对应一个文件夹；
> 2. 若 engine 是 myisam，则一张数据表对应三个文件，后缀为 .frm 的文件存储表结构，后缀为 .MYD 的文件存储表数据，后缀为 . MYI 的文件存储表索引； 
> 3. 若 engine 是 innodb，则一张数据表对应一个后缀为 .frm 的表结构文件，所有使用 innodb 引擎的表数据统一存储在 data\ibdata1 文件中，若数据量很大，则会自动创建 ibdata2, ibdata3, ... 目的就是便于管理；
> 4. 若 engine 是 memory，数据存储在内存中，重启MySQL服务数据丢失，但是读取速度非常快。

## 4.3 字符编码
* 设置编码语法：
```sql
set names 字符编码
```
**注：** 通过 set names 设置，就可以同时改变 character_set_client, character_set_results 的值。

* 查看服务端编码：
```sql
show variables like 'character_set_%';
```
**总结：** 客户端编码，character_set_client，character_set_results 三个编码一致即可操作中文。

# 5.数据类型

## 5.1 整型

|   类型    | 字节  |       范围       |
| :-------: | :--: | :--------------: |
|  tinyint  |  1   |     -128~127     |
| smallint  |  2   |   -32768~32767   |
| mediumint |  3   | -8388608~8388607 |
|    int    |  4   |  -2^31^~2^31^-1  |
|  bigint   |  8   |  -2^63^~2^63^-1  |

* 无符号整数（unsigned），无符号数没有负数，正数部分是有符号的两倍。
```sql
create table stu(
    `id` smallint unsigned auto_increment primary key comment '主键',
    `age` tinyint unsigned not null comment '年龄',
    `money` bigint unsigned comment '存款'
);
Query OK, 0 rows affected (0.06 sec)
```
* 整型支持显示宽度（最小的位数）如：int(5)，如果数值的位数小于5位，前面加上前导0（需要结合zerofill使用），大于5位的则不需要。
```sql
create table stu(
    `id` int(5),
    `age` int(5) zerofill  # 填充前导0
);
Query OK, 0 rows affected (0.02 sec)

insert into stu values (1,11);
Query OK, 0 rows affected (0.02 sec)

insert into stu values (1111111,2222222);
Query OK, 0 rows affected (0.03 sec)

select * from stu;
+---------+---------+
| id      | age     |
+---------+---------+
|       1 |   00011 |
| 1111111 | 2222222 | 
+---------+---------+
2 rows in set (0.00 sec)
```

## 5.2 浮点数

|      浮点型      | 占用字节 |        范围        |
| :--------------: | :------: | :----------------: |
| float（单精度）  |    4     |  -3.4E+38~3.4E+38  |
| double（双精度） |    8     | -1.8E+308~1.8E+308 |

* 语法：
```sql
float(M,D), double(M,D) # M: 总位数, D: 小数位数
```
* 例：
```sql
create table t1(
    num1 float(5,2),   # 总位数是5，小数位数是2，那么整数位数是3
    num2 double(4,1)
);
Query OK, 0 rows affected (0.08 sec)

insert into t1 values (1.23,1.23);   # 如果精度超出了允许的范围，会四舍五入
Query OK, 1 row affected (0.00 sec)

select * from t1;
+------+------+
| num1 | num2 |
+------+------+
| 1.23 |  1.2 |   # 四舍五入的结果
+------+------+
1 row in set (0.00 sec)
```

## 5.3 定点数
> * 定点数是变长的，大致每9个数字用4个字节来存储。定点数之所以能保存精确的小数，因为整数和小数是分开存储的。占用的资源比浮点数要多；
> * 定点数和浮点数都支持显示宽度和无符号数。

* 语法：
```sql
decimal(M,D)
```
```sql
create table t4(
    num decimal(20,19)
);
Query OK, 0 rows affected (0.00 sec)

insert into t4 values (1.1234567890123456789);
Query OK, 1 row affected (0.01 sec)

select * from t4;
+-----------------------+
| num                   |
+-----------------------+
| 1.1234567890123456789 |
+-----------------------+
1 row in set (0.00 sec)
```

## 5.4 字符型

|   数据类型    |   描述   |     长度      |
| :-----------: | :------: | :-----------: |
|  char(长度)   |   定长   |    最大255    |
| varchar(长度) |   变长   |   最大65535   |
|   tinytext    | 大段文本 |  2^8^-1=255   |
|     text      | 大段文本 | 2^16^-1=65535 |
|  mediumtext   | 大段文本 |    2^24^-1    |
|   longtext    | 大段文本 |    2^32^-1    |

* char(10) 和 varchar(10) 的区别？
    * 相同点：它们最多只能保存10个字符；
    * 不同点：char 不回收多余的字符，varchar 会回收多余的字符。 
* char 的最大长度是255；
* varchar 理论长度是65535字节；
* 大块文本（text）不计算在总长度中，一个大块文本只占用10个字节来保存文本的地址。

## 5.5 枚举
* 语法：
```sql
enum(值1, 值2, 值3, ...)
```
* 枚举类型在定义表结构时定义好多个值，插入数据时只能插入已经定义过的一个值：
```sql
create table t8(
    name varchar(20),
    sex enum('男', '女', '保密')
)charset=utf8;
Query OK, 0 rows affected (0.06 sec)

insert into t8 values ('tom', '男');
Query OK, 1 row affected (0.00 sec)

insert into t8 values ('berry', '女');
Query OK, 1 row affected (0.05 sec)

select * from t8;
+-------+------+
| name  | sex  |
+-------+------+
| tom   | 男   |
| berry | 女   |
+-------+------+
```
* MySQL 的枚举类型是通过整数来管理的，第一个值是1，第二个值是2，以此类推：
```sql
select sex+0 from t8;    # 进行运算可以显示数字类型
+-------+
| sex+0 |
+-------+
|     1 |
|     2 |
+-------+
```
* 枚举在数据库内部存储的是整数，可以直接插入数字：
```sql
insert into t8 values ('rose', 3);  # 可以直接插入数字
Query OK, 1 row affected (0.00 sec)

select * from t8;
+-------+------+
| name  | sex  |
+-------+------+
| tom   | 男   |
| berry | 女   |
| rose  | 保密 |
+-------+------+
3 rows in set (0.00 sec)
```

## 5.6 集合
* 集合类型在定义表结构时定义好多个值，插入数据时只能插入已经定义过的一个或多个值：
```sql
create table t9(
    hobby set('爬山', '读书', '游泳', '敲代码')
);
Query OK, 0 rows affected (0.08 sec)

insert into t9 values('爬山');
Query OK, 1 row affected (0.00 sec)

insert into t9 values('爬山, 游泳');
Query OK, 1 row affected (0.00 sec)

insert into t9 values('游泳, 爬山');  # 插入顺序不一样，但是显示的顺序是一样的
Query OK, 1 row affected (0.02 sec)
```
**注：** 每个集合的元素都分配一个固定的数字，从左往右按2的0，1，2，…次方分配。

## 5.7 日期

| 数据类型  |         描述          |   范围     |
| :-------: | :-------------------: | :-------:|
| datetime  | 日期时间，占用8个字节 | 1~9999年  |
|   date    |   日期 占用3个字节    | 1~9999年   |
|   time    |   时间 占用3个字节  | -838:59:59~838:59:59 |
| timestamp |  时间戳，占用4个字节  | 1970~2038年 |
|   year    |  年份   占用1个字节   | 1901-2155(255个) |

* time支持以天的方式插入
```sql
insert into t14 values ('10 10:10:10');  # 10天会转换为240个小时存储
Query OK, 1 row affected (0.02 sec)

select * from t14;
+-----------+
| field     |
+-----------+
| 250:10:10 |
+-----------+
```

## 5.8 Bool型
* MySQL 不支持真正意义上的 boolean 类型，true 和 false 在数据库中对应1和0：
```sql
create table t15(
    field boolean
);
Query OK, 0 rows affected (0.00 sec)

insert into t15 values (true),(false);
Query OK, 2 rows affected (0.00 sec)
Records: 2  Duplicates: 0  Warnings: 0

select * from t15;
+-------+
| field |
+-------+
|     1 |
|     0 |
+-------+
2 rows in set (0.00 sec)
```

# 6.列属性

## 6.1 主键
> * 主键是唯一标识表中记录的一个或一组列；
> * 特点是不能重复，不能为空；
> * 一个表只能有一个主键，主键可以有多个字段组成；
> * 作用是保证数据完整性，加快查询速度。

* 更改表的时候添加主键：
```sql
create table `t20`(
    `id` int,
    `name` varchar(10)
);
Query OK, 0 rows affected (0.00 sec)

alter table `t20` add primary key (`id`);
Query OK, 0 rows affected (0.08 sec)
Records: 0  Duplicates: 0  Warnings: 0
```
* 删除主键：
```sql
alter table `t20` drop primary key;
```

## 6.2 唯一键
> * 特点：不能重复，可以为空，一个表可以有多个；
> * 作用：保证数据不重复，保证数据完整性，加快数据访问。

* 创建表的时候添加唯一键：
```sql
create table `t22`(
    `id` int primary key,   
    `name` varchar(20) unique,   # unique关键字添加唯一键
    `addr` varchar(100) unique
);
Query OK, 0 rows affected (0.00 sec)
```
* 修改表的时候添加唯一键：
```sql
create table `t23`(
    `id` int primary key,
    `name` varchar(20),
    `addr` varchar(20),
    `score` int,
    `age` int
);
Query OK, 0 rows affected (0.02 sec)

alter table `t23` add unique(`name`), add unique(`addr`);
Query OK, 0 rows affected (0.02 sec)
Records: 0  Duplicates: 0  Warnings: 0

alter table `t23` add unique(`score`, `age`);   # 组合唯一键，表示两个字段的值组合起来是唯一的
Query OK, 0 rows affected (0.01 sec)
Records: 0  Duplicates: 0  Warnings: 0
```
* 删除唯一键：
```sql
alter table `t23` drop index `name`;     # 按照唯一键的名字删除，唯一键的名字默认就是字段名称
```

## 6.3 外键

### 6.3.1 添加外键
* 语法：
```sql
foreign key (`外键字段`) references `主表名`(`主键字段`)
```
* 创建表的时候添加外键：
```sql
create table if not exists `stuinfo`(
    `stuno` char(4) primary key,
    `name` varchar(10) not null
);
Query OK, 0 rows affected (0.05 sec)

create table if not exists `stumarks`(
    `stuid` char(4) primary key,
    `score` tinyint unsigned,
    foreign key (`stuid`) references `stuinfo`(`stuno`)
);
Query OK, 0 rows affected (0.09 sec)
```
> * 主表中不存在的，从表中不允许插入；
> * 从表中存在的，主表中不允许删除；
> * 不能更改主表的数据后使得从表孤立。

* 语法：
```sql
alter table `从表名` add foreign key (`从表公共字段`) references `主从名`(`主表的公共字段`)
```
* 修改表的时候添加外键约束：
```sql
create table if not exists `teainfo`(
    `teano` int primary key,
    `name` varchar(10) not null
);
Query OK, 0 rows affected (0.05 sec)

create table if not exists `teamarks`(
    `teaid` int primary key,
    `score` tinyint unsigned
);
Query OK, 0 rows affected (0.05 sec)

alter table `teamarks` add foreign key (`teaid`) references `teainfo`(`teano`);
```
**注：** 要创建外键必须是innodb引擎，myisam不支持外键约束。

### 6.3.2 删除外键：
* 语法：
```sql
alter table `表名` drop foreign key `外建名`
```
* 例：
```sql
show create table `stumarks`\G;
*************************** 1. row ***************************
       Table: stumarks
Create Table: CREATE TABLE `stumarks` (
  `stuid` char(4) NOT NULL,
  `score` tinyint(3) unsigned DEFAULT NULL,
  PRIMARY KEY (`stuid`),
  CONSTRAINT `stumarks_ibfk_1` FOREIGN KEY (`stuid`) REFERENCES `stuinfo` (`stuno`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
1 row in set (0.00 sec)

alter table `stumarks` drop foreign key `stumarks_ibfk_1`;
Query OK, 0 rows affected (0.16 sec)
Records: 0  Duplicates: 0  Warnings: 0
```

### 6.3.3 外键操作：
* 更新时级联，删除时置空（更新主表公共字段时从表也对应的更新，主表删除数据时从表将对应的公共字段置空）：
```sql
create table `stuinfo`(
    `stunno` char(4) primary key,
    `name` var char(10) not null
);
Query OK, 0 rows affected (0.02 sec)

create table `stumarks`(
    `stuid` int auto_increment primary key,
    `stuno` char(4),
    `score` tinyint unsigned,
    foreign key (`stuno`) reference `stuinfo`(`stuno`) on delete set null on update cascade
);
Query OK, 0 rows affected (0.00 sec)

insert into stuinfo values ('s101','tom');
Query OK, 1 row affected (0.00 sec)

insert into stumarks values (null,'s101',88);
Query OK, 1 row affected (0.00 sec)

select * from stuinfo;
+-------+------+
| stuno | name |
+-------+------+
| s101  | tom  |
+-------+------+
1 row in set (0.00 sec)

update stuinfo set stuno='s102' where stuno='s101';   # 更新时级联
Query OK, 1 row affected (0.00 sec)
Rows matched: 1  Changed: 1  Warnings: 0

select * from stumarks;
+-------+-------+-------+
| stuid | stuno | score |
+-------+-------+-------+
|     1 | s102  |    88 |
+-------+-------+-------+
1 row in set (0.00 sec)

delete from stuinfo where stuno='s102';    # 删除时置空
Query OK, 1 row affected (0.02 sec)

select * from stumarks;
+-------+-------+-------+
| stuid | stuno | score |
+-------+-------+-------+
|     1 | NULL  |    88 |
+-------+-------+-------+
1 row in set (0.00 sec)
```

# 7.数据查询

## 7.1 单表查询
* 语法：
```sql
select [选项] `字段名` [from `表名`] [where 条件] [order by `排序字段`] [group by `分组字段`] [having 条件] [limit 限制]
```

### 7.1.1 字段表达式
```sql
select 10 * 10 as `result`;
+--------+
| result |
+--------+
|    100 |
+--------+
1 row in set (0.00 sec)
```

### 7.1.2 from 子句
* 指定从哪张或哪些表中查询，若是多张表，最终结果是所有表结果的笛卡尔积：
```sql
create table t01(
    `id` int,
    `name` varchar(10)
);
Query OK, 0 rows affected (1.82 sec)

create table t02(
    `score` int,
    `address` varchar(10)
);
Query OK, 0 rows affected (0.21 sec)

insert into t01 values(1, 'albert');
Query OK, 1 row affected (0.01 sec)

insert into t01 values(2, 'lily');
Query OK, 1 row affected (0.06 sec)

insert into t02 values(70, 'mit');
Query OK, 1 row affected (0.37 sec)

insert into t02 values(80, 'bkl');
Query OK, 1 row affected (0.19 sec)

select * from t01,t02;
+------+--------+-------+---------+
| id   | name   | score | address |
+------+--------+-------+---------+
|    1 | albert |    70 | mit     |
|    2 | lily   |    70 | mit     |
|    1 | albert |    80 | bkl     |
|    2 | lily   |    80 | bkl     |
+------+--------+-------+---------+
4 rows in set (0.00 sec)
```

### 7.1.3 dual 伪表
* 当将 select 当做表达式使用时，添加 from dual 保证SQL语句的完整性：
```sql
select 10*10 from dual;
+-------+
| 10*10 |
+-------+
|   100 |
+-------+
1 row in set (0.00 sec)
```
### 7.1.4 where 子句
* where子句对查找的数据进行条件过滤：

| 运算符 | 说明 |
|:-----:|:----:|
| `>` | 大于 |
| `<` | 小于 |
| `>=` | 大于等于 |
| `<=` | 小于等于 |
| `!=` | 不等于 |
| `and` | 与 |
| `or` | 或 |
| `not` | 非 |

### 7.1.5 in \| not in
* 判断字段值是否出现在一段指定的集合中。
* 查找住在北京，上海，天津，并且姓名不是Tom的学生：
```SQL
select * from `stu` where `stuaddress` in('北京', '上海', '天津') and `stuName` not in('Tom');
+--------+--------------+--------+--------+---------+------------+------+------+
| stuNo  | stuName      | stuSex | stuAge | stuSeat | stuAddress | ch   | math |
+--------+--------------+--------+--------+---------+------------+------+------+
| s25301 | 张秋丽       | 男     |     18 |       1 | 北京       |   80 | NULL |
| s25302 | 李文才       | 男     |     31 |       3 | 上海       |   77 |   76 |
| s25303 | 李斯文       | 女     |     22 |       2 | 北京       |   55 |   82 |
| s25304 | 欧阳俊雄     | 男     |     28 |       4 | 天津       | NULL |   74 |
| s25318 | 争青小子     | 男     |     26 |       6 | 天津       |   86 |   92 |
+--------+--------------+--------+--------+---------+------------+------+------+
5 rows in set (0.00 sec)
```

### 7.1.6 between ... and \| not between ... and
* 判断字段值是否出现在一段范围中。
* 查找年龄在18~30之间，并且数学成绩不在60~80之间的学生：
```SQL
select * from `stu` where (`stuAge` between 18 and 30) and (`math` not between 60 and 80);
+--------+--------------+--------+--------+---------+------------+------+------+
| stuNo  | stuName      | stuSex | stuAge | stuSeat | stuAddress | ch   | math |
+--------+--------------+--------+--------+---------+------------+------+------+
| s25303 | 李斯文       | 女     |     22 |       2 | 北京       |   55 |   82 |
| s25305 | 诸葛丽丽     | 女     |     23 |       7 | 河南       |   72 |   56 |
| s25318 | 争青小子     | 男     |     26 |       6 | 天津       |   86 |   92 |
+--------+--------------+--------+--------+---------+------------+------+------+
3 rows in set (0.01 sec)
```

### 7.1.7 is null \| is not null
* 判断字段值是否为空。
* 查找语文缺考的，数学没有缺考的同学：
```SQL
mysql> select * from `stu` where `ch` is null and `math` is not null;
+--------+--------------+--------+--------+---------+------------+------+------+
| stuNo  | stuName      | stuSex | stuAge | stuSeat | stuAddress | ch   | math |
+--------+--------------+--------+--------+---------+------------+------+------+
| s25304 | 欧阳俊雄     | 男     |     28 |       4 | 天津       | NULL |   74 |
+--------+--------------+--------+--------+---------+------------+------+------+
1 row in set (0.00 sec)
```
### 7.1.8 聚合函数
* 常用聚合函数：

| 函数 | 说明 |
|:-----:|:----:|
| sum() | 求和 |
| avg() | 求平均值 |
| max() | 求最大值 |
| min() | 求最小值 |
| count() | 求记录数 |

* 例：统计所有同学语文成绩的总值，平均值，最大值，最小值，记录数：
```SQL
select sum(`ch`) sum, avg(`ch`) avg, max(`ch`) max, min(`ch`) min, count(`ch`) count from `stu`;
+------+---------+------+------+-------+
| sum  | avg     | max  | min  | count |
+------+---------+------+------+-------+
|  597 | 74.6250 |   88 |   55 |     8 |
+------+---------+------+------+-------+
1 row in set (0.00 sec)
```

### 7.1.9 like 模糊查询
* 通配符：
    * `_` 表示任意一个字符；
    * `%` 表任意个数的任意字符。

* 查询所有“张”姓的学生：
```SQL
select * from `stu` where `stuname` like '张%';
+--------+-----------+--------+--------+---------+------------+------+------+
| stuNo  | stuName   | stuSex | stuAge | stuSeat | stuAddress | ch   | math |
+--------+-----------+--------+--------+---------+------------+------+------+
| s25301 | 张秋丽    | 男     |     18 |       1 | 北京       |   80 | NULL |
+--------+-----------+--------+--------+---------+------------+------+------+
1 row in set (0.00 sec)
```
### 7.1.10 order by 排序
* asc 升序；
* desc 降序。

* 语文成绩按降序排序。数学成绩按升序排序：
```SQL
select * from `stu` order by `ch` desc;
select * from `stu` order by `nath` asc;
```

### 7.1.11 group by 分组查询
* 语法：
```SQL
group by `字段名`, `字段名`, ...   # 根据相同的字段值将数据分成一组（多个字段会按照字段组合进行分组）
```
* 统计男女生的平均年龄：
```SQL
select avg(`stuAge`) as `平均年龄`, `stusex` from `stu` group by `stusex` desc;
+--------------+--------+
| 平均年龄     | stusex |
+--------------+--------+
|      25.4000 | 男     |
|      22.7500 | 女     |
+--------------+--------+
2 rows in set (0.00 sec)
```
* 统计相同地区男女生的平均年龄：
```SQL
select `stuAddress`, `stuSex`, avg(`stuAge`) from `stu` group by `stuAddress`, `stuSex`;
+------------+--------+---------------+
| stuAddress | stuSex | avg(`stuAge`) |
+------------+--------+---------------+
| 北京       | 男     |       21.0000 |
| 上海       | 男     |       31.0000 |
| 北京       | 女     |       22.0000 |
| 天津       | 男     |       27.0000 |
| 河南       | 女     |       23.0000 |
| 河北       | 女     |       23.0000 |
+------------+--------+---------------+
6 rows in set (0.00 sec)
```

### 7.1.12 having 条件
* where 和 having 的区别：
    * where 是直接对原始数据库的文件按一条条字段进行筛选；
    * having 是对 select...from... 映射出的虚拟表进行筛选。
* where 和 having的区别示例1：
    * where 先从原始库文件的所有数据一条一条过滤，筛选出符合条件的结果，select...from... 再对结果进行抽取映射；
    * select...from... 先从原始库文件中按表按字段抽取数据映射成虚拟表，having 再从虚拟表中按照条件筛选数据。
```SQL
select `stuname` from stu where `stusex`='男';   
+--------------+
| stuname      |
+--------------+
| 张秋丽       |
| 李文才       |
| 欧阳俊雄     |
| 争青小子     |
| Tom          |
+--------------+
5 rows in set (0.00 sec)

select `stuname` from `stu` having  `stusex`='男';
ERROR 1054 (42S22): Unknown column 'stusex' in 'having clause'
```

* where 和 having的区别示例2：
    * where 会先在数据库中对全部数据进行筛选，之后才会通过 select...from... 语句进行按表按字段映射虚拟表，total 字段是基于虚拟表定义的，where 自然无法识别此字段；
    * having 条件过滤操作是在 select...from... 映射虚拟表之后才会执行的，对虚拟表的数据进行筛选，total 是基于虚拟表定义的，所以 having 可以识别并操作。
```SQL
select `stuSex`, count(*) as total from `stu` group by `stuSex` where total >= 5;
ERROR 1064 (42000): You have an error in your SQL syntax; check the manual that corresponds to your MySQL server version for the right syntax to use near 'where total >= 5' at line 1

select `stuSex`, count(*) as total from `stu` group by `stuSex` having total >= 5;
+--------+-------+
| stuSex | total |
+--------+-------+
| 男     |     5 |
+--------+-------+
1 row in set (0.00 sec)
```

### 7.1.13 limit 抽取限制
* 语法：
```SQL
limit index, num      # index 表示限制开始的下标，num表示限制的条目数
```
* 查询班级总分前3的同学：
```SQL
select `stuName`, (`ch` + `math`) as `total` from `stu` order by `total` desc limit 0, 3;
+--------------+-------+
| stuName      | total |
+--------------+-------+
| 争青小子     |   178 |
| Tabm         |   165 |
| 李文才       |   153 |
+--------------+-------+
3 rows in set (0.00 sec)
```

### 7.1.14 查询语句中的选项
> * all 显示所有数据（默认）；
> * distinct 去除结果集中重复的数据。

* 查找学生居住在那些城市：
```SQL
select distinct `stuAddress` from `stu`;
+------------+
| stuAddress |
+------------+
| 北京       |
| 上海       |
| 天津       |
| 河南       |
| 河北       |
+------------+
5 rows in set (0.00 sec)
```

## 7.2 多表查询

### 7.2.1 union 联合
* 语法：
```SQL
select 语句 union [选项] select 语句 union [选项] select ...
选项：all 显示全部，distinct 去重（默认）
```
* 将多条查询语句的结果纵向组合：
```SQL
create table t03(
    `id` int,
    `name` varchar(10)
);
Query OK, 0 rows affected (1.58 sec)

insert into t03 values(1, 'albert'), (2, 'lily');
Query OK, 2 rows affected (0.16 sec)
Records: 2  Duplicates: 0  Warnings: 0

select `stuNo`, `stuName` from `stu` union select `id`, `name` from `t03`;
+--------+--------------+
| stuNo  | stuName      |
+--------+--------------+
| s25301 | 张秋丽       |
| s25302 | 李文才       |
| s25303 | 李斯文       |
| s25304 | 欧阳俊雄     |
| s25305 | 诸葛丽丽     |
| s25318 | 争青小子     |
| s25319 | 梅超风       |
| s25320 | Tom          |
| s25321 | Tabm         |
| 1      | albert       |
| 2      | lily         |
+--------+--------------+
11 rows in set (0.00 sec)
```
> **注：**
> 1. union 两边的 select 语句的字段个数必须一致；
> 2. union 两边的 select 语句的字段名可以不一样，最终结果按照第一个 select 语句的字段名；
> 3. union 两边的 select 语句的数据类型可以不一致。

### 7.2.2 inner join 内连接
> * 内连接获取的是相对于连接字段的两个表的公共数据。

* 语法：
```SQL
select `字段名` from `表1` inner join `表2` on `表1`.`公共字段` = `表2`.`公共字段`
select `字段名` from `表1`, `表2` where `表1`.`公共字段` = `表2`.`公共字段`
```
* 查询学生的全部信息：
```SQL
select stuName, stuSex, stuAge, stuSeat, stuAddress, writtenExam, labExam from stuinfo i join stumarks m on i.stuNo=m.stuNo;
+--------------+--------+--------+---------+------------+-------------+---------+
| stuName      | stuSex | stuAge | stuSeat | stuAddress | writtenExam | labExam |
+--------------+--------+--------+---------+------------+-------------+---------+
| 李斯文       | 女     |     22 |       2 | 北京       |          80 |      58 |
| 李文才       | 男     |     31 |       3 | 上海       |          50 |      90 |
| 欧阳俊雄     | 男     |     28 |       4 | 天津       |          65 |      50 |
| 张秋丽       | 男     |     18 |       1 | 北京       |          77 |      82 |
| 争青小子     | 男     |     26 |       6 | 天津       |          56 |      48 |
+--------------+--------+--------+---------+------------+-------------+---------+
5 rows in set (0.00 sec)
```
* 内连接3张表查询：
```SQL
select * from stuinfo i inner join stumarks m on i.stuNo=m.stuNo inner join stu s on m.stuNo=s.stuNo;
+--------+--------------+--------+--------+---------+------------+---------+--------+-------------+---------+--------+--------------+--------+--------+---------+------------+------+------+
| stuNo  | stuName      | stuSex | stuAge | stuSeat | stuAddress | examNo  | stuNo  | writtenExam | labExam | stuNo  | stuName      | stuSex | stuAge | stuSeat | stuAddress | ch   | math |
+--------+--------------+--------+--------+---------+------------+---------+--------+-------------+---------+--------+--------------+--------+--------+---------+------------+------+------+
| s25303 | 李斯文       | 女     |     22 |       2 | 北京       | s271811 | s25303 |          80 |      58 | s25303 | 李斯文       | 女     |     22 |       2 | 北京       |   55 |   82 |
| s25302 | 李文才       | 男     |     31 |       3 | 上海       | s271813 | s25302 |          50 |      90 | s25302 | 李文才       | 男     |     31 |       3 | 上海       |   77 |   76 |
| s25304 | 欧阳俊雄     | 男     |     28 |       4 | 天津       | s271815 | s25304 |          65 |      50 | s25304 | 欧阳俊雄     | 男     |     28 |       4 | 天津       | NULL |   74 |
| s25301 | 张秋丽       | 男     |     18 |       1 | 北京       | s271816 | s25301 |          77 |      82 | s25301 | 张秋丽       | 男     |     18 |       1 | 北京       |   80 | NULL |
| s25318 | 争青小子     | 男     |     26 |       6 | 天津       | s271819 | s25318 |          56 |      48 | s25318 | 争青小子     | 男     |     26 |       6 | 天津       |   86 |   92 |
+--------+--------------+--------+--------+---------+------------+---------+--------+-------------+---------+--------+--------------+--------+--------+---------+------------+------+------+
5 rows in set (0.00 sec)
```

### 7.2.3 left join 左外连接
> * 左外连接获取的是以左表为主的数据（会保留所有左表数据），右表若没有对应字段的数据，则显示为NULL。

* 语法：
```SQL
select `字段名` from `表1` left join `表2` on `表1`.`公共字段` = `表2`.`公共字段`
```
* 以左边的表为标准，如果右边的表没有对应记录，用 NULL 填充：
```SQL
select stuName, writtenexam, labexam from stuinfo i left join stumarks m on i.stuNo=m.stuNo;
+--------------+-------------+---------+
| stuName      | writtenexam | labexam |
+--------------+-------------+---------+
| 李斯文       |          80 |      58 |
| 李文才       |          50 |      90 |
| 欧阳俊雄     |          65 |      50 |
| 张秋丽       |          77 |      82 |
| 争青小子     |          56 |      48 |
| 诸葛丽丽     |        NULL |    NULL |
| 梅超风       |        NULL |    NULL |
+--------------+-------------+---------+
7 rows in set (0.00 sec)
```

### 7.2.4 right join 右外连接
> * 右外连接获取的是以右表为主的数据（会保留所有右表数据），左表若没有对应字段的数据，则显示为NULL。

* 语法：
```SQL
select `字段名` from `表1` right join `表2` on `表1`.`公共字段` = `表2`.`公共字段`
```
* 例：
```SQL
select stuName, writtenexam, labexam from stuinfo i right join stumarks m on i.stuNo=m.stuNo;
+--------------+-------------+---------+
| stuName      | writtenexam | labexam |
+--------------+-------------+---------+
| 李斯文       |          80 |      58 |
| 李文才       |          50 |      90 |
| 欧阳俊雄     |          65 |      50 |
| 张秋丽       |          77 |      82 |
| 争青小子     |          56 |      48 |
| NULL         |          66 |      77 |
+--------------+-------------+---------+
6 rows in set (0.00 sec)
```

### 7.2.5 cross join 交叉连接
* 语法：
```SQL
select `字段名` from `表1` cross join `表2`
select `字段名` from `表1` cross join `表2` where `表1`.`公共字段` = `表2`.`公共字段`
```
* 生成测试数据：
```SQL
create table t04(
    `id` int,
    `name` varchar(10)
);
Query OK, 0 rows affected (0.46 sec)

create table t05(
    `id` int,
    `score` int
);
Query OK, 0 rows affected (0.18 sec)

insert into t04 values(1, 'albert'), (2, 'lily');
Query OK, 2 rows affected (0.02 sec)
Records: 2  Duplicates: 0  Warnings: 0

insert into t05 values(1, 80), (2, 99);
Query OK, 2 rows affected (0.07 sec)
Records: 2  Duplicates: 0  Warnings: 0
```
* 没有连接表达式（取两表的笛卡尔积）：
```SQL
select * from t04 cross join t05;
+------+--------+------+-------+
| id   | name   | id   | score |
+------+--------+------+-------+
|    1 | albert |    1 |    80 |
|    2 | lily   |    1 |    80 |
|    1 | albert |    2 |    99 |
|    2 | lily   |    2 |    99 |
+------+--------+------+-------+
4 rows in set (0.00 sec)
```
* 有连接表达式（相当于内连接）：
```SQL
select * from t04 cross join t05 where t04.id = t05.id;
+------+--------+------+-------+
| id   | name   | id   | score |
+------+--------+------+-------+
|    1 | albert |    1 |    80 |
|    2 | lily   |    2 |    99 |
+------+--------+------+-------+
2 rows in set (0.00 sec)
```

### 7.2.6 natural 自然连接
> * 自然内连接：natural join；
> * 自然左外连接：natural left join；
> * 自然右外连接：natural right join。

* 自动检测公共字段连接多表：
```SQL
select * from t04 natural join t05;
+------+--------+-------+
| id   | name   | score |
+------+--------+-------+
|    1 | albert |    80 |
|    2 | lily   |    99 |
+------+--------+-------+
2 rows in set (0.00 sec)
```
> **特点：**
    > * 表连接通过同名字段连接；
    > * 没有同名字段返回笛卡尔积；
    > * 自动对结果进行整理，保留一个连接字段，连接字段放最前面。 

### 7.2.7 using()
> * 用于指定多个连接字段；
> * 对连接字段进行整理，整理方式同自然连接。

```SQL
select * from t04 join t05 using(id);
+------+--------+-------+
| id   | name   | score |
+------+--------+-------+
|    1 | albert |    80 |
|    2 | lily   |    99 |
+------+--------+-------+
2 rows in set (0.00 sec)
```

### 7.2.8 子查询
* 查找笔试成绩80分的学生信息：
```SQL
select * from `stuinfo` where `stuNo` = (select stuNo from `stumarks` where `writtenexam` = 80);
+--------+-----------+--------+--------+---------+------------+
| stuNo  | stuName   | stuSex | stuAge | stuSeat | stuAddress |
+--------+-----------+--------+--------+---------+------------+
| s25303 | 李斯文    | 女     |     22 |       2 | 北京       |
+--------+-----------+--------+--------+---------+------------+
1 row in set (0.07 sec)
```
* 查找笔试成绩最高的学生成绩：
```SQL
select * from `stuinfo` where `stuNo` = (select `stuNo` from `stumarks` order by `writtenexam` desc limit 1);
+--------+-----------+--------+--------+---------+------------+
| stuNo  | stuName   | stuSex | stuAge | stuSeat | stuAddress |
+--------+-----------+--------+--------+---------+------------+
| s25303 | 李斯文    | 女     |     22 |       2 | 北京       |
+--------+-----------+--------+--------+---------+------------+
1 row in set (0.00 sec)
```
* 查找笔试成绩及格的学生：
```SQL
select * from `stuinfo` where `stuNo` in (select `stuNo` from `stumarks` where `writtenexam` >= 60);
+--------+--------------+--------+--------+---------+------------+
| stuNo  | stuName      | stuSex | stuAge | stuSeat | stuAddress |
+--------+--------------+--------+--------+---------+------------+
| s25303 | 李斯文       | 女     |     22 |       2 | 北京       |
| s25304 | 欧阳俊雄     | 男     |     28 |       4 | 天津       |
| s25301 | 张秋丽       | 男     |     18 |       1 | 北京       |
+--------+--------------+--------+--------+---------+------------+
3 rows in set (0.04 sec)
```
* 若有学生的成绩大于等于80分，则显示所有学生的信息：
```SQL
select * from `stuinfo` where exists (select * from `stumarks` where `writtenexam` >= 80);
+--------+--------------+--------+--------+---------+------------+
| stuNo  | stuName      | stuSex | stuAge | stuSeat | stuAddress |
+--------+--------------+--------+--------+---------+------------+
| s25301 | 张秋丽       | 男     |     18 |       1 | 北京       |
| s25302 | 李文才       | 男     |     31 |       3 | 上海       |
| s25303 | 李斯文       | 女     |     22 |       2 | 北京       |
| s25304 | 欧阳俊雄     | 男     |     28 |       4 | 天津       |
| s25305 | 诸葛丽丽     | 女     |     23 |       7 | 河南       |
| s25318 | 争青小子     | 男     |     26 |       6 | 天津       |
| s25319 | 梅超风       | 女     |     23 |       5 | 河北       |
+--------+--------------+--------+--------+---------+------------+
7 rows in set (0.00 sec)
```
> **行子查询：** 子查询返回的结果是一行数据

* 从 stu 表中查询成绩最高的男生和女生的信息：
```SQL
# 单表子查询：
select *, (`ch` + `math`) as `total` from `stu` where (`stusex`, `ch` + `math`) in (select `stusex`, max(`ch` + `math`) from `stu` group by `stusex`);
+--------+--------------+--------+--------+---------+------------+------+------+-------+
| stuNo  | stuName      | stuSex | stuAge | stuSeat | stuAddress | ch   | math | total |
+--------+--------------+--------+--------+---------+------------+------+------+-------+
| s25318 | 争青小子     | 男     |     26 |       6 | 天津       |   86 |   92 |   178 |
| s25321 | Tabm         | 女     |     23 |       9 | 河北       |   88 |   77 |   165 |
+--------+--------------+--------+--------+---------+------------+------+------+-------+
2 rows in set (0.00 sec)
```
* 从 stuinfo 和 stumarks 表中查询成绩最高的男生和女生的信息：
```SQL
# 多表子查询：
select * from stuinfo i join stumarks m on i.stuNo = m.stuNo where (stuSex, labExam) in (select stuSex, max(labExam) as t from stuinfo i join stumarks m on i.stuNo = m.stuNo group by stuSex);
+--------+-----------+--------+--------+---------+------------+---------+--------+-------------+---------+
| stuNo  | stuName   | stuSex | stuAge | stuSeat | stuAddress | examNo  | stuNo  | writtenExam | labExam |
+--------+-----------+--------+--------+---------+------------+---------+--------+-------------+---------+
| s25303 | 李斯文    | 女     |     22 |       2 | 北京       | s271811 | s25303 |          80 |      58 |
| s25302 | 李文才    | 男     |     31 |       3 | 上海       | s271813 | s25302 |          50 |      90 |
+--------+-----------+--------+--------+---------+------------+---------+--------+-------------+---------+
2 rows in set (1.72 sec)
```

# 8.视图

## 8.1 创建视图
> * 视图是一张虚拟表，表示一张表的部分或多张表的综合；
> * 视图仅存在表结构，不存在表数据，视图的结构和数据建立在原始表的基础上；
> * 创建视图后，会在数据库文件夹中创建一个以视图名命名的 .frm 文件。

* 语法：
```SQL
create [or replace] view `视图名称`
as 
select 语句
```
* 例：
```SQL
create view `vw_stu`
as
select `stuName`, `stuSex`, `writtenexam`, `labexam` from `stuinfo` i, `stumarks` m where i.`stuNo` = m.`stuNo`;
Query OK, 0 rows affected (0.13 sec)
```

## 8.2 查询视图
```SQL
select * from `vw_stu`;
+--------------+--------+-------------+---------+
| stuName      | stuSex | writtenexam | labexam |
+--------------+--------+-------------+---------+
| 李斯文       | 女     |          80 |      58 |
| 李文才       | 男     |          50 |      90 |
| 欧阳俊雄     | 男     |          65 |      50 |
| 张秋丽       | 男     |          77 |      82 |
| 争青小子     | 男     |          56 |      48 |
+--------------+--------+-------------+---------+
5 rows in set (0.00 sec)
```

## 8.3 查看视图结构
```SQL
desc `vw_stu`;
+-------------+-------------+------+-----+---------+-------+
| Field       | Type        | Null | Key | Default | Extra |
+-------------+-------------+------+-----+---------+-------+
| stuName     | varchar(10) | NO   |     | NULL    |       |
| stuSex      | char(2)     | NO   |     | NULL    |       |
| writtenexam | int(11)     | YES  |     | NULL    |       |
| labexam     | int(11)     | YES  |     | NULL    |       |
+-------------+-------------+------+-----+---------+-------+
4 rows in set (0.14 sec)
```

## 8.4 查看创建视图的SQL
```SQL
show create view `vw_stu`\G;
*************************** 1. row ***************************
                View: vw_stu
         Create View: CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` SQL SECURITY DEFINER VIEW `vw_stu` AS select `i`.`stuName` AS `stuName`,`i`.`stuSex` AS `stuSex`,`m`.`writtenExam` AS `writtenexam`,`m`.`labExam` AS `labexam` from (`stuinfo` `i` join `stumarks` `m`) where (`i`.`stuNo` = `m`.`stuNo`)
character_set_client: utf8
collation_connection: utf8_general_ci
1 row in set (0.03 sec)
```

## 8.5 显示所有视图
```SQL
show tables;
+------------------+
| Tables_in_itcast |
+------------------+
| stu              |
| stuinfo          |
| stumarks         |
| t01              |
| t02              |
| t03              |
| t04              |
| t05              |
| vw_stu           |
+------------------+
9 rows in set (0.03 sec)

select table_name from information_schema.views;
```

## 8.6 更改视图
* 语法：
```SQL
alter view `视图名`
as
select 语句
```
* 例：
```SQL
alter view `vw_stu`
as
select * from `stuinfo`;
Query OK, 0 rows affected (0.33 sec)
```

## 8.7 删除视图
* 语法：
```SQL
drop view [if exists] `视图1`, `视图2`, ...
```
* 例：
```SQL
mysql> drop view `vw_stu`;
Query OK, 0 rows affected (0.29 sec)
```

## 8.8 视图的作用
> * 筛选数据；
> * 隐藏表结构；
> * 降低SQL语句的复杂度。

## 8.9 视图的算法
> * merge：合并算法，将视图的语句和外层的语句合并后在执行；
> * temptable：临时表算法，将视图生成一个临时表，再执行外层语句；
> * undefined：未定义，MySQL 到底用 merge 还是用 temptable 由 MySQL 决定，这是一个默认的算法，一般视图都会选择 merge 算法，因为 merge 效率高。

* 利用视图查询：
```SQL
create algorithm=temptable view `v-stu`
as
select * from `stu` order by `stuSex`;
Query OK, 0 rows affected (0.48 sec)
```

# 9.事务

## 9.1 开启，提交，回滚事务
* 语法：
```SQL
start transaction | begin   # 开启事务
commit      # 提交事务
rollback    # 回滚事务
```
* 例：
```SQL
# 更改定界符
delimiter //     

# 开启事务
start transaction;
update bank set money = money - 300 where cardid = 1001;
update bank set money = money + 300 where cardid = 1002;
//
Query OK, 0 rows affected (0.05 sec)

# 若所有语句全部执行通过，提交事务：
commit //

# 若存在语句执行出错，回滚事务：
rollback //

# 更改定界符
delimiter ;
```

## 9.2 设置事务的回滚点
* 语法：
```SQL
savepoint 回滚点    # 设置自定义回滚点
rollback to 回滚点  # 回滚到自定义的回滚点
```
* 例：
```SQL
delimiter //

insert into bank values(1003, 1000);
savepoint rb;    # 若触发回滚，则从回滚到此处
insert into bank values(1004, 1000);
//
Query OK, 1 row affected (1.93 sec)
Query OK, 0 rows affected (2.11 sec)
Query OK, 1 row affected (2.11 sec)

rollback to rb //
Query OK, 0 rows affected (1.67 sec)

delimiter ;
```

## 9.3 事务的特性
> * 原子性(Atomicity)：事务是一个整体，不可再分，一起执行或一起不执行；
> * 一致性(Consistency)：事务完成时，数据必须处于一致的状态；
> * 隔离性(Lsolation)：每个事务都是相互隔离的；
> * 永久性(Durability)：事务完成后，对数据的修改是永久性的。

# 10.索引

## 10.1 索引的类型
> * 普通索引；
> * 唯一索引（唯一键）；
> * 主键索引：只要是主键就字段创建主键索引，不需要手动创建；
> * 全文索引。

## 10.2 创建普通索引
```SQL
方式1：
create index [`索引名`] on `表名` (`字段名`)

方式2：
alter table `表名` add index [`索引名`] (`字段名`)
```
* 例：
```SQL
# 方式1：直接创建索引
create index i_name on stuinfo(stuName);
Query OK, 0 rows affected (1.42 sec)
Records: 0  Duplicates: 0  Warnings: 0

# 方式2：修改表的方式添加索引
alter table stuinfo add index i_address (stuAddress);
Query OK, 0 rows affected (0.63 sec)
Records: 0  Duplicates: 0  Warnings: 0

# 方式3：创建表时添加索引
create table emp(
    `id` int,
    `name` varchar(10),
    ndex i_name (name)
);
Query OK, 0 rows affected (2.14 sec)
```

## 10.3 创建唯一索引
* 语法：
```SQL
方式1：
create unique index `索引名` on `表名` (`字段名`)

方式2：
alter table `表名` add unique [index] [`索引名`] (`字段名`)
```
* 例：
```SQL
# 方式1：直接创建索引
create unique index UQ_stuName on stu(stuName);
Query OK, 0 rows affected (2.46 sec)
Records: 0  Duplicates: 0  Warnings: 0

# 方式2：修改表的方式添加索引
alter table stu add unique UQ_name (stuName);
Query OK, 0 rows affected, 1 warning (0.26 sec)
Records: 0  Duplicates: 0  Warnings: 1

# 方式3：创建表时添加索引
create table emp2(
    `id` int,
    `name` varchar(20),
    unique index UQ_name(name)
);
Query OK, 0 rows affected (2.06 sec)
```

## 10.4 删除索引
* 语法：
```SQL
drop index `索引名` on `表名`
```
* 例：
```SQL
drop index i_name on stuinfo;
Query OK, 0 rows affected (0.63 sec)
Records: 0  Duplicates: 0  Warnings: 0
```

## 10.5 创建索引的原则
> * 用于频繁搜索的列；
> * 该列用于排序；
> * 公共字段要创建索引；
> * 如果表中数据很少，不需要创建索引，MySQL搜索索引的时间比逐条搜索数据的时间还要长；
> * 如果一个字段上的数据只有几个不同的值（如：性别），不适合创建索引。

# 11.内置函数

## 11.1 数字类
* 随机数：
```SQL
select rand();
+-------------------+
| rand()            |
+-------------------+
| 0.612017822090553 |
+-------------------+
1 row in set (0.00 sec)
```
* 随机抽取2名学生的信息：
```SQL
select * from stu order by rand() limit 2;
+--------+--------------+--------+--------+---------+------------+------+------+
| stuNo  | stuName      | stuSex | stuAge | stuSeat | stuAddress | ch   | math |
+--------+--------------+--------+--------+---------+------------+------+------+
| s25305 | 诸葛丽丽     | 女     |     23 |       7 | 河南       |   72 |   56 |
| s25303 | 李斯文       | 女     |     22 |       2 | 北京       |   55 |   82 |
+--------+--------------+--------+--------+---------+------------+------+------+
2 rows in set (0.01 sec)
```
* 四舍五入：
```SQL
select round(4.5);
+------------+
| round(4.5) |
+------------+
|          5 |
+------------+
1 row in set (0.00 sec)
```
* 向上取整：
```SQL
select ceil(3.3);
+-----------+
| ceil(3.3) |
+-----------+
|         4 |
+-----------+
1 row in set (0.00 sec)
```
* 向下取整：
```SQL
select floor(3.9);
+------------+
| floor(3.9) |
+------------+
|          3 |
+------------+
1 row in set (0.00 sec)
```
* 小数位截取：
```SQL
select truncate(3.1415926, 3);
+------------------------+
| truncate(3.1415926, 3) |
+------------------------+
|                  3.141 |
+------------------------+
1 row in set (0.11 sec)
```

## 11.2 字符串类
* 转换大写：
```SQL
select ucase('Hello MySql');
+----------------------+
| ucase('Hello MySql') |
+----------------------+
| HELLO MYSQL          |
+----------------------+
1 row in set (0.00 sec)
```
* 转换小写：
```SQL
select lcase('Hello MySql');
+----------------------+
| lcase('Hello MySql') |
+----------------------+
| hello mysql          |
+----------------------+
1 row in set (0.00 sec)
```
* 左边截取：
```SQL
select left('abcde', 3);
+------------------+
| left('abcde', 3) |
+------------------+
| abc              |
+------------------+
1 row in set (0.00 sec)
```
* 右边截取：
```SQL
select right('abcde', 3);
+-------------------+
| right('abcde', 3) |
+-------------------+
| cde               |
+-------------------+
1 row in set (0.00 sec)
```
* 取子串：
```SQL
select substring('abcde', 2, 3);     # 从第二个字符开始截取3个字符
+--------------------------+
| substring('abcde', 2, 3) |
+--------------------------+
| bcd                      |
+--------------------------+
1 row in set (0.00 sec)
```
* 连接字符串：
```SQL
select concat('USA', 'MIT');
+----------------------+
| concat('USA', 'MIT') |
+----------------------+
| USAMIT               |
+----------------------+
1 row in set (0.00 sec)
```
* 将学生姓名和性别连接显示：
```SQL
select concat(stuName, '-', stuSex) from stu;
+------------------------------+
| concat(stuName, '-', stuSex) |
+------------------------------+
| 张秋丽-男                    |
| 李文才-男                    |
| 李斯文-女                    |
| 欧阳俊雄-男                  |
| 诸葛丽丽-女                  |
| 争青小子-男                  |
| 梅超风-女                    |
| Tom-男                       |
| Tabm-女                      |
+------------------------------+
9 rows in set (0.00 sec)
```
* NULL值替换：
```SQL
select coalesce(NULL, '空');
+-----------------------+
| coalesce(NULL, '空')  |
+-----------------------+
| 空                    |
+-----------------------+
1 row in set (0.00 sec)
```
* 将没有成绩的学生的成绩字段显示为缺考：
```SQL
select stuName, coalesce(writtenExam, '缺考'), coalesce(labExam, '缺考') from stuinfo i left join stumarks m on i.stuNo = m.stuNo;
+--------------+---------------------------------+-----------------------------+
| stuName      | coalesce(writtenExam, '缺考')   | coalesce(labExam, '缺考')   |
+--------------+---------------------------------+-----------------------------+
| 李斯文       | 80                              | 58                          |
| 李文才       | 50                              | 90                          |
| 欧阳俊雄     | 65                              | 50                          |
| 张秋丽       | 77                              | 82                          |
| 争青小子     | 56                              | 48                          |
| 梅超风       | 缺考                            | 缺考                        |
| 诸葛丽丽     | 缺考                            | 缺考                        |
+--------------+---------------------------------+-----------------------------+
7 rows in set (0.00 sec)
```
* 查看字符串占用的字节数：
```SQL
select length('锄禾日当午');     # UTF-8编码汉字占用3个字节
+---------------------------+
| length('锄禾日当午')      |
+---------------------------+
|                        15 |
+---------------------------+
1 row in set (0.00 sec)
```
* 查看字符串的字符数：
```SQL
mysql> select char_length('锄禾日当午');
+--------------------------------+
| char_length('锄禾日当午')      |
+--------------------------------+
|                              5 |
+--------------------------------+
1 row in set (0.00 sec)
```

## 11.3 时间类
* 获取Unix时间戳：
```SQL
select unix_timestamp();
+------------------+
| unix_timestamp() |
+------------------+
|       1542895353 |
+------------------+
1 row in set (0.01 sec)
```
* 将时间戳转换为日期时间格式：
```SQL
select from_unixtime(unix_timestamp());
+---------------------------------+
| from_unixtime(unix_timestamp()) |
+---------------------------------+
| 2018-11-22 14:02:57             |
+---------------------------------+
1 row in set (0.37 sec)
```
* 获取当前时间的日期时间格式：
```SQL
select now();
+---------------------+
| now()               |
+---------------------+
| 2018-11-22 14:03:05 |
+---------------------+
1 row in set (0.00 sec)
```
* 单独获取年，月，日，时，分，秒：
```SQL
select year(now()) as 年, month(now()) as 月, day(now()) as 日, hour(now()) as 时, minute(now()) as 分, second(now()) as 秒;
+------+------+------+------+------+------+
| 年   | 月   | 日   | 时   | 分   | 秒   |
+------+------+------+------+------+------+
| 2018 |   11 |   22 |   14 |   19 |   39 |
+------+------+------+------+------+------+
1 row in set (0.00 sec)
```
* 获取星期天数等信息：
```SQL
select dayname(now()) as 星期, monthname(now()), dayofyear(now()) as 本年第几天;
+----------+------------------+-----------------+
| 星期     | monthname(now()) | 本年第几天      |
+----------+------------------+-----------------+
| Thursday | November         |             326 |
+----------+------------------+-----------------+
1 row in set (0.00 sec)
```
* 日期计算：
```SQL
select datediff(now(), '2008-8-8');  # 当前时间与2008-8-8相差多少天
+-----------------------------+
| datediff(now(), '2008-8-8') |
+-----------------------------+
|                        3758 |
+-----------------------------+
1 row in set (0.00 sec)
```
* 将now()转换为日期和时间格式：
```SQL
# 方式1：
select convert(now(), date), convert(now(), time);   # 将now()转成日期和时间格式
+----------------------+----------------------+
| convert(now(), date) | convert(now(), time) |
+----------------------+----------------------+
| 2018-11-22           | 14:35:52             |
+----------------------+----------------------+
1 row in set (0.01 sec)

# 方式2：
select cast(now() as date), cast(now() as time);
+---------------------+---------------------+
| cast(now() as date) | cast(now() as time) |
+---------------------+---------------------+
| 2018-11-22          | 14:36:56            |
+---------------------+---------------------+
1 row in set (0.37 sec)
```

## 11.4 加密函数
```SQL
select md5('root'), sha('user');
+----------------------------------+------------------------------------------+
| md5('root')                      | sha('user')                              |
+----------------------------------+------------------------------------------+
| 63a9f0ea7bb98050796b649e85481845 | 12dea96fec20593566ab75692c9949596833adc9 |
+----------------------------------+------------------------------------------+
1 row in set (0.00 sec)
```

## 11.5 判断函数
* 语法：
```SQL
if(表达式, 值1, 值2)
```
* 显示学生的考试通过情况（所有成绩都大于等于60）：
```SQL
select *, if(writtenexam >= 60 && labexam >= 60, '通过', '未通过') as 结果 from stuinfo i, stumarks m where i.stuNo = m..stuNo;
+--------+--------------+--------+--------+---------+------------+---------+--------+-------------+---------+-----------+
| stuNo  | stuName      | stuSex | stuAge | stuSeat | stuAddress | examNo  | stuNo  | writtenExam | labExam | 结果      |
+--------+--------------+--------+--------+---------+------------+---------+--------+-------------+---------+-----------+
| s25303 | 李斯文       | 女     |     22 |       2 | 北京       | s271811 | s25303 |          80 |      58 | 未通过    |
| s25302 | 李文才       | 男     |     31 |       3 | 上海       | s271813 | s25302 |          50 |      90 | 未通过    |
| s25304 | 欧阳俊雄     | 男     |     28 |       4 | 天津       | s271815 | s25304 |          65 |      50 | 未通过    |
| s25301 | 张秋丽       | 男     |     18 |       1 | 北京       | s271816 | s25301 |          77 |      82 | 通过      |
| s25318 | 争青小子     | 男     |     26 |       6 | 天津       | s271819 | s25318 |          56 |      48 | 未通过    |
+--------+--------------+--------+--------+---------+------------+---------+--------+-------------+---------+-----------+
5 rows in set (0.00 sec)
```

# 12.预处理
* 语法：
```SQL
预处理语句：
prepare `预处理名` from 'SQL语句'

执行预处理：
execute `预处理名` [using 变量]
```
* 例：
```SQL
prepare stmt from 'select * from stuinfo where stuSex=? and stuAddress=?';        # ?表示占位符
Query OK, 0 rows affected (0.00 sec)
Statement prepared

set @sex = '男';      # 变量以@开头，通过set给变量赋值
Query OK, 0 rows affected (0.00 sec)

set @addr = '北京';
Query OK, 0 rows affected (0.00 sec)

execute stmt using @sex, @addr;      # 通过using传递变量替换占位符
+--------+-----------+--------+--------+---------+------------+
| stuNo  | stuName   | stuSex | stuAge | stuSeat | stuAddress |
+--------+-----------+--------+--------+---------+------------+
| s25301 | 张秋丽    | 男     |     18 |       1 | 北京       |
+--------+-----------+--------+--------+---------+------------+
1 row in set (0.00 sec)
```

# 13.存储过程

## 13.1 创建存储过程
* 语法：
```SQL
create procedure 存储过程名(参数)
begin
    SQL语句
end
```
* 例：
```SQL
delimiter //

create procedure proc()
begin
select * from stuinfo;
end //
Query OK, 0 rows affected (0.39 sec)

delimiter ;
```

## 13.2 调用存储过程
* 语法：
```SQL
call 存储过程名()
```
* 例：
```SQL
call proc();
+--------+--------------+--------+--------+---------+------------+
| stuNo  | stuName      | stuSex | stuAge | stuSeat | stuAddress |
+--------+--------------+--------+--------+---------+------------+
| s25301 | 张秋丽       | 男     |     18 |       1 | 北京       |
| s25302 | 李文才       | 男     |     31 |       3 | 上海       |
| s25303 | 李斯文       | 女     |     22 |       2 | 北京       |
| s25304 | 欧阳俊雄     | 男     |     28 |       4 | 天津       |
| s25305 | 诸葛丽丽     | 女     |     23 |       7 | 河南       |
| s25318 | 争青小子     | 男     |     26 |       6 | 天津       |
| s25319 | 梅超风       | 女     |     23 |       5 | 河北       |
+--------+--------------+--------+--------+---------+------------+
7 rows in set (0.00 sec)

Query OK, 0 rows affected (0.00 sec)
```

## 13.3 查看存储过程的信息
* 语法：
```SQL
show create procedure `存储过程名`\G
```
* 例：
```SQL
show create procedure proc\G;
*************************** 1. row ***************************
           Procedure: proc
            sql_mode: ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION
    Create Procedure: CREATE DEFINER=`root`@`localhost` PROCEDURE `proc`()
begin
select * from stuinfo;
end
character_set_client: utf8
collation_connection: utf8_general_ci
  Database Collation: utf8_general_ci
1 row in set (0.00 sec)
```

## 13.4 删除存储过程
* 语法：
```SQL
drop procedure [if exists] `存储过程名`
```
* 例：
```SQL
drop procedure proc;
Query OK, 0 rows affected (0.06 sec)
```
## 13.5 存储过程的参数：
例：查找学生的左右同桌
```SQL
delimiter //

create procedure proc(in name varchar(10)) # in表示输入参数
begin
declare seat tinyint;    # 使用declare声明局部变量
select stuSeat into seat from stuinfo where stuName=name;    # 使用into给变量赋值
select * from stuinfo where stuSeat=seat+1 or stuSeat=seat-1;
end //   
Query OK, 0 rows affected (0.04 sec)

call proc('李文才'); //
+--------+--------------+--------+--------+---------+------------+
| stuNo  | stuName      | stuSex | stuAge | stuSeat | stuAddress |
+--------+--------------+--------+--------+---------+------------+
| s25303 | 李斯文       | 女     |     22 |       2 | 北京       |
| s25304 | 欧阳俊雄     | 男     |     28 |       4 | 天津       |
+--------+--------------+--------+--------+---------+------------+
2 rows in set (0.00 sec)

Query OK, 0 rows affected, 1 warning (0.00 sec)

delimiter ;
```
* 输入数字，输出数字的平方：
```SQL
delimiter //

create procedure proc(in num int, out result int)    # out表示输出参数
begin
set result = num * num;
end //
Query OK, 0 rows affected (1.80 sec)

call proc(10, @result); //   # 可在传入时通过@定义全局变量
Query OK, 0 rows affected (0.00 sec)

select @result; //
+---------+
| @result |
+---------+
|     100 |
+---------+
1 row in set (0.00 sec)

delimiter ;
```