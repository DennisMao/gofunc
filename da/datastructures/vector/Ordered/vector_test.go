package Vector

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	"github.com/bradleyjkemp/memviz"
	"github.com/stretchr/testify/assert"
)

var (
	testItem = []Item{"A", "B", "C", "D", "E", "F", "G"}
)

func NewVector() *vector {
	vector := New()

	for i, _ := range testItem {
		vector.PushBack(testItem[i])
	}

	return vector
}

// 生成gv图
func TestVector(t *testing.T) {

	vector := NewVector()

	buf := &bytes.Buffer{}
	memviz.Map(buf, vector)

	ioutil.WriteFile("./Vector.gv", buf.Bytes(), os.ModePerm)
	exec.Command("dot", "-Tpng", "./Vector.gv", "-o", "./Vector.png").Run()
}

// 插入
func TestInsert(t *testing.T) {
	vector := NewVector()
	value, _ := vector.Get(0)
	assert.Equal(t, Item("A"), value)

	vector.InsertOrUpdate(0, "Q")
	value, _ = vector.Get(0)
	assert.Equal(t, Item("Q"), value)

}

/////////////////////////// 查找 //////////////////////////
// 遍历法
func TestSearch(t *testing.T) {

	//assert.Equal(t, Item("DBACFE"), result)

}

// 二分法
func TestSearchBinary(t *testing.T) {

	//assert.Equal(t, Item("DBACFE"), result)

}

// 斐波那契
func TestSearchFib(t *testing.T) {

	//assert.Equal(t, Item("DBACFE"), result)

}

/////////////////////////// 排序 //////////////////////////
// 起泡
func TestSortBubble(t *testing.T) {

	//assert.Equal(t, Item("DBACFE"), result)

}

func TestSortQuick(t *testing.T) {

	//assert.Equal(t, Item("DBACFE"), result)

}

func TestSortHeap(t *testing.T) {

	//assert.Equal(t, Item("DBACFE"), result)

}

func TestSortMerge(t *testing.T) {

	//assert.Equal(t, Item("DBACFE"), result)

}
