package BinaryTree

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	"github.com/bradleyjkemp/memviz"
	"github.com/stretchr/testify/assert"
)

// 生成gv图
func TestBinaryTree(t *testing.T) {

	tree := New(Item("D"))

	st := []Item{"D", "B", "F", "A", "C", "E"}
	for i := 0; i < 6; i++ {
		tree.InsertItem(st[i])
	}

	buf := &bytes.Buffer{}
	memviz.Map(buf, tree)

	ioutil.WriteFile("./BinaryTree.gv", buf.Bytes(), os.ModePerm)
	exec.Command("dot", "-Tpng", "./BinaryTree.gv", "-o", "./BinaryTree.png").Run()
}

// 镜像翻转
func TestMirror(t *testing.T) {
	tree := New(Item("D"))

	st := []Item{"D", "B", "F", "A", "C", "E"}
	for i := 0; i < 6; i++ {
		tree.InsertItem(st[i])
	}

	result := Item("")
	it := func(i Item) bool {
		fmt.Printf(" %s", i)
		result += i
		return true
	}

	resultMirror := Item("")
	itMirror := func(i Item) bool {
		fmt.Printf(" %s", i)
		resultMirror += i
		return true
	}

	tree.PreorderTraversal(tree.Root, it)
	tree.Mirror(tree.Root)
	tree.PreorderTraversal(tree.Root, itMirror)

	assert.Equal(t, Item("DBACFE"), result)
	assert.Equal(t, Item("DFEBCA"), resultMirror)
}

// 前序
func TestPreorderTraversal(t *testing.T) {
	tree := New(Item("D"))

	st := []Item{"D", "B", "F", "A", "C", "E"}
	for i := 0; i < 6; i++ {
		tree.InsertItem(st[i])
	}

	result := Item("")
	it := func(i Item) bool {

		fmt.Printf(" %s", i)
		result += i
		return true
	}

	tree.PreorderTraversal(tree.Root, it)
	assert.Equal(t, Item("DBACFE"), result)

}

// 中序
func TestInorderTraversal(t *testing.T) {
	tree := New(Item("D"))

	st := []Item{"D", "B", "F", "A", "C", "E"}
	for i := 0; i < 6; i++ {
		tree.InsertItem(st[i])
	}

	result := Item("")
	it := func(i Item) bool {
		fmt.Printf(" %s", i)
		result += i
		return true
	}

	tree.InorderTraversal(tree.Root, it)
	assert.Equal(t, Item("ABCDEF"), result)
}

// 后序
func TestPostorderTraversal(t *testing.T) {
	tree := New(Item("D"))

	st := []Item{"D", "B", "F", "A", "C", "E"}
	for i := 0; i < 6; i++ {
		tree.InsertItem(st[i])
	}

	result := Item("")
	it := func(i Item) bool {
		fmt.Printf(" %s", i)
		result += i
		return false
	}

	tree.PostorderTraversal(tree.Root, it)
	assert.Equal(t, Item("ACBEFD"), result)
}
