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

func DeleteNode(node *Element) error {
	if node == nil {
		return fmt.Errorf("node is nil")
	}

	// Connect previous node to the next.
	if node.prev != nil {
		level := len(node.prev.next)
		for i := 0; i < level; i++ {
			node.prev.next[i] = node.next[i]
		}
	}

	node.Reset()
	return nil
}
