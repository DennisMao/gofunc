package skiplist

import (
	"fmt"
)

type Element struct {
	key, Value interface{}
	score      float64 //当前版本  score 直接等于key,后期加入转换函数
	prev       *Element
	next       []ElementLevel
}

type ElementLevel struct {
	next *Element
	span int
}

func (e *Element) Reset() {
	e.key = nil
	e.Value = nil
	e.prev = nil
	e.next = nil
	e.score = 0.0
}

func CreateNode(level int, score float64, k, v interface{}) *Element {
	newNode := Element{
		key:   k,
		Value: v,
		score: score,
		next:  make([]ElementLevel, level),
	}
	return &newNode
}

<<<<<<< HEAD
=======
// DeleteNode 会重置给定的跳表节点对象。对象循环系统会在后续加入。对象循环系统可以收集待删除对象到对象池而不是直接通过
// GC释放掉其内存。此举可以有效减少GC压力。
//
// DeleteNode can reset a specified element of the skiplisk.Object recycle system is moving on.
// It can collect the node object we are going to delete with an object pool instead of freeing by
// garbage collection.This system can effectually reduce the stress of GC(garbage collection).
>>>>>>> f6da005d698eb739f093301078e91f1699b2101a
func DeleteNode(node *Element) error {
	if node == nil {
		return fmt.Errorf("node is nil")
	}

<<<<<<< HEAD
	// Connect previous node to the next.
	if node.prev != nil {
		level := len(node.prev.next)
		for i := 0; i < level; i++ {
			node.prev.next[i] = node.next[i]
		}
	}

=======
>>>>>>> f6da005d698eb739f093301078e91f1699b2101a
	node.Reset()
	return nil
}
