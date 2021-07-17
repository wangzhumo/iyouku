# 高级查询 - 连接、联合、子查询

## 表创建

```sql
-- 创建 imooc_user 表
CREATE TABLE `imooc_user`
(
    `user_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户 id',
    `name`    char(64)   NOT NULL DEFAULT '' COMMENT '姓名',
    `age`     int(11)    NOT NULL DEFAULT '0' COMMENT '年龄',
    PRIMARY KEY (`user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8 COMMENT ='慕课网用户信息表';
  
+-------+-----+---+
|user_id|name |age|
+-------+-----+---+
|1      |qinyi|19 |
|2      |abc  |32 |
|3      |xyz  |30 |
|4      |mno  |29 |
+-------+-----+---+



-- 创建 imooc_course 表
CREATE TABLE `imooc_course`
(
    `id`      bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id` bigint(20) NOT NULL COMMENT '用户 id',
    `cname`   char(64)   NOT NULL DEFAULT '' COMMENT '课程名',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8 COMMENT ='慕课网课程信息表';
  
+--+-------+------+
|id|user_id|cname |
+--+-------+------+
|1 |1      |广告系统  |
|2 |1      |优惠券系统 |
|3 |1      |卡包系统  |
|4 |2      |Java  |
|5 |2      |Python|
|6 |3      |MySQL |
|7 |5      |Linux |
+--+-------+------+  
```



## 连接查询

它是将多个表(绝大多数情况下是两张表)联合起来查询，连接的方式 一共有四种:内连接、外连接、自然连接和交叉连接。连接查询的最终目的是在一次查询中获取多张表的数据。

### 内连接

内连接是有条件匹配的连接，多个表之间依据给定的条件进行连接，并保留符合匹配结果的记录。

语法如下：

```sql
SELECT col1, col2 FROM left [INNER] JOIN right ON left.colx = right.coly
```

从左表(SQL 语句中位于左边的表)中取出一条记录，与右表中的所有记录去匹配，保 留匹配成功的记录，并拼接打印。

```sql
mysql> SELECT * FROM imooc_user AS user INNER JOIN imooc_course AS course ON user.user_id = course.user_id;

| user\_id | name | age | id | user\_id | cname |
| :--- | :--- | :--- | :--- | :--- | :--- |
| 1 | qinyi | 19 | 1 | 1 | 广告系统    |
| 1 | qinyi | 19 | 2 | 1 | 优惠券系统 	|
| 1 | qinyi | 19 | 3 | 1 | 卡包系统    |
| 2 | abc   | 32 | 4 | 2 | Java    		|
| 2 | abc   | 32 | 5 | 2 | Python 		|
| 3 | xyz   | 30 | 6 | 3 | MySQL 			|

```

ON 子句之后也可以指定单表的条件

例如：

```sql
SELECT * FROM imooc_user AS user JOIN imooc_course AS course ON user.user_id = course.user_id AND user.user_id <= 1
```



### 外连接

内连接只会保留两个表中完全匹配的记录，而外连接则不同，不论 “主表” 符不符合匹配 条件，记录都将被保留。

```sql
-- 左外连接
SELECT col1, col2 FROM left LEFT JOIN right ON left.colx = right.coly

-- 右外连接
SELECT col1, col2 FROM left RIGHT JOIN right ON left.colx = right.coly
```

举例如下：

```sql
-- 左外连接  （imooc_user）
SELECT * FROM imooc_user AS user LEFT JOIN imooc_course AS course ON user.user_id = course.user_id;

+-------+-----+---+----+-------+------+
|user_id|name |age|id  |user_id|cname |
+-------+-----+---+----+-------+------+
|1      |qinyi|19 |1   |1      |广告系统  |
|1      |qinyi|19 |2   |1      |优惠券系统 |
|1      |qinyi|19 |3   |1      |卡包系统  |
|2      |abc  |32 |4   |2      |Java  |
|2      |abc  |32 |5   |2      |Python|
|3      |xyz  |30 |6   |3      |MySQL |
|4      |mno  |29 |NULL|NULL   |NULL  |
+-------+-----+---+----+-------+------+

-- 右外连接 - 保留右表 （imooc_course）
SELECT * FROM imooc_user AS user RIGHT JOIN imooc_course AS course ON user.user_id = course.user_id;

+-------+-----+----+--+-------+------+
|user_id|name |age |id|user_id|cname |
+-------+-----+----+--+-------+------+
|1      |qinyi|19  |1 |1      |广告系统  |
|1      |qinyi|19  |2 |1      |优惠券系统 |
|1      |qinyi|19  |3 |1      |卡包系统  |
|2      |abc  |32  |4 |2      |Java  |
|2      |abc  |32  |5 |2      |Python|
|3      |xyz  |30  |6 |3      |MySQL |
|NULL   |NULL |NULL|7 |5      |Linux |
+-------+-----+----+--+-------+------+
```

原始数据中，`imooc_user` 中的用户1-4  但是 `imooc_course` 中没有 4 号用户

`imooc_course`中`7` 是没有`imooc_user`对应用户的

以上两种查询，通过观察就能看出,是以左/右为基础，增加了另一张表的数据

### 交叉连接

交叉连接也被称为 “无条件连接”，它会将左表的每一条记录与右表的每一条记录进行连 接，结果中的列数等于两表列数之和，行数等于两表行数之积。

```sql
SELECT col1, col2 FROM left CROSS JOIN right
```

## 联合查询

UNION，它将两个或多个查询的结果拼接在一起返回。

正是由于它会将不同的查询结果拼接 在一起，所以，每一个查询结果的字段数必须是相同的。特殊的是，拼接过程并不会在意数据类型(但还是强烈推 荐对应列的数据类型一致，否则，很容易在代码中出现类型错误)。

UNION 有个特性，它会去除重复的行(所有的字 段都相同)，如果想要完整的数据，则需要加上 ALL 选项

### 语法

```sql
SELECT user_id FROM imooc_user UNION SELECT user_id FROM imooc_course 

SELECT user_id FROM imooc_user UNION ALL SELECT user_id FROM imooc_course
```



### 对结果进行排序

- 对于单个查询结果排序的话，需要将它的 SELECT 语句用括号括起来，同时配合 LIMIT 使用 (如果不配合 LIMIT，会被语法分析器优化时去除，导致排序失效)
- 最终的查询结果进行排序，在最后一个 SELECT 语句之后使用 ORDER BY 即可

```sql
(SELECT name, age FROM imooc_user ORDER BY age DESC LIMIT 2) UNION SELECT user_id, cname FROM imooc_course;
```

## 子查询

按照它出现的位置(使用的关键字)可以分为三类:FROM 子查询、WHERE 子查询和 EXISTS 子查询

### FROM 子查询

子查询是跟在 FROM 子句之后的，它的语义是:先查出 “一张表”，再去查询这张表。

```sql
-- 给子查询派生出的表指定 user 别名
mysql> SELECT name, age FROM (SELECT * FROM imooc_user WHERE name = 'qinyi') AS user; 

+-------+-----+
|name |age|
+-------+-----+
|qinyi| 19|
+-------+-----+
 
```

FROM 子查询常常用于将复杂的查询分解，将复杂的查询条件拆解到多次查询中去。

### WHERE 子查询

WHERE 子查询是跟在 WHERE 条件中的，它的语义是:先根据条件筛选，再根据条件查询。

```sql
SELECT * FROM imooc_user WHERE user_id IN (SELECT user_id FROM imooc_course WHERE user_id < 3);

SELECT * FROM imooc_user WHERE user_id = (SELECT user_id FROM imooc_course WHERE cname LIKE '广告%');
```

WHERE 子查询的常见用法是 把其他表的查询结果当做当前表的查询条件进行二次查询，非常类似于 JOIN 的思想。

###  EXISTS 子查询

它所表达的语义是:存在才触发

```sql
SELECT * FROM imooc_user WHERE EXISTS(SELECT * FROM imooc_course WHERE cname LIKE '%JAVA%') AND user_id = 1

+---------+-------+-----+
|user_id  |name   |age  |
+---------+-------+-----+
| 1       |qinyi  | 19  |
+---------+-------+-----+

-- 不存在课程名包含 PHP 的课程，所以，子查询不存在
SELECT * FROM imooc_user WHERE EXISTS(SELECT * FROM imooc_course WHERE cname LIKE '%PHP%') AND user_id = 1; 

Empty set (0.00 sec)
```

