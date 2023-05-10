package sort

import "testing"

func TestBasicSelectionSort(t *testing.T) {
	arr := []int{2, 4, 1, 5}
	t.Logf("Input: %+v", arr)
	BasicSelectionSort(arr)
	t.Logf("Output: %+v", arr)
}

func BenchmarkBasicSelectionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []int{2, 5, 1, 7, 6, 3}
		BasicSelectionSort(arr)
	}
}

func TestOptimizeSelectionSort(t *testing.T) {
	arr := []int{2, 4, 1, 5}
	t.Logf("Input: %+v", arr)
	OptimizeSelectionSort(arr)
	t.Logf("Output: %+v", arr)
}
