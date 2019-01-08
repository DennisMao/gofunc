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
	testItem       = []Item{"A", "B", "C", "D", "E", "F", "G"}
	testRandomItem = []Item{"D", "B", "C", "F", "G", "A", "E"}
)

func NewVector() *vector {
	vector := New()

	for i, _ := range testItem {
		vector.PushBack(testItem[i])
	}

	return vector
}

func NewRandomVector() *vector {
	vector := New()

	for i, _ := range testRandomItem {
		vector.PushBack(testRandomItem[i])
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

	vector := NewVector()
	value := vector.Find(Item("A"), -1, -1)
	assert.Equal(t, 0, value)

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

	vector := NewRandomVector()

	assert.Equal(t, testRandomItem, vector.Array())
	vector.SortBubble()
	assert.Equal(t, testItem, vector.Array())

}

func TestSortQuick(t *testing.T) {

	vector := NewRandomVector()

	assert.Equal(t, testRandomItem, vector.Array())
	vector.SortQuick()
	assert.Equal(t, testItem, vector.Array())

}

func TestSortHeap(t *testing.T) {

	//assert.Equal(t, Item("DBACFE"), result)

}

func TestSortMerge(t *testing.T) {

	//assert.Equal(t, Item("DBACFE"), result)

}
