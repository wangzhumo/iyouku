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

