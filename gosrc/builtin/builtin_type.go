package main

import (
	"fmt"
	"unsafe"
)

////////////////// Slice ////////////////////
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

func PrintSlice() {
	k := new([]string)
	fmt.Printf("===newSlice          \nPointer:%#v \nObject:%#v \nDataStructure:%#v\n", k, *k, (*slice)(unsafe.Pointer(k)))

	*k = append(*k, "1")
	fmt.Printf("===newSliceAppended  \nPointer:%#v \nObject:%#v \nDataStructure:%#v\n", k, *k, (*slice)(unsafe.Pointer(k)))

	c := make([]string, 0)
	fmt.Printf("===makeSlice         \nPointer:%#v \nObject:%#v \nDataStructure:%#v\n", &c, c, (*slice)(unsafe.Pointer(&c)))

	c = append(c, "1")
	fmt.Printf("===makeSliceAppended \nPointer:%#v \nObject:%#v \nDataStructure:%#v\n", &c, c, (*slice)(unsafe.Pointer(&c)))
}

////////////////// Map ///////////////////////
const (
	bucketCnt = 8
)

// A header for a Go map.
type hmap struct {
	count      int // # live cells == size of map.  Must be first (used by len() builtin)
	flags      uint8
	B          uint8          // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
	hash0      uint32         // hash seed
	buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
	oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
	nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)
	overflow   *[2]*[]*bmap
}

// A bucket for a Go map.
type bmap struct {
	tophash [bucketCnt]uint8
}

func PrintMap() {
	k := new(map[string]interface{})
	fmt.Printf("===newMap      \nPointer:%#v\nObject:%#v \nDataStructure:%#v\n", k, *k, (*hmap)(unsafe.Pointer(k)))

	c := make(map[string]interface{}, 0)
	fmt.Printf("===makeMap     \nPointer:%#v \nObject:%#v \nDataStructure:%#v\n", &c, c, (*hmap)(unsafe.Pointer(&c)))
}

////////////////// Channel ///////////////////
type hchan struct {
	qcount   uint           // total data in the queue
	dataqsiz uint           // size of the circular queue
	buf      unsafe.Pointer // points to an array of dataqsiz elements
	elemsize uint16
	closed   uint32
	elemtype unsafe.Pointer
	sendx    uint // send index
	recvx    uint // receive index
	recvq    struct {
		first unsafe.Pointer
		last  unsafe.Pointer
	} // list of recv waiters
	sendq struct {
		first unsafe.Pointer
		last  unsafe.Pointer
	} // list of send waiters
	lock struct {
		key uintptr
	}
}

func PrintChannel() {
	k := new(chan int)
	fmt.Printf("===newChannel      \nPointer:%#v \nObject:%#v \nDataStructure:%#v\n", k, *k, (*hchan)(unsafe.Pointer(k)))

	c := make(chan int)
	fmt.Printf("===makeChannel     \nPointer:%#v \nObject:%#v \nDataStructure:%#v\n", &c, c, (*hchan)(unsafe.Pointer(&c)))
}

func main() {
	PrintSlice()
	PrintMap()
	PrintChannel()

}
