// Name: Binary tree
package AvlTree

import (
	BST "goexperience/da/datastructures/trees/BST"
)

const (
	CompareEqual = 1 << iota
	CompareLess
	CompareMore
)

type Node=BST.Node

type AvlTree struct {
	BST.BinarySearchTree
}

// 插入
// @Override
func (tree *AvlTree) Insert(item string) {
	if tree.Root == nil {
		tree.Root = &Node{}
	
	}

	insertPos := searchWithHot(tree.Root, it) // find a suitable insert postition
	if insertPos == nil {
		return nil
	}

	if insertPos.data == it {
		return insertPos // 'it' exists
	}

	//	log.Printf("insertPos:%s Target:%s\n", insertPos.data, it)

	newNode := Node{parent: insertPos, data: it} // create a new node for 'it'

	// connect the new node with it's parents
	if insertPos.data > it {
		insertPos.left = &newNode
	} else {
		insertPos.right = &newNode
	}

}

// 删除
// @Override
func (tree *AvlTree) Remove(item string) {

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

	//log.Printf("当前search:%s  目标:%s \n", sub.data, it)

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
////////////////////// 平衡 ///////////////////
func BalancedFactor(subNode *Node) int {
	return Height(LChild(subNode)) - Height(RChild(subNode))
}

///////////////////// 辅助函数 ////////////////
func LChild(sub *Node) *Node {
	return BST.LChild(sub)
}

func RChild(sub *Node) *Node {
	return BST.RChild(sub)
}

func Height(sub *Node) int {
	return BST.Height(sub)
}
