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
	testItem       = []Item{"A", "A", "A", "B", "C", "D", "E", "F", "G"}
	testRandomItem = []Item{"D", "B", "C", "A", "F", "G", "A", "E", "A"}
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
	value := vector.Find(Item("B"), -1, -1)
	assert.Equal(t, 3, value)

	//assert.Equal(t, Item("DBACFE"), result)

}

// 二分法
func TestSearchBinary(t *testing.T) {

	vector := NewVector()
	value := vector.SearchBinary(Item("B"), -1, -1)
	assert.Equal(t, 3, value)

}

// 斐波那契
func TestSearchFib(t *testing.T) {

	vector := NewVector()
	value := vector.SearchFibnacci(Item("B"), -1, -1)
	assert.Equal(t, 3, value)
}

// 插值法
func TestSearchInsert(t *testing.T) {

	vector := NewVector()
	value := vector.SearchInsert(Item("B"), -1, -1)
	assert.Equal(t, 3, value)
}

/////////////////////////// 排序 //////////////////////////
// 起泡
func TestSortBubble(t *testing.T) {

	vector := NewRandomVector()

	assert.Equal(t, testRandomItem, vector.Array())
	vector.SortBubble()
	assert.Equal(t, testItem, vector.Array())

}

// 快排
func TestSortQuick(t *testing.T) {

	vector := NewRandomVector()

	assert.Equal(t, testRandomItem, vector.Array())
	vector.SortQuick()
	assert.Equal(t, testItem, vector.Array())

}

// 堆排
func TestSortHeap(t *testing.T) {

	vector := NewRandomVector()

	assert.Equal(t, testRandomItem, vector.Array())
	vector.SortHeap()
	assert.Equal(t, testItem, vector.Array())
}

// 归并
func TestSortMerge(t *testing.T) {

	vector := NewRandomVector()

	assert.Equal(t, testRandomItem, vector.Array())
	vector.SortMerge()
	assert.Equal(t, testItem, vector.Array())

}

// 希尔
func TestSortShell(t *testing.T) {

	vector := NewRandomVector()

	assert.Equal(t, testRandomItem, vector.Array())
	vector.SortInsertion()
	assert.Equal(t, testItem, vector.Array())
}

// 直接插入
func TestSortInsertion(t *testing.T) {

	for i := 0; i < 10; i++ {
		t.Log(fib(i))
	}

	vector := NewRandomVector()

	assert.Equal(t, testRandomItem, vector.Array())
	vector.SortInsertion()
	assert.Equal(t, testItem, vector.Array())
}
