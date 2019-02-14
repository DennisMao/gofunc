# golang-lru包学习 - simple-lru

## simple-lru
代码:[simple-lru](https://github.com/hashicorp/golang-lru/blob/master/simplelru/lru.go)  

LRU算法核心是`把最近使用的条目放到靠近缓存的顶部位置,最少使用的条目自然会被移动到末端被淘汰`,其淘汰规则是基于访问时间而不是LFU的基于访问次数,即添加、更新、获取操作都会使得目标条目放到队列的前端(front)。    
simple-lru是lru算法一个简单实现,但是整个包的基础数据结构,接下来的2q和arc算法都是以当前的数据结构进行展开实现。

```
// LRU implements a non-thread safe fixed size LRU cache
type LRU struct {
	size      int							// 队列既定长度
	evictList *list.List					// 元素队列
	items     map[interface{}]*list.Element // 元素哈希索引
	onEvict   EvictCallback					// 淘汰操作的回调函数
}
```

由于链表的添加、删除都会涉及遍历查找操作,即便`list`用的是官方的`container/list`双向链表,但是操作频繁的时候其
O(n)的时间复杂度还是会对性能有影响。simple-lru的底层利用了map来给元素加索引,实现`key -> element(key)`结构,
通过map结构能直接取到元素地址,时间复杂度将为O(1).

>优化: 如果缓存队列长度是既定的,那索引map容量是与队列长度一致,因此底层的map在初始化的时候也建议直接指定cap大小。
>减少底层map的rehash操作。


### 函数

```
type LRUCache interface {
  Add(key, value interface{}) bool 						// 添加元素,若元素不存在则添加元素,当该操作导致其他元素淘汰时候返回true.若元素存在则更新元素
  Get(key interface{}) (value interface{}, ok bool) 	// 获取元素,如果成功找到该元素会自动将当前元素放到队列前端(front)
  Contains(key interface{}) (ok bool) 					// 查看该元素是否存在,但不会更新元素位置
  Peek(key interface{}) (value interface{}, ok bool)	// 获取元素,但不会更新元素位置
  Remove(key interface{}) bool							// 删除元素
  RemoveOldest() (interface{}, interface{}, bool)		// 删除最不常用元素(后端rear元素)
  GetOldest() (interface{}, interface{}, bool)			// 获取最不常用元素(后端rear元素),但不更新元素位置
  Keys() []interface{}									// 以数组形式返回全部key,由最不常用(后端rear)到最常用(前端front)
  Len() int												// 返回队列中元素数量
  Purge()												// 清除队列中全部元素
}
```


### 过程

LRU算法核心过程是Get和Add,我们针对这两个方法进到代码里看看

Add() 添加元素
```
func (c *LRU) Add(key, value interface{}) (evicted bool) {
	// 在map中查找该元素是否存在
	if ent, ok := c.items[key]; ok {  		
		c.evictList.MoveToFront(ent)		// 存在,在队列中把该元素移到前端
		ent.Value.(*entry).value = value	// 更新值
		return false
	}

	ent := &entry{key, value}				// 不存在,创建新元素
	entry := c.evictList.PushFront(ent)		// 添加元素到队列的前端
	c.items[key] = entry					// 添加元素到map中

	evict := c.evictList.Len() > c.size		// 判断当前队列是否超出既定长度
	if evict {
		c.removeOldest()					// 超出,淘汰最不常用(后端rear)元素
	}
	return evict
}
```

Get() 获取元素
```
func (c *LRU) Get(key interface{}) (value interface{}, ok bool) {
	// 在map中查找该元素是否存在	
	if ent, ok := c.items[key]; ok {	
		c.evictList.MoveToFront(ent)			// 存在,在队列中把该元素移到前端
		return ent.Value.(*entry).value, true   // 返回元素
	}
	return
}
```


## 总结
+ simple-lru基础`container/list`包,底层为一个双向链表和map索引
+ simple-lru数据结构并非线程安全,需要外部增加互斥锁或读写锁