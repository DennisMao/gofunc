// Name: Binary tree
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

func New(i string) *BinarySearchTree {

	return &BinarySearchTree{
		Root: &Node{nil, nil, nil, i},
	}

}

func (this *BinarySearchTree) Insert(item string) error {
	return findInsertNode(this.Root, item)
}

func findInsertNode(subNode *Node, item string) error {

	if subNode.data == item {
		return errors.New("item exists")
	}

	if subNode.data > item {
		if subNode.left == nil {
			subNode.left = &Node{
				parent: subNode,
				data:   item,
			}
			return nil
		}

		findInsertNode(subNode.left, item)
	} else {
		if subNode.right == nil {
			subNode.right = &Node{
				parent: subNode,
				data:   item,
			}
			return nil
		}
		findInsertNode(subNode.right, item)
	}

	return nil
}

func (tree *BinarySearchTree) Searchstring(f ItemCompare) (*Node, bool) {
	if tree.Root == nil {
		return nil, false
	}
	currentNode := tree.Root
	for currentNode != nil {

		switch f(currentNode.data) {
		case CompareEqual:
			return currentNode, true
		case CompareLess:
			currentNode = currentNode.right
		case CompareMore:
			currentNode = currentNode.left
		}
	}
	return nil, false
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
