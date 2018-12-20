package BinaryTree

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	"github.com/bradleyjkemp/memviz"
)

// 前序
func TestPreorderTraversal(t *testing.T) {
	tree := New(Item("A"))

	st := 'A'
	for i := 1; i < 6; i++ {
		tree.InsertItem(Item(st + rune(i)))
	}

	it := func(i Item) bool {

		fmt.Printf(" %s", i)
		return true
	}
	tree.PreorderTraversal(nil, it)

	buf := &bytes.Buffer{}
	memviz.Map(buf, tree)
	fmt.Println(buf.String())
	cupaloy.SnapshotT(t, buf.Bytes())

}

// 中序
func TestInorderTraversal(t *testing.T) {
	tree := New(Item("A"))

	st := 'A'
	for i := 1; i < 6; i++ {
		tree.InsertItem(Item(st + rune(i)))
	}

	it := func(i Item) bool {
		fmt.Printf(" %s", i)
		return true
	}

	tree.InorderTraversal(nil, it)

}

// 后序
func TestPostorderTraversal(t *testing.T) {
	tree := New(Item("A"))

	st := 'A'
	for i := 1; i < 6; i++ {
		tree.InsertItem(Item(st + rune(i)))
	}

	it := func(i Item) bool {
		fmt.Printf(" %s", i)
		return true
	}

	tree.PostorderTraversal(nil, it)
}
