// 线性表-单向链表
package LinkedList

type Item string

type linkedList struct {
	root *Node
	last *Node
	size int
}

type Node struct {
	data Item

	// pre *Node  //双向链表的话增加前置节点
	next *Node
}

func New() *linkedList {
	newNode := &Node{}
	return &linkedList{newNode, newNode, 0}
}

func (this *linkedList) Size() int {
	return this.size
}

func (this *linkedList) Search(it Item) *Node {
	return search(this.root, it)
}

func search(curNode *Node, it Item) *Node {
	if curNode.next == nil {
		return nil
	}

	if curNode.data == it {
		return curNode
	}

	return search(curNode.next, it)
}

func (this *linkedList) Insert(it Item) {
	newNode := &Node{
		data: it,
		next: nil,
	}

	this.last.next = newNode
	this.last = newNode
	this.size++
}

func (this *linkedList) Reverse() {

}

func (this *linkedList) Remove() {

}

func reverse(cur, next *Node) *Node {

	var endPoint *Node

	if next.next != nil {
		reverse(next, next.next)
	}

	endPoint = next
	next.next = cur

	return endPoint
}
