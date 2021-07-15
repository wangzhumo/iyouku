# 消息中间件

## 简介

消息中间件就是软件和软件之间发送消息的软件


### 解决问题

- 减少业务调用链，及时返回，减少等待时间
- 部分组件故障不会让整个业务崩溃
- 可以有效的对高峰期起到缓冲作用
- 不会像异步调用的实现方式一样，产生大量线程，内存资源占用少

### 工程中的作用

- 异步处理

- 系统解耦

- 消息的广播

- 流量削峰/控流

- 日志/消息的收集

- 最终一致性

  发出消息后不需要管理，消息接收方会最终处理掉

## 主流的消息中间件

- Apache ActiveMQ  比较老的一个消息中间件
  - Java开发
  - 支持连接多QpenWire/STOMP/REST/XMPP/AMQP
  - 可以使用JDBC连接数据库
  - 监控/重试
  - 不适用于上千个队列

- RabbitMQ  主流使用，通用性好
  - 支持发送确认，可靠性高
  - 基于`Erlang`，支持高并发
  - 社区活跃，文档齐全
  - `AMQP`协议
  - 代理模式下，中央节点增加延迟，影响性能

- Apache RocketMQ 性能强大，但对架构/运维有要求
  - 支持单机1w以上的持久化队列
  - 内存与磁盘都有数据，性能和可靠性强
  - 能保证亿级的能力，可以保障消息顺序
  - 社区活跃但是不及`RabbitMQ`

- Apache Kafka 一般用于大数据和日志，对大量数据支持好，可靠性弱

  - 原生分布式系统，对大数据友好
  - 零拷贝，减少IO
  - 快速持久化，性能高，稳定性好
  - 支持数据批量发送/拉取
  - 但是可靠性不好，实时性不定，不能重试，当单机超过64队列，性能劣化

  

## 安装

> [Downloading and Installing RabbitMQ — RabbitMQ](https://www.rabbitmq.com/download.html)

```shell
docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management

# -p 5672:5672	   RabbitMQ 连接端口
# -p 15672:15672   这个是管理控制台端口

╭─wangzhumo at Wangzhumo in /Applications/Docker.app/Contents using 21-07-15 - 14:16:40
╰─○ docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
Unable to find image 'rabbitmq:3-management' locally
3-management: Pulling from library/rabbitmq
a31c7b29f4ad: Pull complete
0ea5da5fa011: Pull complete
2d9925bd5669: Pull complete
56f5b6dce95d: Pull complete
ae74659cb465: Pull complete
e20048198e4f: Pull complete
9b824a4a94bf: Pull complete
a708661d5d9a: Pull complete
dba4fca7ba03: Pull complete
Digest: sha256:7f0fe23cf3a859298adacebec85c23a1127c1a0515220ed126dae61379096936
Status: Downloaded newer image for rabbitmq:3-management
```

等待第一次下载完成，就会直接开始运行

```shell
Starting RabbitMQ 3.8.19 on Erlang 24.0.3 [jit]
 Copyright (c) 2007-2021 VMware, Inc. or its affiliates.
 Licensed under the MPL 2.0. Website: https://rabbitmq.com

  ##  ##      RabbitMQ 3.8.19
  ##  ##
  ##########  Copyright (c) 2007-2021 VMware, Inc. or its affiliates.
  ######  ##
  ##########  Licensed under the MPL 2.0. Website: https://rabbitmq.com

  Erlang:      24.0.3 [jit]
  TLS Library: OpenSSL - OpenSSL 1.1.1k  25 Mar 2021

  Doc guides:  https://rabbitmq.com/documentation.html
  Support:     https://rabbitmq.com/contact.html
  Tutorials:   https://rabbitmq.com/getstarted.html
  Monitoring:  https://rabbitmq.com/monitoring.html

  Logs: <stdout>

  Config file(s): /etc/rabbitmq/rabbitmq.conf

  Starting broker...2021-07-15 06:24:15.428 [info] <0.273.0>
 node           : rabbit@2930c41a60ea
 home dir       : /var/lib/rabbitmq
 config file(s) : /etc/rabbitmq/rabbitmq.conf
 cookie hash    : vVmyiI/oawcjtRgBCA1RUw==
 log(s)         : <stdout>
 database dir   : /var/lib/rabbitmq/mnesia/rabbit@2930c41a60ea
```

其中说明了

版本：RabbitMQ 3.8.19 on Erlang 24.0.3
Config: Config file(s): /etc/rabbitmq/rabbitmq.conf

以及其他一些启动后的配置说明.



## 启动管理器

> http://127.0.0.1:15672/

### 新建Exchange

|              |                                                              |
| -----------: | ------------------------------------------------------------ |
|        Name: | *                                                            |
|        Type: | direct                    fanout                    headers                    topic |
|  Durability: | Durable(持久化)          Transient                           |
| Auto delete: | No          Yes                                              |
|    Internal: | No          Yes                                              |
|   Arguments: |                                                              |

- name     新建交换机的名字
- Type      交换机的类型
- Durability    是否需要持久化
- Auto delete   是否在无人绑定的情况下，关闭这个交换机

### 新建Queue

|              |                                    |
| -----------: | ---------------------------------- |
|        Type: | Classic                     Quorum |
|        Name: | *                                  |
|  Durability: | Durable          Transient         |
| Auto delete: | No          Yes                    |

和Exchange的参数一样的意思



### Exchange绑定Queue

| 											 | 		  |
| :--------------------- | ---- |
| To queue             : | *    |
| Routing key:           |      |

写入Queue地址，填写Routing Key



## 命令行工具

1.没有端口的情况下

2.需要统一配置的时候

- List         查看
- purge     清空
- Delete    删除

### Status

- Connection   `rabbitmqctl list_connection`
- Consumers   `rabbitmqctl list_consumers`
- Exchange   `rabbitmqctl list_exchanges`  

### User

- 新建  `rabbitmqctl add_user`
- 修改密码  `rabbitmqctl change_password`
- 删除用户  `rabbitmqctl delete_user`
- 查看  `rabbitmqctl list_users`
- 角色  `rabbitmqctl set_user_tags`

### App

- 启动  `rabbitmqctl start_app`
- 关闭(除Erlang)  `rabbitmqctl stop_app`
- 关闭  `rabbitmqctl stop`

