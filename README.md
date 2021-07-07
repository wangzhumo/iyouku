## 用户表


```sql

create table ucenter
(
    uid      bigint           not null
        primary key comment '用户ID',
    nick     varchar(60)      not null comment '用户昵称',
    password varchar(36)      not null comment '用户密码',
    status   int    default 0 not null comment '用户状态',
    `create` bigint default 0 not null comment '创建日期',
    mobile   varchar(20)      not null comment '用户手机号',
    avatar   varchar(64)      not null comment '用户头像',
    name     varchar(20)      not null comment '用户姓名'
)
    comment '用户表'

```