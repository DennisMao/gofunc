# 数据结构
使用Go语言实现的常用数据结构用于练手,代码的实现结合了清华大学邓俊辉《数据结构课程》和《大话数据结构》的课设，部分设计思路参考了《Redis 2.4》和《Go源码 1.10》.
单元测试采用[testify](github.com/stretchr/testify/assert)框架,Dot图生成采用[memviz](github.com/bradleyjkemp/memviz).

使用说明:

+ 1.安装[graphviz](http://www.graphviz.org/),若无需生成Dot图可忽略此步
+ 2.cd 到指定数据结构目录
+ 3.go test 
+ 4.程序会自动进行单元测试并生成出dot文件 XXX.gv和可视图 XXX.png。

HashMap:
![HashMap](./hashmap/HashMap.png)



## 目录
### 向量
结构名称|代码|实现
:-|:-|:-
向量|[Vector](./vector)|√|

### 链表
结构名称|代码|实现
:-|:-|:-
双向链表|[LinkedList](./lists/LinkedList)|√|
环|[Ring](./lists/Ring)|√|
跳跃表|[SkipList](./lists/SkipList)||

### 串
结构名称|代码|实现
:-|:-|:-
串|[String](./strings)|√|

### 树
结构名称|代码|实现
:-|:-|:-
二叉树|[BinaryTree](./trees/BinaryTree)|√|
BST|[BinarySearchTree](./trees/BinarySearchTree)|√|
AVL|[AvlTree](./trees/AvlTree)||
B树|[Btree](./trees/Btree)||
红黑树|[RedBlackTree](./trees/RedBlackTree)||

### HashTable
结构名称|代码|实现
:-|:-|:-
哈希表|[HashMap](./hashmap)|√|

