# Mysql

## 权限

用户与权限信息存储在 mysql 系统库中，一共用 5 张表去表达，分别是:user、db、tables_priv、columns_priv 和 procs_priv

- user 表:它属于全局层级，存储用户账户信息以及全局级别的权限，这些全局权限适用于系统中所有的数据 
- db 表:它属于数据库层级，适用于一个给定数据库中的所有目标
- tables_priv 表:它属于表层级，适用于一个给定表中的所有列
- columns_priv 表:它属于列层级，适用于一个给定表中的单一列
- procs_priv 表:它属于子程序层级，存放存储过程和函数级别的权限



## 访问控制

### 连接

- 用户连接时，MySQL 服务器首先匹配 mysql.user 表中的主机名、用户名和密码，匹配不到则拒绝当前的连接请求
- 检查 mysql.user 表的 max_connections 和 max_user_connections 字段值，如果超过设置的值范围则拒绝连接请求
- 检查 mysql.user 表的 SSL 安全连接配置，如果有配置 SSL，则需要用户提供证书且是合法的



不论我们是通过哪一种方式与 MySQL 服务器建立连接，MySQL 都会依次执行以上的 3 个权限校验工作。只有当 所有的检查都通过后，服务器才会与客户端建立连接。当连接建立之后，用户使用 SQL 语句执行增删改查 时，MySQL 服务器又需要做执行查询阶段的验证工作。



### 执行查询阶段

- 检查 mysql.user 表的 max_questions 和 max_updates 字段值，如果超过上限值，则拒绝执行 SQL 语句 
- 检查 mysql.user 表，如果拥有全局性权限，直接执行;否则，继续下一步检查
- 检查 mysql.db 表，如果拥有数据库级别的权限，则执行;否则，继续下一步检查
- 检查 mysql.tables_priv，mysql.columns_priv，mysql.procs_priv 表，如果拥有相应对象的权限，则执行;否 则，上报权限不足错误

从以上的检查过程可以看出，第一个检查是对属性规定的要求检查，后三个才是对权限的检查。同时，也可以看 出，MySQL 的权限检查是一个比较复杂的过程。所以，为了提高性能，MySQL 在启动时会把 5 张权限表加载到 内存中，是典型的空间换时间的思想。



## 数据类型

- 字符串类型:以 char、varchar、text 为代表，用于存储字符、字符串数据
- 日期 / 时间类型:以 date、time、datetime、timestamp 为代表，用于存储日期或时间，这种数据类型也是比 较难抉择的
- 数值类型:以 tinyint、int、bigint、float、double、decimal 为代表，用于存储整数或小数
- 二进制类型:以 tityblob、blob、mediumblob、longblob 为代表，用于存储二进制数据，适用场景最为受限

### 字符串类型

#### char

存储一些长度固定的数据，如果未达到指定长度，则会使用空格填充到指定长度

例如:手 机号码、身份证号等等。

#### varchar

可变长度的字符串，需要考虑边界问题，检索速度要慢于 char

例如:姓名、邮箱地址等等。

#### tinytext、text、mediumtext、longtext

- tinytext:最大长度是(2^8 - 1)个字符 

- text:最大长度是(2^16 - 1)个字符 

- mediumtext:最大长度是(2^24 - 1)个字符 

- longtext:最大长度是(2^32 - 1)个字符

建议数据量超过 500 个字符时，就应该考虑使用文本，文本类型不能有默认值，且在创建索引时需要指定前多少个 字符。

### 日期 / 时间类型

#### date

存储日期，存储范围是 ‘1000-01-01’ 到 ‘9999-12-31’

例如：出生日期

#### time

不仅可以表示一天中的时间，也可以用于表示两个时间的时间间隔。它的取值范围是 ‘- 838:59:59’ to ‘838:59:59’

#### datetime

‘1000-01-01 00:00:00.000000’ 到 ‘9999-12-31 23:59:59.999999’，是最常见，用途最广的数据类型

例如:存储数据插入时 间、订单完成时间等等。

#### timestamp

与 datetime 存储的数据格式是一样的，它的取值范围是:‘1970-01-01 00:00:01.000000’ UTC 到 ‘2038-01-19 03:14:07.999999’ UTC

timestamp 是与时区相关的，能够反映 “当前时间”,当插入时间时，会先转换为本地时区后再存储;查询时 间时，会转换为本地时区后再显示。



### 数值类型

#### 整数类型

主要支持 5 个整数类型:tinyint、smallint、mediumint、int、bigint

