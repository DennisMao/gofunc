package quick

import (
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

func TestSortQuick(t *testing.T) {
	item := NewRandomItem()

	assert.Equal(t, testRandomItem, item)
	SortQuick(&item)
	assert.Equal(t, testItem, item)
}
