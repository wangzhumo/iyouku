



# Mysql索引

## 索引概述

索引是存储引擎用于快速找到记录的一种数据结构

> 一个常见而又简单的索 引例子是书籍的目录

### 优缺点

#### 索引的优点:

- 减少扫描的数据量，加速查询
- 减少或完全消除数据库的排序操作(ORDER BY)，因为索引是有序的
- 将服务器的随机 IO 变为顺序 IO，例如，想要查询 salary 处于 1500 ~ 2100 的员工，就可以按照索引顺序查 询

#### 索引的缺点:

- 索引会占据额外的存储空间(毕竟它是数据结构)，包括磁盘和内存 
- 由于对数据需要排序，自然会影响到数据更新(插入、更新、删除)的速度

### Mysql中索引原理

InnoDB 而言，它的内部实现使用的是 B+ 树。

#### B 树性质

- 根节点至少有两个子节点
- 每个节点包含 k - 1 个元素和 k 个子节点，其中 m/2 <= k <= m(元素是存储的数据) 
- 每个叶子节点都包含 k - 1 个元素，且位于同一层，其中 m/2 <= k <= m 
- 每个节点中的元素从小到大排列，类似于一个有序数组

#### B+ 树

B+ 树是在 B 树之上改进得到的，它又添加了两项约束

- 除叶子节点之外的其他节点都不保存数据，所以，数据在同一层

- 叶子节点之间按照排列顺序链接在一起，形成了一个有序链表



### 索引分类

- **普通索引**:针对于单个列创建的索引，之所以说它普通是因为它对列值没有什么限制，允许被索引的列包含重 复的值
- **唯一索引**:正如它的关键字一样，它要求列值是唯一的，这个索引保证了数据记录的唯一性 
- **主键索引**:它是一种特殊的唯一索引，在一张表中只能定义一个(但不是必须)主键索引 
- **联合索引**:也被称为复合索引，它是将多个列值绑定在一起作为索引
