# B-tree
## 库介绍
[btree](https://github.com/google/btree)本库是go语言的Btree的一个实现。Btree跟二叉树相比,它是多路搜索,支持范围搜索而且搜索效率稳定。

## 使用方法
测试代码实现了自定义一个数据结构,并用B树进行存储,然后对其进行一系列的操作与查找.代码地址:https://github.com/DennisMao/goexperience/google/btree


### 数据对象
```
// Item represents a single object in the tree.
type Item interface {
	// Less tests whether the current item is less than the given argument.
	//
	// This must provide a strict weak ordering.
	// If !a.Less(b) && !b.Less(a), we treat this to mean a == b (i.e. we can only
	// hold one of either a or b in the tree).
	Less(than Item) bool
}
``` 

我们实现一个数据单元,来应用这个接口

```
type Data struct {
	Id   int64  //数据编号
	Content string //数据内容
}

func (this *Data) Less(item btree.Item) bool {
	if v, ok := item.(*Data); ok {
		return this.Id < v.Id
	}

	fmt.Println("Error: not valid item")
	return false
}
```

### 初始化
```
//初始化一个二阶的B树
//并初始化一个独立的节点链表,默认长度32
tree := btree.New(2)  

//初始化一个二阶B树并指定节点链表
freeList := btree.NewFreeList(64)
tree := btree.NewWithFreeList(degree,freeList)
```


### 基本操作
-|-
Len|获取长度
Min|获取最小对象
Max|获取最大对象
Has|判断对象是否存在
Get|获取对象,若不存在返回nil
Clone|返回当前树的一个副本
Delete|删除指定对象
DeleteMax|删除最大对象
DeleteMin|删除最小对象
Clear|清空树
ReplaceOrInsert|替换或者插入,如果是已经存在的则返回原始对象

### 迭代器操作
-|-
Ascend|升序迭代获取整个树
AscendGreaterOrEqual|升序获取大于等于`[a,+∞)`
AscendLessThan|升序获取小于`(-∞,b)`
AscendRange|升序获取范围`[a,b)`
Descend|降序迭代获取整个树
DescendGreaterOrEqual|降序获取大于等于`[a,+∞)`
DescendLessThan|降序获取小于`(-∞,b)`
DescendRange|降序获取范围`[a,b)`

要使用迭代器,需要先实现`btree.ItemIterator`,核心是一个命名函数。在树进行范围查找时候符合条件的对象会调用该函数进行返回,因此我们需要在方法中实现对结果的收集。
```
// ItemIterator allows callers of Ascend* to iterate in-order over portions of
// the tree.  When this function returns false, iteration will stop and the
// associated Ascend* function will immediately return.
type ItemIterator func(i Item) bool
```

