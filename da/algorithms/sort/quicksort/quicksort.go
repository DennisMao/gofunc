package quicksort

import (
	"log"
	"time"
)

// SortQuick can sort the data set in Quick-Sort way.
// Avg : O(nlogn)
// Worst： O(n^2)
func SortQuick(array *[]string) {
	if len(*array) < 2 {
		return
	}

	ret := quickSort(*array)
	array = &ret
}

func SortQuickInplace(array *[]string) {
	if len(*array) < 2 {
		return
	}

	//this.data = quickSort(array)
	quickSortInplace(*array)
}

func quickSort(array []string) []string {
	if len(array) < 2 {
		return array
	}

	median := array[len(array)/2] //中值点
	//log.Printf("median:%s  raw:%v\n", median, array)

	low_part := make([]string, 0)   //values less than median
	high_part := make([]string, 0)  // values higher than median
	pivot_part := make([]string, 0) // values equal to median

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
//  [  <pivot ]lo[  U ]hi[ >pivot ]
func quickSortInplace(array []string) {
	if len(array) < 2 {
		return
	}
	log.Printf("Raw== array:%v    \n", array)

	pivotIdx := 0            //轴点
	pivot := array[pivotIdx] //轴点值
	lo := 1
	hi := len(array) - 1

	for {

		for array[lo] < pivot && lo < hi {
			lo++
		}
		array[lo], array[pivotIdx] = array[pivotIdx], array[lo]
		pivotIdx = lo

		for array[hi] > pivot && lo < hi {
			hi--
		}
		array[hi], array[pivotIdx] = array[pivotIdx], array[hi]
		pivotIdx = hi

		if lo == hi {
			break
		}

		log.Printf("CurArray: %#v \n", array)
		time.Sleep(200 * time.Millisecond)
	}

	array[pivotIdx] = pivot
	log.Printf("Sort== arrayLeft:%v   Idx: %s %d  arrayRight:%v \n", array[:pivotIdx], pivot, pivotIdx, array[pivotIdx+1:])
	quickSortInplace(array[:pivotIdx])
	quickSortInplace(array[pivotIdx+1:])

}
