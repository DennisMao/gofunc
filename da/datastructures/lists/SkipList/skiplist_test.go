package skiplist

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"sort"
	"testing"

	"github.com/bradleyjkemp/memviz"
	"github.com/stretchr/testify/assert"
)

// Returns a target data structure object we are going to test.
func NewExampleList() *SkipList {
	l := Create(0)

	k := int64(1)
	score := float64(1.0)
	for i := 0; i < 20; i++ {
		l.Insert(score, k, "1")
		k++
		score += 1.0
	}

	return l
}

// Rand 算法
func TestRandLevel(t *testing.T) {
	l := Create(0)

	res := make([]int, 100)
	for i, _ := range res {
		res[i] = l.RandomLevel()
	}
	sort.Sort(sort.IntSlice(res))

	t.Log(res)
}

// 生成gv图
func TestSkipList(t *testing.T) {

	l := NewExampleList()

	buf := &bytes.Buffer{}
	memviz.Map(buf, l)

	ioutil.WriteFile("./SkipList.gv", buf.Bytes(), os.ModePerm)
	exec.Command("dot", "-Tpng", "./SkipList.gv", "-o", "./SkipList.png").Run()
}

func TestRange(t *testing.T) {
	l := NewExampleList()

	for i := MAX_LEVEL; i > 0; i-- {
		r := l.RangeLevel(i)
		fmt.Printf("========= level:%d ========\n", i)
		for _, node := range r {
			if node.key == nil {
				continue
			}
			fmt.Printf("(%.1f,%d)>", node.score, node.key.(int64))
		}
		fmt.Printf("\n")

	}

}

func TestSearch(t *testing.T) {
	l := NewExampleList()

	node := l.Find(float64(1.0), int64(1))
	if node == nil {
		assert.Fail(t, "Element '1' is losing,check 'Insert' function.")
	}

	assert.Equal(t, "1", node.Value.(string))
}

func TestInsert(t *testing.T) {
	l := NewExampleList()

	l.Insert(float64(1024.0), int64(1024), "testInsert")

	node := l.Find(float64(1024.0), int64(1024))
	if node == nil {
		assert.Fail(t, "Element '10' is losing,check 'Insert' function.")
	}
	assert.Equal(t, "testInsert", node.Value.(string))
}

//func TestReverse(t *testing.T) {
//	l := NewExampleList()

//	bNode := l.Front()
//	assert.Equal(t, Item("A"), bNode.Data())
//	l.Reverse()

//	bNode = l.Front()
//	if bNode == nil {
//		assert.Fail(t, "Back element is nil")
//		return
//	}

//	assert.Equal(t, Item("F"), bNode.Data())
//}
