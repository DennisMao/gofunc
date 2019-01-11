//https://github.com/Workiva/go-datastructures/blob/master/tree/avl/avl.go
// Name: Binary search tree
package BinarySearchTree

import (
	"log"
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

	_high int // high

	data string
}

type BinarySearchTree struct {
	Root *Node
}

func New(it string) *BinarySearchTree {
	root := new(Node)
	root._high = 0
	root.data = it

	return &BinarySearchTree{root}
}

func (this *BinarySearchTree) Search(it string) *Node {

	_hot := searchIn(it, this.Root)
	if _hot.data != it {
		return nil
	} else {
		return _hot
	}
}

func searchIn(it string, _hot *Node) *Node {
	if _hot == nil {
		return nil
	}

	if _hot.data == it {
		return _hot
	}

	if it < _hot.data {
		if _hot.left != nil {
			return searchIn(it, _hot.left)
		}
	} else {
		if _hot.right != nil {
			return searchIn(it, _hot.right)
		}
	}

	return _hot
}

func updateAboveHigh(_hot *Node) {

}

func (this *BinarySearchTree) Insert(it string) *Node {
	_hot := searchIn(it, this.Root)
	if _hot.data == it {
		return _hot
	}

	newNode := new(Node)
	newNode.parent = _hot
	newNode.data = it
	newNode._high = _hot._high + 1

	if _hot.left == nil {
		_hot.left = newNode
	} else {
		_hot.right = newNode
	}

	return newNode

}

func (this *BinarySearchTree) Remove(it string) {

}

// 镜像翻转
func (tree *BinarySearchTree) Mirror(subNode *Node) {
	if subNode == nil {
		return
	}

	subNode.left, subNode.right = subNode.right, subNode.left
	tree.Mirror(subNode.left)
	tree.Mirror(subNode.right)
}

// 前序
//（1）访问根结点
//（2）前序遍历左子树
//（3）前序遍历右子树
func (tree *BinarySearchTree) PreorderTraversal(subNode *Node, it ItemIterator) {
	if subNode == nil {
		return
	}

	it(subNode.data)
	tree.PreorderTraversal(subNode.left, it)
	tree.PreorderTraversal(subNode.right, it)

}

// 中序
//（1）中序遍历左子树
//（2）访问根结点
//（3）中序遍历右子树
func (tree *BinarySearchTree) InorderTraversal(subNode *Node, it ItemIterator) {
	if subNode == nil {
		return
	}

	tree.InorderTraversal(subNode.left, it)
	it(subNode.data)
	tree.InorderTraversal(subNode.right, it)
}

// 后序
//（1）后序遍历左子树
//（2）后序遍历右子树
//（3）访问根结点
func (tree *BinarySearchTree) PostorderTraversal(subNode *Node, it ItemIterator) {
	if subNode == nil {
		return
	}

	if subNode.left != nil {
		tree.PostorderTraversal(subNode.left, it)
	}
	if subNode.right != nil {
		tree.PostorderTraversal(subNode.right, it)
	}
	it(subNode.data)

}

// 层级遍历
// (1)从根部开始入栈
// (2)出栈,输出
// (3)左孩子入栈、右孩子入栈
func (tree *BinarySearchTree) LevelTraversal(subNode *Node, it ItemIterator) {
	if subNode == nil {
		return
	}

	stack := make(chan *Node, 128)
	stack <- subNode
	for {
		if len(stack) == 0 {
			break
		}

		curNode := <-stack
		if curNode == nil {
			break
		}

		it(curNode.data) //输出

		if curNode.left != nil {
			stack <- curNode.left
		}
		if curNode.right != nil {
			stack <- curNode.right
		}
	}

	close(stack)
}
