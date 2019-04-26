package quick

import (
	"log"
)

func SortQuick(data *[]string) {
	d := *data
	lenD := len(d)
	if lenD < 2 {
		return
	}

	quickSortInplace(d)

}

// 原地快排
func quickSortInplace(data []string) {
	if len(data) < 2 {
		return
	}

	pivot := partition(data) // 轴点坐标
	quickSortInplace(data[:pivot])
	quickSortInplace(data[pivot+1:])

	return
}

func partition(data []string) (pivot int) {
	pivot = 0               // 轴点
	pivotVal := data[pivot] //哨兵

	lo := 1
	hi := len(data) - 1
	for lo <= hi {
		// |  < Pivot |  U  |  > Pivot |
		// Prev < Pivot
		if data[lo] < pivotVal {
			data[lo], data[pivot] = data[pivot], data[lo]
			pivot = lo
			lo++
		} else {
			// Pivot <= Tail
			data[hi], data[lo] = data[lo], data[hi]
			hi--
		}

		log.Printf("CurArr: %#v", data)
	}

	data[pivot] = pivotVal
	log.Printf("EndArr: %#v", data)
	return pivot
}
