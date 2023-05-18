package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	arr := []int{7, 3, 6, 5, 1}
	t.Logf("Original arr: %+v", arr)

	BubbleSort(arr)
	t.Logf("BubbleSort arr: %+v", arr)
}

func TestSelectionSort(t *testing.T) {
	arr := []int{7, 3, 6, 5, 1}
	t.Logf("Original arr: %+v", arr)

	SelectionSort(arr)
	t.Logf("SelectSort arr: %+v", arr)
}

func TestInsertionSort(t *testing.T) {
	arr := []int{7, 3, 6, 5, 1}
	t.Logf("Original arr: %+v", arr)

	InsertionSort(arr)
	t.Logf("SelectSort arr: %+v", arr)
}

func TestMergeSort(t *testing.T) {
	arr := []int{7, 3, 6, 5, 1}
	t.Logf("Original arr: %+v", arr)

	MergeSortRecursive(arr)
	t.Logf("MergeSort arr: %+v", arr)
}
