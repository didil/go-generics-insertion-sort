package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertInt(t *testing.T) {
	items := []int{6, 55, 4, 9, 8, 7, 67, 51}

	InsertSortInt(items)

	assert.Equal(t, items, []int{4, 6, 7, 8, 9, 51, 55, 67})
}

func TestInsertSortGenericInt(t *testing.T) {
	items := []int{3, 55, 4, 9, 8, 7, 67, 51}

	InsertSort(items)

	assert.Equal(t, items, []int{3, 4, 7, 8, 9, 51, 55, 67})
}

func TestInsertSortGenericString(t *testing.T) {
	items := []string{"mike", "john", "adil", "aaron", "julia"}

	InsertSort(items)

	assert.Equal(t, items, []string{"aaron", "adil", "john", "julia", "mike"})
}
