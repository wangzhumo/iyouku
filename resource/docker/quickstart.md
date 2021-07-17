# Docker

## 源

Docker for mac 

设置参数为如下所示，而后重启即可

```json

{
  "builder": {
    "gc": {
      "defaultKeepStorage": "20GB",
      "enabled": true
    }
  },
  "experimental": true,
  "debug":true,
  "registry-mirrors":[
     "http://hub-mirror.c.163.com",
     "https://docker.mirrors.ustc.edu.cn"
  ],
  "features": {
    "buildkit": true
  }
}

```

## 安装

### MSYQL

```shell
docker run -itd -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 --name mysql mysql:5.7.34
```
参数：
- -i：以交互模式运行，配合-t
- -t：为容器重新分配一个伪输入终端，配合-i
- -d：后台运行容器
- -p：端口映射，格式为主机端口:容器端口
- -e：设置环境变量，这里设置的是mysql的root密码
- --name：设置容器名


## 状态

### Images

#### 已经拉取到本地的Images
```shell
╭─wangzhumo at Wangzhumo in ~ using 21-07-17 - 17:28:51
╰─○ docker images                                                                           
REPOSITORY      TAG            IMAGE ID       CREATED       SIZE
rabbitmq        3-management   85e83aca5d60   3 days ago    249MB
elasticsearch   7.13.3         84840c8322fe   2 weeks ago   1.02GB
redis           6.2            08502081bff6   3 weeks ago   105MB
mysql           5.7.34         09361feeb475   3 weeks ago   447MB
```

#### Search

```shell
docker search {name}
```
#### 运行
```shell
docker ps

CONTAINER ID   IMAGE                   COMMAND                  CREATED         STATUS         PORTS                                                                                                                                                 NAMES
93729334f616   mysql:5.7.34            "docker-entrypoint.s…"   6 minutes ago   Up 6 minutes   0.0.0.0:3306->3306/tcp, :::3306->3306/tcp, 33060/tcp                                                                                                  mysql
1e6afe1467a1   rabbitmq:3-management   "docker-entrypoint.s…"   2 days ago      Up 31 hours    4369/tcp, 5671/tcp, 0.0.0.0:5672->5672/tcp, :::5672->5672/tcp, 15671/tcp, 15691-15692/tcp, 25672/tcp, 0.0.0.0:15672->15672/tcp, :::15672->15672/tcp   rabbitmq
cebcfd58c059   redis:6.2               "docker-entrypoint.s…"   5 days ago      Up 31 hours    0.0.0.0:6379->6379/tcp, :::6379->6379/tcp                                                                                                             redis
```
这里常用的就是 CONTAINER ID , Port 了


## 进入容器

```shell
docker exec -it 93729334f616 /bin/bash

╭─wangzhumo at Wangzhumo in ~ using 21-07-17 - 17:30:36
╰─○ docker exec -it 93729334f616 /bin/bash
root@93729334f616:/# 
```

