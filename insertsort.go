package sort

import "golang.org/x/exp/constraints"

func InsertSortInt(items []int) {
	for i := 1; i < len(items); i++ {
		// loop through subslices
		for j := i; j > 0; j-- {
			if items[j] < items[j-1] {
				// swap elements if not ordered
				items[j], items[j-1] = items[j-1], items[j]
			}
		}
	}
}

// Does not compile, the type constraint any is too broad, the items need to be comparable with <
/*func InsertSortAny[T any](items []T) {
	for i := 1; i < len(items); i++ {
		// loop through subslices
		for j := i; j > 0; j-- {
			if items[j] < items[j-1] {
				// swap elements if not ordered
				items[j], items[j-1] = items[j-1], items[j]
			}
		}
	}
}*/

func InsertSort[T constraints.Ordered](items []T) {
	for i := 1; i < len(items); i++ {
		// loop through subslices
		for j := i; j > 0; j-- {
			if items[j] < items[j-1] {
				// swap elements if not ordered
				items[j], items[j-1] = items[j-1], items[j]
			}
		}
	}
}
