// 线性表-跳跃表
package skiplist

import (
	"math/rand"
	"time"
)

const (
	DEFAULT_SIZE = 1024
	MAX_LEVEL    = 16
)

// TODO
// 1.改抽象接口,若Go2有泛型直接用泛型替代
// 2.提取后单独写处理函数
var DefaultCompareKey = func(key1, key2 interface{}) int {
	ret := -1
	k1 := key1.(int64)
	k2 := key2.(int64)
	if k1 > k2 {
		return ret
	}

	if k1 == k2 {
		return 0
	} else {
		return 1
	}
}

type SkipList struct {
	front, rear *Element
	compareKey  func(key1, key2 interface{}) int

	random *rand.Rand
	level  int
	length int
	size   int
}

func Create(size int) *SkipList {
	if size == 0 {
		size = DEFAULT_SIZE
	}

	newList := &SkipList{CreateNode(MAX_LEVEL, 0, nil, nil),
		nil,
		DefaultCompareKey,
		rand.New(rand.NewSource(time.Now().Unix())),
		0,
		0,
		size,
	}

	for i, _ := range newList.front.next {
		newList.front.next[i].next = nil
		newList.front.next[i].span = 0
	}

	return newList
}

func (this *SkipList) Size() int {
	return this.size
}

func (this *SkipList) Len() int {
	return this.length
}

func (this *SkipList) Insert(score float64, k, v interface{}) error {
	updateNodes := make([]*Element, MAX_LEVEL) // 记录需要更新的节点
	rank := make([]int, MAX_LEVEL)             // 记录span查找的跨度
	x := new(Element)

	// 从头节点front开始搜索,按level从高到低逐层查找直到最后一层,把需要更新的节点放入到updateNodes数组中.
	x = this.front
	for i := this.level; i >= 0; i-- {
		rank[i] = 0
		if i != this.level-1 {
			rank[i] = rank[i+1]
		}

		// 当当前遍历节点分数小于目标分数或者当前遍历节点key小于目标key值,继续遍历
		for x.next[i].next != nil && (x.next[i].next.score < score && this.compareKey(k, x.next[i].next.key) < 0) {
			rank[i] += x.next[i].span
			x = x.next[i].next
		}

		updateNodes[i] = x
	}

	// 添加新节点
	// 由于skiplist的Insert操作默认认为当前添加的节点是不存在的,所以允许重复的
	// 节点加入,但在redis实际应用是禁止重复的key加入,因此在skiplist的外部需要
	// 用hashtable来对key进行排重。

	// 随机新节点的层数(抛硬币法每个等级 平均概率)
	level := this.RandomLevel()
	// 若新插入节点level大于当前跳表level -> 升级
	for level > this.level {
		// 升级
		for i := this.level; i < level; i++ {
			rank[i] = 0
			// 升级的最顶层的前驱一定是根节点
			updateNodes[i] = this.front
			updateNodes[i].next[i].span = this.length

		}
		// 更新调表结构的level值
		this.level = level
	}

	// 插入新节点到对应的层
	x = CreateNode(level, score, k, v)
	for i := 0; i < level; i++ {
		// 链表插入节点
		x.next[i].next = updateNodes[i].next[i].next
		updateNodes[i].next[i].next = x

	}

	// 若新插入节点level未大于当前跳表level,高于新节点Level的span需要自增,因为跳表长度增加
	for i := level; i < this.level; i++ {
		updateNodes[i].next[i].span++
	}

	// 更新新节点的后继指针
	if x.next[0].next == nil {
		this.rear = x
	} else {
		x.next[0].next.prev = x
	}

	// 更新新节点的前驱指针
	if updateNodes[0] == this.front {
		x.prev = nil
	} else {
		x.prev = updateNodes[0]
	}

	this.length++
	return nil
}

func (this *SkipList) Delete(score float64, k interface{}) error { return nil }

// RandomLevel 函数会返回一个随机的等级,用于节点创建使用。其返回值区间是 [1,MAX_LEVEL]
//
// RandomLevel returns a random level for the new node we are going to create.
// And the return value of this function is between 1 and MAX_LEVEL (both inclusive,[1,MAX_LEVEL]).
func (this *SkipList) RandomLevel() int {
	level := 1

	// SkipList 随机节点的算法
	// Redis 原始实现:
	//  for ( random() & 0xFFFF < 0.25 * 0xFFFF ) {level ++ ;}
	// 16384 ~= 65535 * 0.25
	for level < MAX_LEVEL && ((rand.Int63() & 0xFFFF) < int64(16384)) { // 保证生成出的level的概率平均
		level++
	}

	return level
}

// Find 函数会查找指定score和k.若查询到,返回指定的节点.若未查询到,返回nil.
// 查询支持score重复的情况,首先匹配分值score,再匹配key.
//
// Find returns a pointer to a element we are going to find by score and key.
// This function supports storing the key with same score's value.
func (this *SkipList) Find(score float64, k interface{}) *Element {

	// 从头节点开始搜索
	x := this.front
	curLevel := this.level - 1
	for ; curLevel >= 0; curLevel-- {
		// 当当前遍历节点分数小于目标分数或者当前遍历节点key小于目标key值,继续遍历
		for x.next[curLevel].next != nil && x.next[curLevel].next.score <= score {
			x = x.next[curLevel].next
		}
	}

	// 同score级别比较key的值
	// 支持score相同的存储 (与论文不同)
	for ; curLevel >= 0; curLevel-- {
		for x.next[curLevel].next != nil && x.next[curLevel].next.score == score && this.compareKey(x.next[curLevel].next.key, k) < 0 {
			x = x.next[curLevel].next
		}
	}
	if x == this.front {
		return nil
	}

	return x
}

// 遍历迭代每一层
func (this *SkipList) RangeLevel(level int) []*Element {
	if level > this.level {
		return nil
	}
	if level < 0 {
		return nil
	}

	ret := make([]*Element, 0)
	x := this.front
	for x.next[level-1].next != nil {
		ret = append(ret, x)
		x = x.next[level-1].next
	}

	return ret
}
