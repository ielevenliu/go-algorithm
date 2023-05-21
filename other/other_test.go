package other

import "testing"

func TestPrintMatrix(t *testing.T) {
	m := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	ans := PrintMatrix(m)
	t.Logf("matrix: %+v", ans)
}
