# ElasticSearch

## 概述

### Lucene

  **倒排索引**  是Apache旗下的一款全文搜索引擎

> [终于有人把Elasticsearch原理讲透了！ - 知乎 (zhihu.com)](https://zhuanlan.zhihu.com/p/62892586)

### Elasticsearch

是一个**基于Lucene搜索引擎**为核心构建的开源，分布式，RESTful**搜索服务器**

> [Elasticsearch：官方分布式搜索和分析引擎 | Elastic](https://www.elastic.co/cn/elasticsearch/)

**特点：**

- 速度快(近乎实时)
- 灵活
- 稳定,简单
- 可以做集群    

**储存：**

集群 节点 分片 文档 副本

**分词：**

- 多种分词器支持
- 可以配置分词库（优化搜索结果）
- 可以配置屏蔽词库（屏蔽某些结果）

### 分词器  

> [medcl/elasticsearch-analysis-ik: The IK Analysis plugin integrates Lucene IK analyzer into elasticsearch, support customized dictionary. (github.com)](https://github.com/medcl/elasticsearch-analysis-ik)

The IK Analysis plugin integrates Lucene IK analyzer (http://code.google.com/p/ik-analyzer/) into elasticsearch, support customized dictionary.

Analyzer: `ik_smart` , `ik_max_word` , Tokenizer: `ik_smart` , `ik_max_word`

### 工作流

#### 输入

- 文档输入
- Analyzer  分词器
- Index Document 
- Index Directory

当文档输入之后，会进行倒叙索引，而这个过程中是需要分词器配合的，索引建立成功，进行存储

#### 输出（查询）

- Search
- Build Query  分词器
- Ececute

查询时其实在构造查询时，也会进行分词，构建为一个Query的结构，最后执行查询



## 安装

### 下载

```shell
# 下载镜像，这里指定需要下载7.13.3版本
╭─wangzhumo at Wangzhumo in ~ using 21-07-17 - 18:01:14
╰─○ docker pull elasticsearch:7.13.3
7.13.3: Pulling from library/elasticsearch
Digest: sha256:759533051d0d5c67f7b09c8655ea00af3a8f15f98e2df4ea76ff2b47967488a5
Status: Image is up to date for elasticsearch:7.13.3
docker.io/library/elasticsearch:7.13.3

# 查看镜像
╭─wangzhumo at Wangzhumo in ~ using 21-07-17 - 18:01:55
╰─○ docker images
REPOSITORY      TAG            IMAGE ID       CREATED       SIZE
rabbitmq        3-management   85e83aca5d60   3 days ago    249MB
elasticsearch   7.13.3         84840c8322fe   2 weeks ago   1.02GB
redis           6.2            08502081bff6   3 weeks ago   105MB
mysql           5.7.34         09361feeb475   3 weeks ago   447MB

# elasticsearch 已经成功下载了
```



### 启动

```shell
# 启动 elasticsearch

╭─wangzhumo at Wangzhumo in ~ using 21-07-17 - 17:59:01
╰─○ docker run -itd --name es -p 9300:9300 -p 9200:9200 -e "discovery.type=single-node" -e "ES_JAVA_OPTS=-Xms512m -Xmx512m" elasticsearch:7.13.3

# 输出启动后的ID
d9fd0e6986394a801ddc3daeada968421faeccdc89f9a61b8ec32ecca338f4fe
```

- -p 9300:9300 -p 9200:9200  指定端口
- -e "discovery.type=single-node"   指定单节点运行
- -e "ES_JAVA_OPTS=-Xms512m -Xmx512m"  指定JVM的内存设置

因为这里只是测试使用，所以指定了内存 默认是1G

```shell
CONTAINER ID   IMAGE                   COMMAND                  CREATED          STATUS          PORTS                                                                                                                                                 NAMES
d9fd0e698639   elasticsearch:7.13.3    "/bin/tini -- /usr/l…"   4 minutes ago    Up 4 minutes    0.0.0.0:9200->9200/tcp, :::9200->9200/tcp, 0.0.0.0:9300->9300/tcp, :::9300->9300/tcp                                                                  es
93729334f616   mysql:5.7.34            "docker-entrypoint.s…"   41 minutes ago   Up 41 minutes   0.0.0.0:3306->3306/tcp, :::3306->3306/tcp, 33060/tcp                                                                                                  mysql
1e6afe1467a1   rabbitmq:3-management   "docker-entrypoint.s…"   2 days ago       Up 32 hours     4369/tcp, 5671/tcp, 0.0.0.0:5672->5672/tcp, :::5672->5672/tcp, 15671/tcp, 15691-15692/tcp, 25672/tcp, 0.0.0.0:15672->15672/tcp, :::15672->15672/tcp   rabbitmq
cebcfd58c059   redis:6.2               "docker-entrypoint.s…"   5 days ago       Up 32 hours     0.0.0.0:6379->6379/tcp, :::6379->6379/tcp                                                                                                             redis
```

elasticsearch:7.13.3  已经启动



######`----------------------------------Waring-----------------------------------`

下次一定要注意

ElasticSearch version

Ik  version

需要一致,Ik最新的是7.13.2 ，所以我们需要降级

`---------------------------------------------------------------------------------------`

## IK安装

### 在线安装

#### 1.进入镜像

```shell
# 进入镜像
docker exec -it es /bin/bash

cd bin
ls -l 
-rwxr-xr-x 1 elasticsearch root     2896 Jul  2 12:03 elasticsearch
-rwxr-xr-x 1 elasticsearch root      501 Jul  2 12:07 elasticsearch-certgen
-rwxr-xr-x 1 elasticsearch root      493 Jul  2 12:07 elasticsearch-certutil
-rwxr-xr-x 1 elasticsearch root      996 Jul  2 12:03 elasticsearch-cli
-rwxr-xr-x 1 elasticsearch root      443 Jul  2 12:07 elasticsearch-croneval
-rwxr-xr-x 1 elasticsearch root     4859 Jul  2 12:11 elasticsearch-env
-rwxr-xr-x 1 elasticsearch root     1828 Jul  2 12:03 elasticsearch-env-from-file
-rwxr-xr-x 1 elasticsearch root      168 Jul  2 12:03 elasticsearch-geoip
-rwxr-xr-x 1 elasticsearch root      184 Jul  2 12:03 elasticsearch-keystore
-rwxr-xr-x 1 elasticsearch root      450 Jul  2 12:07 elasticsearch-migrate
-rwxr-xr-x 1 elasticsearch root      126 Jul  2 12:03 elasticsearch-node
-rwxr-xr-x 1 elasticsearch root      172 Jul  2 12:03 elasticsearch-plugin
-rwxr-xr-x 1 elasticsearch root      441 Jul  2 12:07 elasticsearch-saml-metadata
-rwxr-xr-x 1 elasticsearch root      439 Jul  2 12:07 elasticsearch-service-tokens
-rwxr-xr-x 1 elasticsearch root      448 Jul  2 12:07 elasticsearch-setup-passwords
-rwxr-xr-x 1 elasticsearch root      118 Jul  2 12:03 elasticsearch-shard
-rwxr-xr-x 1 elasticsearch root      483 Jul  2 12:07 elasticsearch-sql-cli
-rwxr-xr-x 1 elasticsearch root 21956553 Jul  2 12:07 elasticsearch-sql-cli-7.13.3.jar
-rwxr-xr-x 1 elasticsearch root      436 Jul  2 12:07 elasticsearch-syskeygen
-rwxr-xr-x 1 elasticsearch root      436 Jul  2 12:07 elasticsearch-users
-rwxr-xr-x 1 elasticsearch root      356 Jul  2 12:07 x-pack-env
-rwxr-xr-x 1 elasticsearch root      364 Jul  2 12:07 x-pack-security-env
-rwxr-xr-x 1 elasticsearch root      363 Jul  2 12:07 x-pack-watcher-env

```

#### 2.安装