| 数据类型  | 占据空间 | 范围(有符号)   | 范围(无符号)             | 描述       |
| --------- | -------- | -------------- | ------------------------ | ---------- |
| tinyint   | 1 个字节 | -2^7-2^7-1     | 0-255                    | 小整数值   |
| smallint  | 2 个字节 | -2^15-2^15-1   | 0 - 65535                | 大整数值   |
| mediumint | 3 个字节 | -2^23-2^23-1   | 0 - 16777215             | 大整数值   |
| int       | 4 个字节 | -2^31-2^31-1   | 0 - 4294967295           | 大整数值   |
| bigint    | 8 个字节 | -2^63 - 2^63-1 | 0 - 18446744073709551615 | 极大整数值 |

关于整数类型还有一个特性:显示宽度

```sql
`a` bigint(20) NOT NULL COMMENT 'a', 
`b` int(11) NOT NULL COMMENT 'b'
```

20 和 11 就是可选的显示宽度，这会让 MySQL 对 SQL 标准进行扩展，当从数据库检索一个值时，可以把 这个值延长到指定的宽度。例如，这里的 b 定义的类型为 int (11)，就可以保证 b 这一列少于 11 个字符宽度时自 动使用空格填充。但同时，需要注意，定义宽度并不会影响字段的大小和存储值的取值范围。

#### 浮点类型

支持两个浮点类型:float、double。

- float 用于表示单精度浮点数值，占用 4 个字节
- double 用于表 示双精度浮点数值，占用 8 个字节
- float (M, D):其中 M 定义显示长度，D 定义小数位数。但是它们是可选的，且默认值是 float (10, 2)，2 是小 数的位数，10 是数字的总长(包括小数)。它的小数精度可以到 24 个浮点。
- double (M, D):M 和 D 的含义与 float 是相同的，默认值是 double (16, 4)。它的小数精度可以达到 53 位。

#### 定点类型

MySQL 中的 decimal 被称为定点数据类型，由于它保存的是精确值,所以它通常用于精度要求非常高的计算中



由于 decimal 需要比较大的空间和计算开销，它的计算效率也就没有 float 和 double 那么高，所以应该只有要求精 确计算的场景下才考虑去使用 decimal。



### 二进制类型

二进制数据类型理论上可以存储任何数据，可以是文本数据，也可以存储图像或者其他多媒体数据。二进制数据类

型相对于其他的数据类型来说，使用频率是比较低的

- tityblob:最大支持 255 字节 

- blob:最大支持 64KB 

- mediumblob:最大支持 16MB 

- longblob:最大支持 4GB



## Schema设计

Schema 设计指的是对数据表的的设计

### 设计原则

- 使用小写的名称，且只有英文字母
- 取一个有意义的名称，单词之间使用下划线连接
- 记住 “够用且尽量小” 的原则
- 不要使用物理外键，用代码去管理这种关系
- 表一定要有主键
- 保持一致的字符集:库、表、数据列的字符集都应该是一致的，统一为 utf8 或 utf8mb4



#### 创建表

```sql
CREATE TABLE `ad_unit` (
`id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
`user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '标记当前记录所属用户', `unit_name` varchar(48) NOT NULL COMMENT '推广单元名称',
`unit_status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '推广单元状态: 0-正常, 1-失效', `position_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '广告位类型(1,2,3)',
`budget` bigint(20) NOT NULL COMMENT '预算(单位: 元)',
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='广告-推广单元表';


+----+---------+-----------------+-------------+---------------+--------+
| id | user_id | unit_name | unit_status | position_type | budget | 
+----+---------+-----------------+-------------+---------------+--------+
```





## 聚合函数

常用的聚合 函数有五个:AVG、COUNT、MIN、MAX、SUM。

| 语法                  | 功能               | 备注                                    |
| --------------------- | ------------------ | --------------------------------------- |
| AVG ([DISTINCT] expr) | 返回 expr 的平均值 | DISTINCT 选项用于去除字段值重复的行记录 |
| COUNT(expr)           | 统计表中的行数     |                                         |
| MIN ([DISTINCT] expr) | 返回 expr 的最小值 |                                         |
| MAX ([DISTINCT] expr) | 返回 expr 的最大值 |                                         |
| SUM([DISTINCT] expr)  | 返回 expr 的合计值 |                                         |

一些共性:

- 每个聚合函数接受一个参数，参数可以是数据表列，也可以是函数表达式 
- 默认情况下，聚合函数会忽略列值为 NULL 的行，不参与计算 
- 聚合函数不允许嵌套，例如:COUNT(SUM(expr)) 是不合法的 
- 一次查询中可以出现多个聚合函数，例如:SELECT MAX(expr), MIN(expr) FROM ...

表`worker`如下：

```sql
mysql> SELECT id, type, name, salary FROM worker; 
+----+------+--------+--------+
|id|type|name |salary|
+----+------+--------+--------+
|1 |A 	|tom   |1800| 
|2 |B 	|jack  |2100| 
|3 |C 	|pony  |NULL| 
|4 |B 	|tony  |3600|
|5 |B 	|marry |1900| 
|6 |C 	|tack  |1200| 
|7 |A 	|tick  |NULL|
|8 |B 	|clock |2000| 
|9 |C 	|noah  |1500| 
|10|C 	|jarvis|1800
| +----+------+--------+--------+
```

