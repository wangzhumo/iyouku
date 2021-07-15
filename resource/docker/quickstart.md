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