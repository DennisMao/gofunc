package quick

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testItem       = []string{"A", "A", "A", "B", "C", "D", "E", "F", "G"}
	testRandomItem = []string{"D", "B", "C", "A", "F", "G", "A", "E", "A"}
)

func NewRandomItem() []string {
	n := make([]string, len(testRandomItem))
	copy(n, testRandomItem)
	return n
}

func TestSortMerge(t *testing.T) {
	item := NewRandomItem()

	assert.Equal(t, testRandomItem, item)
	SortMerge(&item)
	assert.Equal(t, testItem, item)
}

func TestRuntimeStats(t *testing.T) {

	item := NewRandomItem()
	assert.Equal(t, testRandomItem, item)

	for i := 0; i < 10; i++ {
		runtime.GC()
	}
	fmt.Println("-------- BEFORE ----------")
	printCurMems()

	SortMerge(&item) // Excute sort

	fmt.Println("-------- AFTER ----------")
	printCurMems()

	assert.Equal(t, testItem, item)
}

func printCurMems() {

	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	fmt.Printf("Alloc:%d  TotalAlloc:%d \n", stats.Alloc, stats.TotalAlloc)
}
