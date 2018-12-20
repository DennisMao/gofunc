// Name: Binary tree
package BinaryTree

const (
	CompareEqual = 1 << iota
	CompareLess
	CompareMore
)

// ItemIterator allows callers of Ascend* to iterate in-order over portions of
// the tree.  When this function returns false, iteration will stop and the
// associated Ascend* function will immediately return.
type ItemIterator func(i Item) bool

// ItemCompare allows callers specify the compare function for items searching.
// Return the compare state
type ItemCompare func(i Item) int

type Item string

type Node struct {
	left   *Node
	right  *Node
	parent *Node
	data   Item
}

type BinaryTree struct {
	root *Node
}

func New(i Item) *BinaryTree {

	return &BinaryTree{
		root: &Node{nil, nil, nil, i},
	}

}

func (this *BinaryTree) InsertItem(item Item) {

	curNode := this.root
	for {
		if curNode.left == nil {
			newNode := &Node{
				parent: curNode,
				data:   item,
			}
			curNode.left = newNode
			return
		} else {
			curNode = curNode.left
		}

		if curNode.right == nil {
			newNode := &Node{
				parent: curNode,
				data:   item,
			}
			curNode.right = newNode
			return
		} else {
			curNode = curNode.right
		}
	}
}

func (tree *BinaryTree) SearchItem(f ItemCompare) (*Node, bool) {
	if tree.root == nil {
		return nil, false
	}
	currentNode := tree.root
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

// 前序
func (tree *BinaryTree) PreorderTraversal(subNode *Node, it ItemIterator) {
	if subNode == nil {
		return
	}

	it(subNode.data)

	tree.PreorderTraversal(subNode.left, it)
	tree.PreorderTraversal(subNode.right, it)

}

// 中序
func (tree *BinaryTree) InorderTraversal(subNode *Node, it ItemIterator) {

	if subNode == nil {
		return
	}

	tree.InorderTraversal(subNode.left, it)
	it(subNode.data)
	tree.InorderTraversal(subNode.right, it)
}

// 后序
func (tree *BinaryTree) PostorderTraversal(subNode *Node, it ItemIterator) {
	if subNode == nil {
		return
	}

	if subNode.right != nil {
		tree.PostorderTraversal(subNode.right, it)
	}
	if subNode.left != nil {
		tree.PostorderTraversal(subNode.left, it)
	}
	it(subNode.data)

}
