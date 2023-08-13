package search

import "testing"

func TestBinarySearchTarget(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}

	idx := BinarySearchTarget(nums, 9)
	t.Logf("Index: %d", idx)
}

func TestBinarySearchLeftEdge(t *testing.T) {
	nums := []int{-1, 1, 9, 9, 9, 12}

	idx := BinarySearchLeftEdge(nums, 9)
	t.Logf("Index: %d", idx)
}

func TestBinarySearchRightEdge(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}

	idx := BinarySearchRightEdge(nums, 6)
	t.Logf("Index: %d", idx)
}
