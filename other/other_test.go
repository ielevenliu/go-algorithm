package other

import "testing"

func TestPrintMatrix(t *testing.T) {
	m := [][]int{
		{1, 2, 3, 4},
		{4, 5, 6, 6},
		{7, 8, 9, 9},
	}
	t.Logf("Orginal matrix: ")
	for idx := 0; idx < len(m); idx++ {
		t.Logf("%+v", m[idx])
	}

	ans := PrintMatrix(m)
	t.Logf("Print matrix: %+v", ans)
}

func TestRotateMatrix(t *testing.T) {
	m := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	t.Logf("Orginal matrix: ")
	for idx := 0; idx < len(m); idx++ {
		t.Logf("%+v", m[idx])
	}

	RotateMatrix(m)
	t.Logf("Rotate matrix: ")
	for idx := 0; idx < len(m); idx++ {
		t.Logf("%+v", m[idx])
	}
}
