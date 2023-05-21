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
		tC++
		dC--
		tR++
		dR--
	}
	for tC <= dC {
		ans = append(ans, printEdge(m, tR, tC, dR, dC)...)
		tC++
		dC--
	}
	for tR <= dR {
		ans = append(ans, printEdge(m, tR, tC, dR, dC)...)
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
