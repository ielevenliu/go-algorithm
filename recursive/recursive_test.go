package recursive

import "testing"

func TestGetMax(t *testing.T) {
	arr := []int{4, 3, 2, 1}
	max := GetMax(arr, 0, len(arr)-1)
	t.Logf("max: %d", max)
}
