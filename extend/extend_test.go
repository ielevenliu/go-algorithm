package extend

import "testing"

func TestSimplePartition(t *testing.T) {
	arr := []int{6, 5, 4, 3, 2, 1}
	t.Logf("Original arr: %+v", arr)

	SimplePartition(arr, 0, len(arr)-1, 3)
	t.Logf("Simple partition arr: %+v", arr)
}

func TestHollandFlag(t *testing.T) {
	arr := []int{6, 5, 4, 3, 2, 1}
	t.Logf("Original arr: %+v", arr)

	HollandFlag(arr, 0, len(arr)-1, 3)
	t.Logf("Holland flag arr: %+v", arr)
}
