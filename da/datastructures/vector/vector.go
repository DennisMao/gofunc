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
	if lo == -1 {
		lo = 0
	}

	if hi == -1 {
		hi = this.lastIdx
	}

	return searchBinarySearch(this.data, it, lo, hi)
}

func searchBinarySearch(arr []Item, it Item, lo, hi int) int {

	pivot := (lo + hi) / 2

	if lo == hi && it != arr[pivot] {
		return -1
	}

	log.Printf("pivot:%d  value:%s \n", pivot, arr[pivot])

	switch {
	case it < arr[pivot]:
		return searchBinarySearch(arr, it, lo, pivot-1)
	case it > arr[pivot]:
		return searchBinarySearch(arr, it, pivot+1, hi)
	case it == arr[pivot]:
		return pivot
	}

	return -1
}

var fibArray []int

func init() {
	fibArray := make([]int, 20)

	for n := 0; n < 20; n++ {
		fibArray[n] = fib(n)
	}
}

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

func GetFib(n int) int {
	if n > len(fibArray)-1 {
		return fib(n)
	}

	return fibArray[n]
}

// SearchBinary can search a specified item in fibnacci way.
func (this *vector) SearchFibnacci(it Item, lo, hi int) int {
	return searchFibnacci(this.data, it, 0, this.lastIdx)
}

func searchFibnacci(arr []Item, it Item, lo, hi int) int {

	k := hi - 1
	pivot := GetFib(k)
	if pivot > hi {
		for pivot > hi {
			k--
			pivot = GetFib(k)
		}
	}

	if pivot < lo {
		pivot = (lo + hi) / 2 //当[lo,hi]处在FIB区间内的时候用二分法
	}

	if lo == hi && it != arr[pivot] {
		return -1
	}

	log.Printf("arr:%v pivot:%d  value:%s \n", arr, pivot, arr[pivot])

	switch {
	case it < arr[pivot]:
		return searchFibnacci(arr, it, lo, pivot-1)
	case it > arr[pivot]:
		return searchFibnacci(arr, it, pivot+1, hi)
	case it == arr[pivot]:
		return pivot
	}

	return -1
}

// SearchInsert can search a specified item in insert-search way.
func (this *vector) SearchInsert(it Item) int {
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
	if len(this.data) < 2 {
		return
	}

	//this.data = quickSort(this.data)
	quickSortInplace(this.data)
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
func quickSortInplace(array []Item) {
	if len(array) < 2 {
		return
	}
	log.Printf("Raw== array:%v    \n", array)

	pivotIdx := 0            //基准点
	pivot := array[pivotIdx] //哨兵
	i := 1
	j := len(array) - 1

	for i <= j {

		if array[i] < pivot {
			array[i], array[pivotIdx] = array[pivotIdx], array[i]
			i++
			pivotIdx++
		} else {
			array[j], array[i] = array[i], array[j]
			j--
		}

		time.Sleep(200 * time.Millisecond)
	}

	array[pivotIdx] = pivot
	log.Printf("Sort== arrayLeft:%v   Idx: %s %d  arrayRight:%v \n", array[:pivotIdx], pivot, pivotIdx, array[pivotIdx+1:])
	quickSortInplace(array[:pivotIdx])
	quickSortInplace(array[pivotIdx+1:])

}

// SortHeap can sort the data set in Heap-Sort way.
func (this *vector) SortHeap() {
	if len(this.data) < 2 {
		return
	}

	heapSort(this.data, 0, this.lastIdx)

}

func heapSort(arr []Item, left, right int) {
	first := left
	lo := 0
	hi := right - left

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(arr, i, hi, first)
	}

	log.Printf("arr:%v left:%s \n", arr, arr[left])

	// Pop elements, largest first, into end of data.
	for i := hi - 1; i >= 0; i-- {
		arr[first], arr[first+i] = arr[first+i], arr[first]
		siftDown(arr, lo, i, first)
		log.Printf("arr:%v first:%v arr[i]:%v \n", arr, arr[first], arr[first+i])
	}
}

// siftDown implements the heap property on data[lo, hi).
// first is an offset into the array where the root of the heap lies.
func siftDown(arr []Item, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && arr[first+child] < arr[first+child+1] {
			child++
		}
		if !(arr[first+root] < arr[first+child]) {
			return
		}
		arr[first+root], arr[first+child] = arr[first+child], arr[first+root]

		root = child
	}
}

// SortMerge can sort the data set in Merge-Sort way.
func (this *vector) SortMerge() {
	if len(this.data) < 2 {
		return
	}

	this.data = mergeSortSort(this.data)

}

func mergeSortSort(arr []Item) []Item {
	if len(arr) < 2 {
		return arr
	}

	pivot := len(arr) / 2

	left := mergeSortSort(arr[:pivot])
	right := mergeSortSort(arr[pivot:])

	return mergeSortMerge(left, right)
}

func mergeSortMerge(left, right []Item) []Item {
	lenLeft := len(left)
	lenRight := len(right)

	ret := make([]Item, lenLeft+lenRight)
	i := 0
	j := 0

	// Scan both left and right arrays
	for i < lenLeft && j < lenRight {

		if left[i] < right[j] {
			ret[i+j] = left[i]
			i++
			continue
		} else {
			ret[i+j] = right[j]
			j++
			continue
		}
	}

	log.Printf("i:%d j:%d", i, j)

	// Append elements left
	if i < lenLeft {
		for ; i < lenLeft; i++ {
			ret[i+j] = left[i]
		}

	}
	if j < lenRight {
		for ; j < lenRight; j++ {
			ret[i+j] = right[j]
		}
	}

	log.Printf("left:%v right:%v merge:%v \n", left, right, ret)

	return ret
}

// SortShell can sort the data set in Shell-Sort way.
func (this *vector) SortShell() {
	if len(this.data) < 2 {
		return
	}

	for i := 1; i < len(this.data); i++ {
		j := i
		for ; j > 0 && this.data[j-1] > this.data[j]; j-- {
			this.data[j], this.data[j-1] = this.data[j-1], this.data[j]
		}
		log.Printf("Pick:%s In:%d Arr:%v\n", this.data[j], j, this.data)
	}

}

// SortInsertion can sort the data set in Insertion-Sort way.
func (this *vector) SortInsertion() {
	if len(this.data) < 2 {
		return
	}

	for i := 1; i < len(this.data); i++ {
		temp := this.data[i]
		j := i
		for ; j > 0 && this.data[j-1] >= temp; j-- {
			this.data[j], this.data[j-1] = this.data[j-1], this.data[j]
		}

		log.Printf("Pick:%s In:%d Arr:%v\n", this.data[j], j, this.data)
		this.data[j] = temp
	}

}
