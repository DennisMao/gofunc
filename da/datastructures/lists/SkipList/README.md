
# 跳跃表 skiplist


## 基本类型

### SkipList 跳跃表

type SkipList struct{
    front,rear *Element // 首末节点(rear非论文标准)
    level int           // 当前等级
    used int            // 存在元素统计
    length int          // 长度

}


###  Elemenet 节点

type Element struct {
    key,value interface{}   // k,v存储元素(key非论文标准)
    score float64           // 分值
    next []ElementLevel     // 后继节点
    prev *Element           // 前驱节点(非论文标准)
}

type ElementLevel struct{
	span int	// 该层跨越的节点数量(非论文标准)
	next *Element
}


## ADT 抽象接口

### SkipList接口
|函数名 |作用|复杂度|
|:-|:-|:-|
|Create|新建并初始化一个跳跃表|O(L)|
|Free|释放给定的跳跃表|O(N)|
|RandomLevel|得到新节点的层数(抛硬币法的改进)|O(1)|
|Insert|将给定的score与member新建节点并添加到跳表中|O(logN)|
|Delete|删除给定的score与member在跳表中匹配的节点|O(logN)|
|IsInRange|检查跳表中的元素score值是否在给定的范围内|O(1)|
|FirstInRange|查找第一个符合给定范围的节点|O(logN)|
|LastInRange|查找最后一个符合给定范围的节点|O(logN)|
|DeleteRangeByScore|删除score值在给定范围内的节点|O(logN)+O(M)|
|DeleteRangeByRank|删除排名在给定范围内的节点|O(logN)+O(M)|
|GetRank|返回给定score与member在集合中的排名|O(logN)|
|GetElementByRank|根据给定的rank来查找元素|O(logN)|

### Node接口
|函数名 |作用|复杂度|
|:-|:-|:-|
|CreateNode|新建并返回一个跳表节点|O(1)|
|FreeNode|释放给定的节点|O(1)|
|DeleteNode|删除给定的跳表节点|O(L)|

## 引用
+ [跳跃表wiki](https://zh.wikipedia.org/wiki/%E8%B7%B3%E8%B7%83%E5%88%97%E8%A1%A8)