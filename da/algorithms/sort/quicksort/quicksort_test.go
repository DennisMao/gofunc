package quicksort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testItem       = []string{"A", "A", "A", "B", "C", "D", "E", "F", "G"}
	testRandomItem = []string{"D", "B", "C", "A", "F", "G", "A", "E", "A"}
)

func NewRandomVector() []string {
	return testRandomItem
}

// 快排
func TestSortQuick(t *testing.T) {

	vector := NewRandomVector()

	assert.Equal(t, testRandomItem, vector)
	vector.SortQuick()
	assert.Equal(t, testItem, vector)

}

// 原地快排
func TestSortQuickInplace(t *testing.T) {

	vector := NewRandomVector()

	assert.Equal(t, testRandomItem, vector)
	vector.SortQuickInplace()
	assert.Equal(t, testItem, vector)

}
