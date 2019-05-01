package quick

import (
	"log"
)

func SortMerge(data *[]string) {
	d := *data
	lenD := len(d)
	if lenD < 2 {
		return
	}

	newD := sortSort(d)
	log.Printf("End: %#v\n", newD)

	return
}

func sortSort(data []string) []string {
	lenD := len(data)
	if lenD < 2 {
		return data
	}
	if lenD == 2 {
		if data[0] <= data[1] {
			return data
		} else {
			data[0], data[1] = data[1], data[0]
			return data
		}
	}

	log.Printf("CurSort %v\n", data)
	pivot := len(data) / 2 //分割点
	leftU := sortSort(data[:pivot])
	rightU := sortSort(data[pivot:])

	mergeU := sortMerge(leftU, rightU)
	log.Printf("Pivot:%s Left:%v Right:%v  Result:%v\n", data[pivot], leftU, rightU, mergeU)
	return mergeU
}

func sortMerge(left, right []string) []string {
	log.Printf("Left:%v Right:%v\n", left, right)

	lenLeft := len(left)
	lenRight := len(right)

	if lenLeft+lenRight == 0 {
		return nil
	}
	if lenLeft == 0 {
		return right
	}
	if lenRight == 0 {
		return left
	}

	//  LeftU   RightU
	//     \      /
	//     SortedU
	newU := make([]string, lenLeft+lenRight)

	i := 0
	j := 0
	k := 0

	for i < lenLeft && j < lenRight {
		if left[i] <= right[j] {
			newU[k] = left[i]
			i++
		} else {
			newU[k] = right[j]
			j++
		}
		k++
	}

	if i < lenLeft {
		newU[k] = left[i]
		k++
		i++
	}
	if j < lenRight {
		newU[k] = left[j]
		k++
		j++
	}

	return newU
}