#### AVG

AVG 只适用于数值类型的列，因为对于像日期、字符串等类型求平均本身就是没有意义的。

```sql
mysql> SELECT AVG(salary) FROM worker WHERE id < 3; 
+-------------+
| AVG(salary) |
+-------------+
| 1950.0000   | 
+-------------+
```

#### COUNT

COUNT(n)、COUNT(*)、COUNT(expr)、COUNT(DISTINCT expr)

COUNT(n) 中的 n 可以是任何整数或小数，它与 COUNT(*) 的查询结果是一样的

```sql
mysql> SELECT COUNT(0), COUNT(1), COUNT(9.9), COUNT(*) FROM worker; 
+----------+----------+------------+----------+
| COUNT(0) | COUNT(1) | COUNT(9.9) | COUNT(*) | 
+----------+----------+------------+----------+
| 10			 | 10				| 10				 | 10				| 
+----------+----------+------------+----------+
```

- COUNT(n) 和 COUNT(*) 统计的总行数是包含 NULL 值的

- COUNT(expr) 和 COUNT(DISTINCT expr) 由于需要传入列作为参数，它们统计的是非 NULL 的行数。

```sql
mysql> SELECT COUNT(salary), COUNT(DISTINCT salary) FROM worker; 
+---------------+------------------------+
| COUNT(salary) | COUNT(DISTINCT salary) | 
+---------------+------------------------+
|8							|7											 | 
+---------------+------------------------+
```



- COUNT(n) 和 COUNT(*) 用于统计表中的总行数，不关心列值是否为 NULL 
- COUNT(expr) 用于统计列值非 NULL 的行记录数
- COUNT(DISTINCT expr) 用于统计列值不同且非 NULL 的行记录数



#### MIN、MAX

不同于 AVG，它们的适用范围不只是数值类型，日期类型、 字符串类型也同样是允许的。

```sql
mysql> SELECT MIN(salary), MAX(salary) FROM worker; 
+-------------+-------------+
| MIN(salary) | MAX(salary) |
+-------------+-------------+
| 1200 			  | 3600   		  | 
+-------------+-------------+
```



#### SUM

只能用于数值类型的列，且会忽略值为 NULL 的列



## 分组 GROUP BY

数据根据某一列或者某几列分类,GROUP BY 结合聚合函数就可以实现将表数据分类再汇总的效果

语法如下：

```sql
SELECT
<列名1>,
<列名2>...... FROM
<表名> WHERE
......
GROUP BY
<列名1>, <列名2>......;
```



#### 按照 type 分组对数据进行统计

```sql
mysql> SELECT type,SUM(salary) as sum_s FROM worker GROUP BY type ORDER BY sum_s DESC;
+------+-------+
| type | sum_s |
+------+-------+
| B    |  9600 |
| C    |  4500 |
| A    |  1800 |
+------+-------+
```



#### 对分组聚合结果进行排序

```sql
mysql> SELECT type, AVG(salary), COUNT(1), MIN(salary), MAX(salary), SUM(salary) FROM worker GROUP BY type DESC;
+------+-------------+----------+-------------+-------------+-------------+
| type | AVG(salary) | COUNT(1) | MIN(salary) | MAX(salary) | SUM(salary) |
+------+-------------+----------+-------------+-------------+-------------+
| C    |   1125.0000 |        4 |           0 |        1800 |        4500 |
| B    |   2400.0000 |        4 |        1900 |        3600 |        9600 |
| A    |    900.0000 |        2 |           0 |        1800 |        1800 |
+------+-------------+----------+-------------+-------------+-------------+
3 rows in set, 1 warning (0.00 sec)
```



#### 对分组结果进行过滤 - HAVING

这里过滤的是分组后的聚合结果，而不是数据表中的原始记录。

想要 获取 SUM(salary) 大于 4000 的分组，可以这样做:

```sql
mysql> SELECT type, AVG(salary), COUNT(1), SUM(salary) FROM worker GROUP BY type HAVING SUM(salary) > 4000;
+------+-------------+----------+-------------+
| type | AVG(salary) | COUNT(1) | SUM(salary) |
+------+-------------+----------+-------------+
| B    |   2400.0000 |        4 |        9600 |
| C    |   1125.0000 |        4 |        4500 |
+------+-------------+----------+-------------+
```

HAVING 的使用方法与 WHERE 是相似的

- WHERE 子句在分组前对原始记录进行过滤 

- HAVING 子句在分组后对记录进行过滤





