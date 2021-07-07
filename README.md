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


## 视频表

```sql

create table video
(
    id         int not null
        primary key comment '视频的ID',
    title      varchar(255) comment '视频标题',
    sub_title  varchar(255) comment '视频子标题',
    status     tinyint comment '-1下线  1上线',
    add_time   int comment '创建时间',
    imgh        varchar(255) comment '视频图片-竖屏',
    imgv        varchar(255) comment '视频图片-横屏',
    sort       varchar(255) comment '排序',
    channel_id int comment '视频归属channel的ID',
    type_id  int comment '频道类型ID',
    region_id  int comment '地区类型ID',
    user_id int comment '用户ID',
    episodes_count int comment '视频集数',
    episodes_update int comment '更新时间',
    is_end tinyint comment '是否完结 0 否  1是',
    is_hot tinyint comment '是否热播 0否  1是',
    is_recommend tinyint comment '是否推荐 0否  1是',
    comment int comment '评论数'

) DEFAULT CHARSET utf8 
  COLLATE utf8_general_ci  comment '视频数据表';


```

## 视频频道地区表

```sql

create table channel_region
(
    id         int not null
        primary key comment '频道地区的ID',
    name       varchar(50) comment '频道地区',
    channel_id int comment '频道ID',
    add_time   int comment '创建时间',
    status     tinyint comment '0下线 1上线',
    sort       int comment '排序'
) DEFAULT CHARSET utf8
  COLLATE utf8_general_ci comment '视频频道地区表';

```

## 频道类型表

```sql

create table channel_type
(
    id         int not null
        primary key comment '频道地区的ID',
    name       varchar(50) comment '频道地区',
    channel_id int comment '频道ID',
    add_time   int comment '创建时间',
    status     tinyint comment '0下线 1上线',
    sort       int comment '排序'
) DEFAULT CHARSET utf8
  COLLATE utf8_general_ci comment '频道类型表';

```

## 视频频道

```sql

create table channel
(
    id int not null
        primary key comment '频道的ID',
    name varchar(50) comment '频道的名字',
    add_time   int comment '创建时间',
    status     tinyint comment '0下线 1上线'
) DEFAULT CHARSET utf8
  COLLATE utf8_general_ci comment '视频频道';


```

## 广告栏数据表

```sql

create table advert
(
    id         int not null
        primary key comment '广告的ID',
    title      varchar(255) comment '广告标题',
    sub_title  varchar(255) comment '广告的子标题',
    channel_id int comment '当前广告归属channel的ID',
    img        varchar(255) comment '广告图片',
    sort       varchar(255) comment '排序',
    add_time   int comment '创建时间',
    url        varchar(255) comment '跳转地址',
    status     tinyint comment '0下线 1上线'
) DEFAULT CHARSET utf8
  COLLATE utf8_general_ci comment '广告栏数据表';

```