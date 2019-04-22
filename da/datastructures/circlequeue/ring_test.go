package Ring

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
func NewExampleRing() *ring {
	l := New(0)

	st := []Item{"A", "B", "C", "D", "E", "F"}
	for i, _ := range st {
		l.PushBack(st[i])
	}

	return l
}

// 生成gv图
func TestRing(t *testing.T) {

	l := NewExampleList()

	buf := &bytes.Buffer{}
	memviz.Map(buf, l)

	ioutil.WriteFile("./Ring.gv", buf.Bytes(), os.ModePerm)
	exec.Command("dot", "-Tpng", "./Ring.gv", "-o", "./Ring.png").Run()
}

func TestPushBack(t *testing.T) {
	l := NewExampleList()
	l.PushBack(Item("G"))

	assert.Equal(t, 7, l.Len())
}

func TestSearch(t *testing.T) {
	l := NewExampleList()

	bNode := l.Search("B")

	assert.Equal(t, Item("B"), bNode.Data())
}

func TestInsert(t *testing.T) {
	l := NewExampleList()

	nodeF := l.Search("F")
	if nodeF == nil {
		assert.Fail(t, "Element 'F' is losing,check 'PushBack' function.")
	}

	l.Insert(Item("K"), nodeF)

	//log.Printf("Range:%v \n", l.RangeFrom(nodeF.Data()))

	assert.Equal(t, Item("K"), nodeF.next.Data())
}

func TestReverse(t *testing.T) {
	l := NewExampleList()

	bNode := l.Front()
	assert.Equal(t, Item("A"), bNode.Data())
	l.Reverse()

	bNode = l.Front()
	if bNode == nil {
		assert.Fail(t, "Back element is nil")
		return
	}

	assert.Equal(t, Item("F"), bNode.Data())
}
