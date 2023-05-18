package search

import "testing"

func TestTraversalSearch(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}
	b := []int{8, 9, 2, 3}

	ans := TraversalSearch(a, b)
	t.Logf("TraversalSearch ans: %+v", ans)
}

func TestBinarySearch(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}
	b := []int{8, 9, 2, 3}

	ans := BinarySearch(a, b)
	t.Logf("BinarySearch ans: %+v", ans)
}

func TestSortSearch(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}
	b := []int{8, 9, 2, 3}

	ans := SortSearch(a, b)
	t.Logf("SortSearch ans: %+v", ans)
}
