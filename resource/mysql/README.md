# Mysql

## 安装

1.从dockerhub下载5.7.34版本的Mysql

```shell
docker pull mysql:5.7.34

5.7.34: Pulling from library/mysql
Digest: sha256:1a2f9cd257e75cc80e9118b303d1648366bc2049101449bf2c8d82b022ea86b7
Status: Image is up to date for mysql:5.7.34
docker.io/library/mysql:5.7.34
```

2.启动

```shell
docker run -it --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=wzmmysql mysql:5.7.34
```

- -p 3306:3306 指定了映射的Port
- MYSQL_ROOT_PASSWORD=wzmmysql    指定root密码

PS：如果需要指定`config`或者其他路径可以通过`-v localPath:dockerPath` 



## 创建

### 建立数据库

```shell
# 建立database,并使用utf8
create database iyouku character set utf8

1 row affected in 13 ms

# 切换到iyouku库
use iyouku

completed in 9 ms
```

### 建立表

