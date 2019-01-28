package AvlTree

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
func TestBinarySearchTree(t *testing.T) {

	tree := New()

	st := []string{"D", "B", "F", "A", "C", "E"}
	for i := 0; i < 6; i++ {
		tree.Insert(st[i])
	}

	buf := &bytes.Buffer{}
	memviz.Map(buf, tree)

	ioutil.WriteFile("./BinarySearchTree.gv", buf.Bytes(), os.ModePerm)
	exec.Command("dot", "-Tpng", "./BinarySearchTree.gv", "-o", "./BinarySearchTree.png").Run()
}

// 搜索
func TestSearch(t *testing.T) {
	tree := New()

	st := []string{"D", "B", "F", "A", "C", "E"}
	for i := 0; i < 6; i++ {
		tree.Insert(st[i])
	}

	result := tree.Search("B")
	if result == nil {
		t.Fatal("Search failed")
	}

	assert.Equal(t, "B", result.data)
}

// 删除
func TestDelete(t *testing.T) {
	tree := New()

	st := []string{"D", "B", "F", "A", "C", "E"}
	for i := 0; i < 6; i++ {
		tree.Insert(st[i])
	}

	result := ""
	it := func(i string) bool {

		result += i
		return true
	}

	tree.InorderTraversal(tree.Root, it)
	assert.Equal(t, "ABCDEF", result)

	tree.Delete("B")

	result = ""
	tree.InorderTraversal(tree.Root, it)
	assert.Equal(t, "ACDEF", result)

}

// 前序
// Root [Left] [Right]
func TestPreorderTraversal(t *testing.T) {
	tree := New()

	st := []string{"D", "B", "F", "A", "C", "E"}
	for i := 0; i < 6; i++ {
		tree.Insert(st[i])
	}

	result := ""
	it := func(i string) bool {

		fmt.Printf(" %s", i)
		result += i
		return true
	}

	tree.PreorderTraversal(tree.Root, it)
	assert.Equal(t, "DBACFE", result)

}

// 中序
// [Left] Root [Right]
func TestInorderTraversal(t *testing.T) {
	tree := New()

	st := []string{"D", "B", "F", "A", "C", "E"}
	for i := 0; i < 6; i++ {
		tree.Insert(st[i])
	}

	result := ""
	it := func(i string) bool {
		fmt.Printf(" %s", i)
		result += i
		return true
	}

	tree.InorderTraversal(tree.Root, it)
	assert.Equal(t, "ABCDEF", result)
}

// 后序
// [Right] [Left] Root
func TestPostorderTraversal(t *testing.T) {
	tree := New()

	st := []string{"D", "B", "F", "A", "C", "E"}
	for i := 0; i < 6; i++ {
		tree.Insert(st[i])
	}

	result := ""
	it := func(i string) bool {
		fmt.Printf(" %s", i)
		result += i
		return false
	}

	tree.PostorderTraversal(tree.Root, it)
	assert.Equal(t, "ACBEFD", result)
}

// 层级遍历
// ROOT [h2] ... [hn]
func TestLevelTraversal(t *testing.T) {
	tree := New()

	st := []string{"D", "B", "F", "A", "C", "E"}
	for i := 0; i < 6; i++ {
		tree.Insert(st[i])
	}

	result := ""
	it := func(i string) bool {
		fmt.Printf(" %s", i)
		result += i
		return false
	}

	tree.LevelTraversal(tree.Root, it)
	assert.Equal(t, "DBFACE", result)
}
