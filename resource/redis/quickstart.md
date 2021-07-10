## Docker Redis

### 1.下载

>WebSite: https://redis.io/
>DockerHub: https://hub.docker.com/_/redis/ 

```shell
docker pull redis:6.2.4

C:\Workspace\GolandProjects\iyouku>docker images
REPOSITORY     TAG       IMAGE ID       CREATED         SIZE
redis          6.2.4     08502081bff6   2 weeks ago     105MB
ffdfgdfg/npc   latest    cad4d73b481e   3 months ago    12MB
mysql          5.7.30    9cfcce23593a   13 months ago   448MB
```
现在Redis的镜像已经下载完毕了，下一步启动

### 2.配置

配置文件：

`resource/redis/redis.conf`

可以去Github拉：

`https://github.com/redis/redis`

> bind 127.0.0.1 #只接受本地访问，注视掉使redis可以外部访问 
> 
> daemonize no #用守护线程的方式启动
> 
> requirepass #redis设置密码 
> 
> appendonly yes#redis持久化　　 
> 
> tcp-keepalive 300 #防止出现远程主机强迫关闭了一个现有的连接的错误 默认是300

### 3.映射文件
创建本地与docker映射的目录

```shell

# redis目录
mkdir ~/Docker/redis  
# redis的配置目录
mkdir ~/Docker/redis/conf
# redis的数据
mkdir ~/Docker/redis/data

```
然后将上一步配置好的`redis.conf`放到`~/Docker/redis/conf`目录下


### 4.docker启动redis

```shell

C:\Users\wangzhumo>docker run -p 6100:6100 --name redis -v C:\Develop\Docker\redis\conf\redis.conf:/etc/redis/redis.conf -v C:\Develop\Docker\redis\data:/data -d redis redis-server /etc/redis/redis.conf --appendonly yes
Unable to find image 'redis:latest' locally

```

此时问题是，我指定下载的`6.2.4`，docker没有找到`latest`的版本
只需要启动指定版本的`redis`即可

```shell

C:\Users\wangzhumo>docker run -p 6100:6100 --name redis -v C:\Develop\Docker\redis\conf\redis.conf:/etc/redis/redis.conf -v C:\Develop\Docker\redis\data:/data -d redis:6.2.4 redis-server /etc/redis/redis.conf --appendonly yes
a2cf524a3c7d156c78126131d93f64c496889f347cab80f5b8aa23cd4e1aaded

```
- -p 6100:6100          指定端口
- --name redis          指定启动后的名字
- -v {localPath}:{dockerPath}  建立映射关系
- -d redis:6.2.4        指定启动的镜像(这里指定了6.2.4版本)
- redis-server /etc/redis/redis.conf 这里指定了启动的配置文件
- –appendonly yes       redis启动后数据持久化

### 5.查看docker运行状态
```shell

C:\Users\wangzhumo>docker ps
CONTAINER ID   IMAGE         COMMAND                  CREATED          STATUS          PORTS                                       NAMES
a2cf524a3c7d   redis:6.2.4   "docker-entrypoint.s…"   22 minutes ago   Up 22 minutes   0.0.0.0:6379->6379/tcp, :::6379->6379/tcp   redis

```

可以看到，我们的`redis`已经成功运行了

### 6.进入docker中的容器

```shell

docker exec -it redis bash

C:\Users\wangzhumo>docker exec -it redis bash
root@a2cf524a3c7d:/data# ls
appendonly.aof

```

#### 查看日志：

```shell

C:\Users\wangzhumo>docker logs redis
1:C 09 Jul 2021 16:55:33.182 # oO0OoO0OoO0Oo Redis is starting oO0OoO0OoO0Oo
1:C 09 Jul 2021 16:55:33.182 # Redis version=6.2.4, bits=64, commit=00000000, modified=0, pid=1, just started
1:C 09 Jul 2021 16:55:33.183 # Configuration loaded
1:M 09 Jul 2021 16:55:33.183 * monotonic clock: POSIX clock_gettime
1:M 09 Jul 2021 16:55:33.184 * Running mode=standalone, port=6379.
1:M 09 Jul 2021 16:55:33.185 # Server initialized
1:M 09 Jul 2021 16:55:33.185 # WARNING overcommit_memory is set to 0! Background save may fail under low memory condition. To fix this issue add 'vm.overcommit_memory = 1' to /etc/sysctl.conf and then reboot or run the command 'sysctl vm.overcommit_memory=1' for this to take effect.
1:M 09 Jul 2021 16:55:33.185 * Ready to accept connections

```
发现了一个问题，提示我们没有设置`overcommit_memory`,那按照提示去修改一下

- 0 表示内核将检查是否有足够的可用内存供应用进程使用
  - 如果有足够的可用内存，内存申请允许
  - 否则，内存申请失败，并把错误返回给应用进程
  
- 1 表示内核允许分配所有的物理内存，而不管当前的内存状态如何。
  
- 2 表示内核允许分配超过所有物理内存和交换空间总和的内存

修改：
```shell

C:\Users\wangzhumo>docker exec -it redis bash
root@a2cf524a3c7d:/data# cat /proc/sys/vm/overcommit_memory
0
root@a2cf524a3c7d:/data# echo "vm.overcommit_memory=1" >> /etc/sysctl.conf
exit
C:\Users\wangzhumo>docker stop redis
redis
C:\Users\wangzhumo>docker start redis
redis




```

发现docker中好像无法处理,后续在linux上尝试
> https://github.com/bkuhl/redis-overcommit-on-host

### 7.连接Redis

#### redis-cli

```shell
╭─wangzhumo at Wangzhumo in /usr/local/etc using 21-07-10 - 15:20:41
╰─○ redis-cli -h 127.0.0.1 -p 6379
127.0.0.1:6379>

```
进入实例就可以进行操作了

#### 客户端：
> https://github.com/qishibo/AnotherRedisDesktopManager

## End