package LinkedList

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	"github.com/bradleyjkemp/memviz"
	"github.com/stretchr/testify/assert"
)

// Return an example with datas "A,B,C,D,E,F"
func NewExampleList() *arrayList {
	list := New()

	st := []Item{"A", "B", "C", "D", "E", "F"}
	for i, _ := range st {
		list.Insert(st[i])
	}

	return list
}

// 生成gv图
func TestLinkedList(t *testing.T) {

	list := NewExampleList()

	buf := &bytes.Buffer{}
	memviz.Map(buf, list)

	ioutil.WriteFile("./LinkedList.gv", buf.Bytes(), os.ModePerm)
	exec.Command("dot", "-Tpng", "./LinkedList.gv", "-o", "./LinkedList.png").Run()
}

func TestInsert(t *testing.T) {
	list := NewExampleList()
	list.Insert(Item("G"))

	assert.Equal(t, 7, list.Size())
}

func TestSearch(t *testing.T) {
	list := NewExampleList()

	bNode := list.Search("B")

	assert.Equal(t, Item("C"), bNode.next.data)
}
