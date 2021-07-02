## 用户表


```sql


Data Source: qnap Schema: iyouku Table: ucenter  -- 用户表
-- auto-generated definition
create table ucenter
(
uid      bigint           not null
primary key,
nick     varchar(60)      not null,
password varchar(36)      not null,
status   int    default 0 not null,
`create` bigint default 0 not null,
mobile   varchar(20)      not null,
avatar   varchar(64)      not null,
name     varchar(20)      not null
)
comment '用户表';

```