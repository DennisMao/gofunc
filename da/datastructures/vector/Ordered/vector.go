// 线性表-序列
package Vector

import (
	"errors"
	"log"
	"time"
)

const (
	DefaultCap = 64
	IsDebug    = true
)

type Item string

type vector struct {
	data    []Item
	lastIdx int
}

func New() *vector {
	return &vector{make([]Item, 0, DefaultCap), -1}
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
// lo,hi specifies the search range of internal data array,setting '-1'
// can automatically
// O(1) ~ O(n)
func (this *vector) Find(it Item, lo, hi int) int {
	if lo == -1 {
		lo = 0
	}

	if hi == -1 {
		hi = this.lastIdx
	}

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
	dataLen := this.lastIdx + 1
	for s := 0; s < dataLen; s++ {
		for e := 0; e < dataLen; e++ {

			//log.Printf("s:%d e:%d\n", s, e)

			if this.data[s] < this.data[e] {
				//log.Printf("swap A:%s B:%s \n", this.data[s], this.data[e])
				this.data[s], this.data[e] = this.data[e], this.data[s]
			}

		}
	}
}

// SortQuick can sort the data set in Quick-Sort way.
// Avg : O(nlogn)
// Worst： O(n^2)
func (this *vector) SortQuick() {
	//this.data = quickSort(this.data)
	quickSortInplace(this.data, 0, this.lastIdx)
}

func quickSort(array []Item) []Item {
	if len(array) < 2 {
		return array
	}

	median := array[len(array)/2] //中值点
	//log.Printf("median:%s  raw:%v\n", median, array)

	low_part := make([]Item, 0)   //values less than median
	high_part := make([]Item, 0)  // values higher than median
	pivot_part := make([]Item, 0) // values equal to median

	for i, _ := range array {
		switch {
		case array[i] == median:
			pivot_part = append(pivot_part, array[i])
		case array[i] < median:
			low_part = append(low_part, array[i])
		case array[i] > median:
			high_part = append(high_part, array[i])
		}
	}

	//分治
	low_part = quickSort(low_part)
	high_part = quickSort(high_part)

	//log.Printf("low_part:%v \n", low_part)
	//log.Printf("pivot_part:%v \n", pivot_part)
	//log.Printf("high_part:%v \n", high_part)

	//汇总结果
	low_part = append(low_part, pivot_part...)
	low_part = append(low_part, high_part...)
	return low_part
}

// 原地分割
func quickSortInplace(array []Item, left, right int) {

	privotIdx := right / 2

	for right > left {
		newPrivotIdx := quickSortPartition(array, left, right, privotIdx)
		log.Printf("oldPrivotIdx:%d  newPrivotIdx:%d array:%v\n", privotIdx, newPrivotIdx, array)
		time.Sleep(1 * time.Second)

		quickSortInplace(array, left, newPrivotIdx)
		quickSortInplace(array, newPrivotIdx+1, right)
	}
}

//从左往右找出合适的位置
func quickSortPartition(array []Item, left, right, privotIdx int) int {
	privotValue := array[privotIdx]

	//log.Printf("raw array: %v \n", array)

	//将privot值移到最右缓存起来
	array[privotIdx], array[right] = array[right], array[privotIdx]
	//log.Printf("setBuf array: %v \n", array)

	suitIdx := left
	for i := left; i < right; i++ {

		if array[i] < privotValue {
			array[i], array[suitIdx] = array[suitIdx], array[i] //移动合适位置
			suitIdx++
		}
	}
	//log.Printf("finding suitIdx :%d array: %v \n", suitIdx, array)

	//将privot放到合适位置
	array[suitIdx], array[right] = array[right], array[suitIdx]
	//log.Printf("reset  array: %v \n", array)
	return suitIdx
}

// SortHeap can sort the data set in Heap-Sort way.
func (this *vector) SortHeap() {

}

// SortMerge can sort the data set in Merge-Sort way.
func (this *vector) SortMerge() {

}
