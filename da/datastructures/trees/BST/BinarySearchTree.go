// Name: Binary tree
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
	left, right, parent *Node
	data                string
}

type BinarySearchTree struct {
	Root *Node
}

func New(i string) *BinarySearchTree {

	return &BinarySearchTree{
		Root: &Node{nil, nil, nil, i},
	}

}

func (tree *BinarySearchTree) Insert(it string) *Node {

	insertPos := searchWithHot(tree.Root, it) // find a suitable insert postition
	if insertPos == nil {
		return nil
	}

	if insertPos.data == it {
		return insertPos // 'it' exists
	}

	log.Printf("insertPos:%s Target:%s\n", insertPos.data, it)

	newNode := Node{parent: insertPos, data: it} // create a new node for 'it'

	// connect the new node with it's parents
	if insertPos.data > it {
		insertPos.left = &newNode
	} else {
		insertPos.right = &newNode
	}

	return &newNode
}

func (tree *BinarySearchTree) Search(it string) *Node {
	if tree.Root == nil {
		return nil
	}

	return search(tree.Root, it)
}

// searchWithHot can find the data 'it' in a tree using recursive algorithm,
// and return a pointer pointed to 'it' if find while return a pointer
// poined to 'it' 's father node if not find.
func searchWithHot(sub *Node, it string) *Node {
	if sub == nil {
		return nil
	}

	if sub.data == it {
		return sub
	}

	log.Printf("当前search:%s  目标:%s \n", sub.data, it)

	if sub.data > it {
		if sub.left != nil {
			return searchWithHot(sub.left, it)
		} else {
			return sub
		}
	} else {
		if sub.right != nil {
			return searchWithHot(sub.right, it)
		} else {
			return sub
		}
	}

	return nil
}

// search can find the data 'it' in a tree using recursive algorithm,
// and return a pointer pointed to 'it' if find while return 'nil' if not find.
func search(sub *Node, it string) *Node {
	if sub == nil {
		return nil
	}

	if sub.data == it {
		return sub
	}

	if sub.data < it {
		return search(sub.left, it)
	} else {
		return search(sub.right, it)
	}

	return nil
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