package other

// 转圈打印矩阵
// 解决方案：从外到内按层转圈打印
func PrintMatrix(m [][]int) []int {
	if len(m) == 0 || len(m[len(m)-1]) == 0 {
		return []int{}
	}
	ans := []int{}
	tC, tR, dC, dR := 0, 0, len(m[len(m)-1])-1, len(m)-1
	for tC <= dC && tR <= dR {
		ans = append(ans, printEdge(m, tR, tC, dR, dC)...)
		tC, tR, dC, dR = tC+1, tR+1, dC-1, dR-1
	}
	return ans
}

func printEdge(m [][]int, tR, tC, dR, dC int) []int {
	ans := []int{}
	if tR == dR {
		for idx := tC; idx <= dC; idx++ {
			ans = append(ans, m[tR][idx])
		}
	} else if tC == dC {
		for idx := tR; idx <= dR; idx++ {
			ans = append(ans, m[idx][tC])
		}
	} else {
		// tR
		for idx := tC; idx <= dC; idx++ {
			ans = append(ans, m[tR][idx])
		}
		// dC
		for idx := tR + 1; idx <= dR-1; idx++ {
			ans = append(ans, m[idx][dC])
		}
		// dR
		for idx := dC; idx >= tC; idx-- {
			ans = append(ans, m[dR][idx])
		}
		// tC
		for idx := dR - 1; idx >= tR+1; idx-- {
			ans = append(ans, m[idx][tC])
		}
	}
	return ans
}

// 旋转矩阵
// 解决方案：从外到内按层旋转
func RotateMatrix(m [][]int) {
	if len(m) == 0 || len(m[len(m)-1]) == 0 || len(m) != len(m[len(m)-1]) {
		return
	}
	tR, tC, dR, dC := 0, 0, len(m)-1, len(m[len(m)-1])-1
	for tR <= dR && tC <= dC {
		rotateEdge(m, tR, tC, dR, dC)
		tR, tC, dR, dC = tR+1, tC+1, dR-1, dC-1
	}
}

func rotateEdge(m [][]int, tR, tC, dR, dC int) {
	for idx := 0; idx < dC-tC; idx++ {
		m[tR][tC+idx], m[tR+idx][dC], m[dR][dC-idx], m[dR-idx][tC] =
			m[dR-idx][tC], m[tR][tC+idx], m[tR+idx][dC], m[dR][dC-idx]
	}
}
