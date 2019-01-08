// 线性表-序列
package Vector

import (
	"errors"
)

const (
	DefaultCap = 64
)

type Item string

type vector struct {
	data    []Item
	lastIdx int
}

func New() *vector {
	return &vector{make([]Item, 0, DefaultCap), 0}
}

func (this *vector) InsertOrUpdate(idx int, it Item) error {
	//Range or Dichotomy
	if idx > this.lastIdx {
		return errors.New("over the range of list")
	}

	this.data[idx] = it
	return nil
}

func (this *vector) Get(idx int) (Item, error) {
	//Range or Dichotomy
	if idx > this.lastIdx {
		return Item(""), errors.New("over the range of list")
	}

	return this.data[idx], nil
}

// Replace will replace the old data on specified idx
// and return the older one.
func (this *vector) Replace(idx int, it Item) (Item, error) {
	if idx > this.lastIdx {
		return Item(""), errors.New("over the range of list")
	}

	old := this.data[idx]
	this.data[idx] = it

	return old, nil
}

func (this *vector) PushBack(it Item) int {
	//Range or Dichotomy
	this.data = append(this.data, it)
	this.lastIdx++
	return this.lastIdx
}

func (this *vector) PopBack() Item {
	return this.data[this.lastIdx]
}

func (this *vector) Search(it Item) int {
	//Range or Dichotomy
	for i, _ := range this.data {
		if it == this.data[i] {
			return i
		}
	}

	return -1
}

func (this *vector) Array() []Item {
	ret := make([]Item, this.lastIdx+1)
	copy(ret, this.data[:this.lastIdx+1])
	return ret
}

// Sort will resort over the data storage.
// If 'asc' is true,it will be sorted by ascending order while
// by descending order if false.
func (this *vector) Sort(asc bool) {
	// Adapt 'sort' package which combines quick sort  and heap sort
	// For details,you can find on 'https://sourcegraph.com/github.com/golang/go/-/blob/src/sort/sort.go#L183:6'

}

func (this *vector) Remove() {

}

// 唯一化
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func (this *vector) Uniquify() {

}

/////////////  算法 //////////////////

// SearchDichotomy can find a specified item in traversal way.
// O(1) ~ O(n)
func (this *vector) Find(it Item, lo, hi int) int {
	if hi > this.lastIdx {
		hi = this.lastIdx
	}

	for lo <= hi {
		if this.data[lo] == it {
			return lo
		}
		lo++
	}

	return -1
}

// SearchBinary can search a specified item in binary way.
// O(1.5*logn)
func (this *vector) SearchBinary(it Item, lo, hi int) int {

	return -1
}

// SearchBinary can search a specified item in fibnacci way.
func (this *vector) SearchFibnacci(it Item) int {
	return -1
}

// SortBubble can sort the data set in Bubble-Sort way.
func (this *vector) SortBubble() {

}

// SortQuick can sort the data set in Quick-Sort way.
func (this *vector) SortQuick() {

}

// SortHeap can sort the data set in Heap-Sort way.
func (this *vector) SortHeap() {

}

// SortMerge can sort the data set in Merge-Sort way.
func (this *vector) SortMerge() {

}
