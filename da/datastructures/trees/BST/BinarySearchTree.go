
https://github.com/Workiva/go-datastructures/blob/master/tree/avl/avl.go
// Name: Binary search tree
package BinarySearchTree

import (
	"errors"
)

const (
	CompareEqual = 1 << iota
	CompareLess
	CompareMore
)

// ItemIterator allows callers of Ascend* to iterate in-order over portions of
// the tree.  When this function returns false, iteration will stop and the
// associated Ascend* function will immediately return.
type ItemIterator func(i string) bool

// stringCompare allows callers specify the compare function for items searching.
// Return the compare state
type ItemCompare func(i string) int

type Node struct {
	left   *Node
	right  *Node
	parent *Node
	data   string
}

type BinarySearchTree struct {
	Root *Node
}

func (this *BinarySearchTree) Search(it string) {

}

func (this *BinarySearchTree) Insert(it string) {

}

func (this *BinarySearchTree) Remove(it string) {

}
