package search

import "sort"

/*
Question: 有序数组a长度为N，无序数组b长度为M，找出b中不在a中的元素
*/

// TraversalSearch: 时间复杂度O(N*M)
func TraversalSearch(a, b []int) []int {
	if len(a) == 0 || len(b) == 0 {
		return b
	}

	ans := []int{}
	for _, elemB := range b {
		flag := false
		for _, elemA := range a {
			if elemB == elemA {
				flag = true
				break
			}
		}
		if !flag {
			ans = append(ans, elemB)
		}
	}
	return ans
}

// BinarySearch: 时间复杂度O(N*logM)，以2为底
func BinarySearch(a, b []int) []int {
	if len(a) == 0 || len(b) == 0 {
		return b
	}

	ans := []int{}
	for _, elemB := range b {
		lowIdx := 0
		highIdx := len(a) - 1
		for lowIdx <= highIdx {
			midIdx := (lowIdx + highIdx) / 2
			if a[midIdx] < elemB {
				lowIdx = midIdx + 1
			} else if a[midIdx] > elemB {
				highIdx = midIdx - 1
			} else {
				break
			}
		}
		if lowIdx > highIdx {
			ans = append(ans, elemB)
		}
	}
	return ans
}

// SortSearch: 时间复杂度O(M*logM)+O(N+M)，以2为底
func SortSearch(a, b []int) []int {
	if len(a) == 0 || len(b) == 0 {
		return b
	}

	ans := []int{}
	sort.Ints(b)
	var idxA, idxB int
	for idxA < len(a) && idxB < len(b) {
		if a[idxA] < b[idxB] {
			idxA++
		} else if a[idxA] > b[idxB] {
			ans = append(ans, b[idxB])
			idxB++
		} else {
			idxB++
		}
	}
	for ; idxB < len(b); idxB++ {
		ans = append(ans, b[idxB])
	}
	return ans
}
