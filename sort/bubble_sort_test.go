package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	arr := []int{4, 3, 5, 2}
	t.Logf("Input: %+v", arr)
	BubbleSort(arr)
	t.Logf("Output: %+v", arr)
}
