package sort

import (
	"math/rand"
	"time"
)

// BubbleSort: 冒泡排序时间复杂度O(N^2)
func BubbleSort(arr []int) {
	for end := len(arr) - 1; end > 0; end-- {
		for idx := 0; idx < end; idx++ {
			if arr[idx] > arr[idx+1] {
				arr[idx], arr[idx+1] = arr[idx+1], arr[idx]
			}
		}
	}
}

// SelectionSort: 选择排序时间复杂度O(N^2)
func SelectionSort(arr []int) {
	for start := 0; start < len(arr)-1; start++ {
		minIdx := start
		for idx := start + 1; idx < len(arr); idx++ {
			if arr[minIdx] > arr[idx] {
				minIdx = idx
			}
		}
		if minIdx != start {
			arr[start], arr[minIdx] = arr[minIdx], arr[start]
		}
	}
}

// InsertionSort: 插入排序时间复杂度 最好情况O(N)、最差情况O(N^2)、平均情况O(N)
func InsertionSort(arr []int) {
	for start := 0; start < len(arr)-1; start++ {
		for idx := start + 1; idx > 0 && arr[idx] < arr[idx-1]; idx-- {
			arr[idx-1], arr[idx] = arr[idx], arr[idx-1]
		}
	}
}

// MergeSortRecursive: 归并排序时间复杂度T(N) = 2T(N/2) + O(N) => O(N*logN); 额外空间复杂度O(N)
func MergeSortRecursive(arr []int) {
	if len(arr) < 2 {
		return
	}
	mergeSortProcess(arr, 0, len(arr)-1)
}

func mergeSortProcess(arr []int, l, r int) {
	if l == r {
		return
	}
	// 防止l+r数据过大溢出
	mid := l + ((r - l) >> 2)
	mergeSortProcess(arr, l, mid)
	mergeSortProcess(arr, mid+1, r)
	merge(arr, l, mid, r)
}

func merge(arr []int, l, mid, r int) {
	tmp := []int{}
	idxL, idxR := l, mid+1
	for idxL <= mid && idxR <= r {
		if arr[idxL] <= arr[idxR] {
			tmp = append(tmp, arr[idxL])
			idxL++
		} else {
			tmp = append(tmp, arr[idxR])
			idxR++
		}
	}
	for idxL <= mid {
		tmp = append(tmp, arr[idxL])
		idxL++
	}
	for idxR <= r {
		tmp = append(tmp, arr[idxR])
		idxR++
	}
	for idx := range tmp {
		arr[l+idx] = tmp[idx]
	}
}

// todo
func MergeSortNonRecursive() {

}

// QuickSort: 随机快排长期期望的时间复杂度T(N) = 2T(N/2) + O(N) => O(N*logN); 额外空间复杂度O(logN)
func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	quickSortProcess(arr, 0, len(arr)-1)
}

func quickSortProcess(arr []int, l, r int) {
	if l >= r {
		return
	}
	// 随机一个索引位置用做partition对比，相较于经典快排直接用最后一个数而言，屏蔽了与数据状况[4,3,2,1]的相关性
	rand.Seed(time.Now().UnixNano())
	swapIdx := rand.Intn(r-l) + l
	arr[swapIdx], arr[r] = arr[r], arr[swapIdx]
	less, more := partition(arr, l, r)
	quickSortProcess(arr, l, less-1)
	quickSortProcess(arr, more+1, r)
}

func partition(arr []int, l, r int) (int, int) {
	less, more := l-1, r
	for l < more {
		if arr[l] < arr[r] {
			arr[less+1], arr[l] = arr[l], arr[less+1]
			less++
		} else if arr[l] > arr[r] {
			arr[more-1], arr[l] = arr[l], arr[more-1]
			more--
			continue
		}
		l++
	}
	arr[r], arr[more] = arr[more], arr[r]
	return less + 1, more
}

// todo
func QuickSortNonRecursive(arr []int) {

}

// HeapSort:
func HeapSort(arr []int) {

}
