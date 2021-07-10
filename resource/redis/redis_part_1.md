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

> [Redis 命令参考 — Redis 命令参考 (redisfans.com)](http://doc.redisfans.com/)



## Redis 配置

#### 简介

Redis中使用byte为大小的单位

同时，不区分大小写

```shell
# Redis configuration file example.
#
# Note that in order to read the configuration file, Redis must be
# started with the file path as first argument:
#
# ./redis-server /path/to/redis.conf

# Note on units: when memory size is needed, it is possible to specify
# it in the usual form of 1k 5GB 4M and so forth:
#
# 1k => 1000 bytes
# 1kb => 1024 bytes
# 1m => 1000000 bytes
# 1mb => 1024*1024 bytes
# 1g => 1000000000 bytes
# 1gb => 1024*1024*1024 bytes
#
# units are case insensitive so 1GB 1Gb 1gB are all the same.
```

#### INCLUDES 

可以在这里，链接其他文件中的配置

```shell
################################## INCLUDES ###################################
# include /path/to/local.conf
# include /path/to/other.conf
# include /path/to/fragments/*.conf
```

#### NETWORK

```shell
################################## NETWORK #####################################
# Examples:
#
# bind 192.168.1.100 10.0.0.1     # listens on two specific IPv4 addresses
# bind 127.0.0.1 ::1              # listens on loopback IPv4 and IPv6
# bind * -::*                     # like the default, all available interfaces
#
# You will also need to set a password unless you explicitly disable protected
# mode.
# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
bind 127.0.0.1 -::1
protected-mode yes

# Accept connections on the specified port, default is 6379 (IANA #815344).
# If port 0 is specified Redis will not listen on a TCP socket.
port 6379

# TCP listen() backlog.
tcp-backlog 511

# Unix socket.
# unixsocket /run/redis.sock
# unixsocketperm 700

# Close the connection after a client is idle for N seconds (0 to disable)
timeout 0

# TCP keepalive.
tcp-keepalive 300
```

- `bind 127.0.0.1 -::1    `      只能本地访问，如果需要remote访问，注释掉

- `protected-mode yes`          需要remote访问，则关闭保护模式

- `port 6379`     端口号

- `tcp-backlog 511`   连接队列的总和

  backlog连接队列  = 未完成三次握手队列 + 已经完成三次握手队列

- `timeout 0`       超时的时间，连接有效时间,单位是: 秒

- `tcp-keepalive 300`            Tcp心跳检测时间，300秒

#### GENERAL 

```shell
################################# GENERAL #####################################

# By default Redis does not run as a daemon. Use 'yes' if you need it.
# Note that Redis will write a pid file in /var/run/redis.pid when daemonized.
# When Redis is supervised by upstart or systemd, this parameter has no impact.
daemonize no

# If you run Redis from upstart or systemd, Redis can interact with your
# supervision tree. Options:
#   supervised no      - no supervision interaction
#   supervised upstart - signal upstart by putting Redis into SIGSTOP mode
#                        requires "expect stop" in your upstart job config
#   supervised systemd - signal systemd by writing READY=1 to $NOTIFY_SOCKET
#                        on startup, and updating Redis status on a regular
#                        basis.
#   supervised auto    - detect upstart or systemd method based on
#                        UPSTART_JOB or NOTIFY_SOCKET environment variables
#
# supervised auto

# If a pid file is specified, Redis writes it where specified at startup
# and removes it at exit.
# Note that on modern Linux systems "/run/redis.pid" is more conforming
# and should be used instead.
pidfile /var/run/redis_6379.pid

# Specify the server verbosity level.
# This can be one of: debug verbose notice warning
loglevel notice

# Specify the log file name. default logs will be sent to /dev/null
logfile ""

# To enable logging to the system logger
# syslog-enabled no

# Specify the syslog identity.
# syslog-ident redis

# Specify the syslog facility. Must be USER or between LOCAL0-LOCAL7.
# syslog-facility local0

# To disable the built in crash log, which will possibly produce cleaner core
# dumps when they are needed, uncomment the following:
#
# crash-log-enabled no

# To disable the fast memory check that's run as part of the crash log, which
# will possibly let redis terminate sooner, uncomment the following:
#
# crash-memcheck-enabled no

# Set the number of databases. The default database is DB 0, you can select
# a different one on a per-connection basis using SELECT <dbid> where
# dbid is a number between 0 and 'databases'-1
databases 16

# By default Redis shows an ASCII art logo only when started to log to the
# standard output and if the standard output is a TTY and syslog logging is
# disabled. Basically this means that normally a logo is displayed only in
# interactive sessions.
#
# However it is possible to force the pre-4.0 behavior and always show a
# ASCII art logo in startup logs by setting the following option to yes.
always-show-logo no

# By default, Redis modifies the process title (as seen in 'top' and 'ps') to
# provide some runtime information. It is possible to disable this and leave
# the process name as executed by setting the following to no.
set-proc-title yes

# When changing the process title, Redis uses the following template to construct
# the modified title.
#
# Template variables are specified in curly brackets. The following variables are
# supported:
#
# {title}           Name of process as executed if parent, or type of child process.
# {listen-addr}     Bind address or '*' followed by TCP or TLS port listening on, or
#                   Unix socket if only that's available.
# {server-mode}     Special mode, i.e. "[sentinel]" or "[cluster]".
# {port}            TCP port listening on, or 0.
# {tls-port}        TLS port listening on, or 0.
# {unixsocket}      Unix domain socket listening on, or "".
# {config-file}     Name of configuration file used.
#
proc-title-template "{title} {listen-addr} {server-mode}"
```

- `daemonize no`    	守护进程，后台启动，一般就yes
- `pidfile /var/run/redis_6379.pid`                 进程号写入的文件
- `loglevel`   日志级别
- `logfile`      日志路径
- `databases`    数据库数量

#### SECURITY 

```shell
################################## SECURITY ##################################
# ACL LOG
#
# The ACL Log tracks failed commands and authentication events associated
# with ACLs. The ACL Log is useful to troubleshoot failed commands blocked
# by ACLs. The ACL Log is stored in memory. You can reclaim memory with
# ACL LOG RESET. Define the maximum entry length of the ACL Log below.
acllog-max-len 128

# Using an external ACL file
#
# aclfile /etc/redis/users.acl

# Clients will still authenticate using AUTH <password> as usually
# if they follow the new protocol: both will work.
#
# The requirepass is not compatible with aclfile option and the ACL LOAD
# command, these will cause requirepass to be ignored.
#
requirepass psdcustom

```

- `requirepass {password}`

#### CLIENTS

```shell
################################### CLIENTS ####################################

# Set the max number of connected clients at the same time.
#
# maxclients 10000
```

- `maxclients `  最大连接数



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



