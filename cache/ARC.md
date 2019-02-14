# golang-lru包学习 - ARC算法

## ARC算法
代码:[ARC](https://github.com/hashicorp/golang-lru/blob/master/arc.go)  
Wiki:[Adaptive replacement cache](https://en.wikipedia.org/wiki/Adaptive_replacement_cache)
论文:[A low overhead high performance buffer replacement algorithm](http://www.vldb.org/conf/1994/P439.PDF)

### ARC算法的描述
>ARC算法是为了整个LFU和LRU两种替换算法的优点,实现他们之间的自适应的workload调整。算法把有效缓存队列划分为两个,T1和T2,其中T1是LRU队列用于保存最近使用的条目、T2是LFU队列用于保存最常使用的条目。
他们各自都包含了一个名为ghost list的淘汰队列,分别命名为B1、B2,但tracking操作(即外部添加、查找操作)只针对T1、T2进行.ARC算法会根据缓存命中情况自动调整T1、T2的大小,以保证整个缓存既定长度的恒定。当新元素较多的时候,T1长度会增长T2长度会缩小,而当旧元素命中较多时候
T2长度会增长T1长度会缩小。

理论ARC算法核心结构:

+ T1, 用于保存最近的条目
+ T2, 用于保存最常用的条目,至少被引用两次
+ B1, 用于保存从T1淘汰的条目
+ B2, 用于保存从T2淘汰的条目

```
. . . [   B1  <-[     T1    <-!->      T2   ]->  B2   ] . .
      [ . . . . [ . . . . . . ! . .^. . . . ] . . . . ]
                [   fixed cache size (c)    ]
				[             L             ]
L可视为一个有效的缓存队列整体,其长度恒定不变。通过T1、T2长度的变化改变偏重P的大小。

				
首次添加  Add("A")
          | T1(LRU)  |      | B1(LRU)  |       | T2 (LFU)|     | B2 (LFU)|
		  |----------|      |----------|       |---------|     |---------|
ele(A)->  |  ele(A)  |      |          |       |         |     |         |
		  |----------|      |----------|       |---------|     |---------|
		
第二次添加或者访问 Add("A") / Get("A")
          | T1(LRU)  |      | B1(LRU)  |       | T2 (LFU)|     | B2 (LFU)|
		  |----------|      |----------|       |---------|     |---------|
          |          |      |          |       |  ele(A) |     |         |
		  |----------|      |----------|       |---------|     |---------|
		
对于已被T1或者T2淘汰的数据,再次添加 Add("A"),会直接进入T2
          | T1(LRU)  |      | B1(LRU)  |       | T2 (LFU)|     | B2 (LFU)|
		  |----------|      |----------|       |---------|     |---------|
          |          |      |  ele(A)  |       |         |     |         |
		  |----------|      |----------|       |---------|     |---------|
		
          | T1(LRU)  |      | B1(LRU)  |       | T2 (LFU)|     | B2 (LFU)|
		  |----------|      |----------|       |---------|     |---------|
          |          |      |          |       |  ele(A) |     |         |
		  |----------|      |----------|       |---------|     |---------|		
```



### 理论ARC算法替换过程

新元素:
+ 若空间不足,淘汰T2
+ 添加新元素到T1

已存在元素:
+ 若在B1或B2存在,移动到T2
+ 若空间不足,淘汰T1

查询命中
+ 若T1查询命中,移动到T2


## 算法实现:
本包跟2Q算法实现类似,将T2用LRU队列来实现。其结构如下图所示。T1,T2,B1,B2依然按照理论ARC中的设计,区别只是T2 B2用的是LRU算法而不是LFU。

本包算法实现结构:
```
          | T1(LRU)  |      | B1(LRU)  |       | T2 (LRU)|     | B2 (LRU)|
		  |----------|      |----------|       |---------|     |---------|
ele(A)->  |  ele(A)  |      |          |       |         |     |         |
		  |----------|      |----------|       |---------|     |---------|
```


ARC算法结构对象:
```
type ARCCache struct {
	size int // 整体缓存的既定长度
	p    int // P T1和T2的侧重值

	t1 simplelru.LRUCache // T1 最近使用队列
	b1 simplelru.LRUCache // B1 T1淘汰队列

	t2 simplelru.LRUCache // T2 最常使用队列
	b2 simplelru.LRUCache // B2 T2淘汰队列

	lock sync.RWMutex  // 读写锁
}
```



### 函数

```
type ARCCache interface {
  NewARC(size int) (*TwoQueueCache, error) 				// 创建一个ARC算法的缓存对象
  Add(key, value interface{}) 					    	// 添加元素,若元素不存在则添加元素，若元素存在则更新元素
  Get(key interface{}) (value interface{}, ok bool) 	// 获取元素
  Contains(key interface{}) (ok bool) 					// 查看该元素是否存在,但不会更新元素位置
  Peek(key interface{}) (value interface{}, ok bool)	// 获取元素,但不会更新元素位置
  Remove(key interface{}) bool							// 删除元素
  Len() int												// 返回队列中元素数量
  Purge()												// 清除队列中全部元素
}
```


### 过程

ARC算法核心过程是Add(),在Add()方法中需要针对不同情况处理四个队列的淘汰问题和T1、T2长度问题,最后更新偏重值P。

Add() 添加元素
```
func (c *ARCCache) Add(key, value interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.t1.Contains(key) {		// 若在T1存在,移动到T2
		c.t1.Remove(key)
		c.t2.Add(key, value)
		return
	}

	if c.t2.Contains(key) {		// 若在T2存在,更新值
		c.t2.Add(key, value)
		return
	}

	if c.b1.Contains(key) {		// 若在B1存在,移动到T2
		delta := 1				// 侧重增量
		b1Len := c.b1.Len()
		b2Len := c.b2.Len()
		if b2Len > b1Len {
			delta = b2Len / b1Len	
		}
		if c.p+delta >= c.size { 	// 调整侧重,靠近T1这边
			c.p = c.size
		} else {
			c.p += delta
		}

		if c.t1.Len()+c.t2.Len() >= c.size {
			c.replace(false)		// 淘汰T2,缓存窗口往T1移动
		}

		c.b1.Remove(key)
		
		c.t2.Add(key, value)
		return
	}

	if c.b2.Contains(key) {		// 若在B2存在,移动到T2
		delta := 1				// 侧重增量
		b1Len := c.b1.Len()
		b2Len := c.b2.Len()
		if b1Len > b2Len {
			delta = b1Len / b2Len
		}
		if delta >= c.p {		// 调整侧重,靠近T2这边
			c.p = 0
		} else {
			c.p -= delta
		}

		if c.t1.Len()+c.t2.Len() >= c.size {
			c.replace(true)		// 淘汰T1,缓存窗口往T2移动
		}

		c.b2.Remove(key)

		c.t2.Add(key, value)
		return
	}

	// T1 T2 B1 B2 都无该元素,添加到T1
	if c.t1.Len()+c.t2.Len() >= c.size {	// 超出整体缓存既定大小
		c.replace(false)					// T2淘汰
	}

	if c.b1.Len() > c.size-c.p {
		c.b1.RemoveOldest()
	}
	if c.b2.Len() > c.p {
		c.b2.RemoveOldest()
	}

	c.t1.Add(key, value)
	return
}
```

Get方法与2Q相似,T1命中的则移动到T2队列  
Get() 获取元素
```
func (c *ARCCache) Get(key interface{}) (value interface{}, ok bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if val, ok := c.t1.Peek(key); ok {		// 若在T1存在,移动到T2
		c.t1.Remove(key)
		c.t2.Add(key, val)
		return val, ok
	}

	if val, ok := c.t2.Get(key); ok {
		return val, ok
	}

	// No hit
	return nil, false
}
```


## 总结
+ ARC算法最大优势是自动调整两种队列的负载,也无需像2Q算法那样需要设置参数
+ ARC算法可减弱遍历访问造成的缓存污染