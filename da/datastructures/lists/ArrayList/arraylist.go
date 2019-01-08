// 线性表-序列
package ArrayList

import (
	"errors"
)

const (
	DefaultCap = 256
)

type Item string

type arrayList struct {
	data    []Item
	lastIdx int
}

func New() *arrayList {
	return &arrayList{make([]Item, 0, 256), 0}
}

func (this *arrayList) InsertOrUpdate(idx int, it Item) error {
	//Range or Dichotomy
	if idx > this.lastIdx {
		return errors.New("over the range of list")
	}

	this.data[idx] = it
	return nil
}

// Replace will replace the old data on specified idx
// and return the older one.
func (this *arrayList) Replace(idx int, it Item) (Item, error) {
	if idx > this.lastIdx {
		return Item(""), errors.New("over the range of list")
	}

	old := this.data[idx]
	this.data[idx] = it

	return old, nil
}

func (this *arrayList) PushBack(it Item) int {
	//Range or Dichotomy
	this.data = append(this.data, it)
	this.lastIdx++
	return this.lastIdx
}

func (this *arrayList) PopBack() Item {
	return this.data[this.lastIdx]
}

func (this *arrayList) Search(it Item) int {
	//Range or Dichotomy
	for i, _ := range this.data {
		if it == this.data[i] {
			return i
		}
	}

	return -1
}

// Sort will resort over the data storage.
// If 'asc' is true,it will be sorted by ascending order while
// by descending order if false.
func (this *arrayList) Sort(asc bool) {
	// Adapt 'sort' package which combines quick sort  and heap sort
	// For details,you can find on 'https://sourcegraph.com/github.com/golang/go/-/blob/src/sort/sort.go#L183:6'

}

func (this *arrayList) Remove() {

}

// 唯一化
func (this *arrayList) Uniquify() {

}


/////////////  算法 //////////////////

// SearchDichotomy can find a specified item in traversal way.
// O(1) ~ O(n)
func(this *arrayList) Find(it Item,lo,hi int) int{
	if  hi > this.lastIdx  {
		hi = this.lastIdx
	}
	
	
	for lo <= hi {
		
	}
	
}

// SearchBinary can search a specified item in binary way.
// O(1.5*logn)
func (this *arrayList) SearchBinary(it Item) int {
	return -1
}

// SearchBinary can search a specified item in fibnacci way.
func (this *arrayList) SearchFibnacci(it Item) int {
	return -1
}

// SortBubble can sort the data set in Bubble-Sort way.
func （this *arrayList）SortBubble() {
	
}

// SortQuick can sort the data set in Quick-Sort way.
func (this *arrayList) SortQuick() {
	
}

// SortHeap can sort the data set in Heap-Sort way.
func (this *arrayList) SortHeap() {
	
}
