# Redis 

## NoSQL

`Not Only SQL`

### 适用场景

- 一些高并发的读写操作
- 大量数据的读写操作
- 对数据扩展性有要求

### 不适用

- 需要事务的支持
- 需要结构化的查询，有复杂的相关关系

### 常用的NoSQL

- Memcache   不能持久化
- Redis    支持多种数据结构，可以持久化
- mongoDB    文档型 json，二进制,大文档



## Redis简介

Redis 是单线程 + 多路IO复用

可以用来作为数据库、缓存和消息队列

Redis 比其他 key-value 缓存产品有以下三个特点：

- Redis 支持数据的持久化，可以将内存中的数据保存在磁盘中，重启的时候可以再次加载到内存使用。
- Redis 不仅支持简单的 key-value 类型的数据，同时还提供 list，set，zset，hash 等数据结构的存储。
- Redis 支持主从复制，即 master-slave 模式的数据备份。

Redis 主要由有两个程序组成：

- Redis 客户端 redis-cli
- Redis 服务器 redis-server



## Redis通用的key操作

### 1.查询

#### 所有key

`keys *`

```shell
127.0.0.1:6379[1]> keys *
1) "key1"
2) "key2"
3) "key3"
```

#### Key是否存在

`exists {key}`

```shell
127.0.0.1:6379[1]> exists key1
(integer) 1   # 存在
127.0.0.1:6379[1]> exists key4
(integer) 0   # 不存在
127.0.0.1:6379[1]>
```

#### 查询Key的类型

`type {key}`

```shell
127.0.0.1:6379[1]> type key1
string
```

#### 查询key

`get {key}`

#### 查询是否过期

`ttl {key}`

```shell
127.0.0.1:6379[1]> set key2 value2
OK
127.0.0.1:6379[1]> ttl key2    
(integer) -1    # 没有设置key2的过期时间，所以一直存在
127.0.0.1:6379[1]> expire key2 10
(integer) 1	
127.0.0.1:6379[1]> ttl key2
(integer) 8	    # 剩余时间	
127.0.0.1:6379[1]> ttl key2
(integer) -2	# 已经过期
127.0.0.1:6379[1]>
```

- -1   永不过期
- -2    过期
- other   剩余时间



### 2.删除

#### 删除key

`del {key}`

```shell
127.0.0.1:6379[1]> del key1
(integer) 1
```

#### 非阻塞删除

先把`key`从`keyspace`元数据中删除，真正的删除异步操作删除

`unlink {key}`

#### 过期时间

`expire {key} {seconds}`

```shell
127.0.0.1:6379[1]> expire key2 10
(integer) 1
127.0.0.1:6379[1]> get key2
"value2"
127.0.0.1:6379[1]> get key2
(nil)
```





## Redis中的数据类型

### - String

Redis中的基本类型，一个Redis中的string字符串最大可以是**512M**

#### 插入

-  `set {key} {value}` ,key唯一，如果重复会覆盖

-  `setnx {key} {value} ` 当key不存在的时候插入key，value

-  `setex {key} {过期时间} {value}  ` 插入并设置过期时间

-  `mset {key1} {value1} ... {keyN} {valueN}`   同时设置

-  `msetnx {key1} {value1} ... {keyN} {valueN}`  同时setnx

  当**所有key都不存在**的情况下，才会发生赋值操作
  
- `setrange {key}  {startIndex} {value}` 插入指定范围

  会把`key`中的值覆盖，从`startIndex`开始直到`startIndex + length(value)`

#### 查询

- `get {key}`  查询

-  `mget {key1}...{keyN} `  同时查询 

-  `strlen {key}`    长度

- `getrange {key} {startIndex} {endIndex}`  获取指定范围 

  ```shell
  127.0.0.1:6379[1]> set key3 abcdefg
  OK
  127.0.0.1:6379[1]> getrange key3 0 3
  "abcd"
  ```

  这个取值的范围是**前包后包**的取值

#### 修改

-  `incr {key}`  可以对数字类型的key，做自增操作
-  `decr {key}`   可以对数字类型的key，减去1
-  `incrby/decrby {key} {步长}`    可以对数字类型的key，进行-+ 步长的操作
-  `append {key} {value} `   追加数据

#### 其他

-  `getset {key} {value}`  先查询key值，并赋值一个新的值



### - List

单键多值，它的底层是一个**双向链表**，所以查询效率比较低，增删快

#### 插入

- `lpush/rpush {key} {value1} ... {valueN}`  

  从左/右边添加一个key

  ```shell
  127.0.0.1:6379[1]> lpush keyx a b c d e f g
  (integer) 7
  127.0.0.1:6379[1]> lrange keyx 0 3
  1) "g"
  2) "f"
  3) "e"
  4) "d"
  ```

  注意观察它的**顺序**

- `rpoplpush {key1} {keyX} `

  列表右边获取`key1`的一个值,并将之歌值插入到`keyX`列表的左边

  ```shell
  127.0.0.1:6379[1]> rpush keyy  x y z
  (integer) 3
  127.0.0.1:6379[1]> rpoplpush keyx keyy
  "a"
  127.0.0.1:6379[1]> lrange keyy 0 3
  1) "a"
  2) "x"
  3) "y"
  4) "z"
  ```

  我们可以看到，keyx 中的最后一个元素被插入了 keyy 的左边

- `linsert {key} before {value} {newValue}`    在value前插入newValue值

- `set {key} {index} {value}`   覆盖在index位置的值为value

#### 查询

- `lpop/rpop`  

  从左/右边获取一个key ，**值在键在，值空键亡**

- `lrange {key} {start} {end}`   按照索引获取元素

- `lindex {key} {index}`  按照index取key中的元素

- `llen {key}`  获取长度

#### 删除

- `lrem {key}{count} {value} `   删除key中 count 个为value的元素

  ```shell
  127.0.0.1:6379[1]> lrange keyx 0 7
  1) "d"
  2) "b"
  3) "c"
  4) "b"
  127.0.0.1:6379[1]> lrem keyx 1 b
  (integer) 1
  
  # 可以看到,只是删除了一个 `b` 
  127.0.0.1:6379[1]> lrange keyx 0 7
  1) "d"
  2) "c"
  3) "b"
  ```

  
