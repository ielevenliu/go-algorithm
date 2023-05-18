package sort

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

// MergeSortRecursive: 归并排序时间复杂度T(N) = 2T(N/2) + O(N) => O(N*logN)
func MergeSortRecursive(arr []int) {
	if len(arr) < 2 {
		return
	}
	mergeSortRecursive(arr, 0, len(arr)-1)
}

func mergeSortRecursive(arr []int, l, r int) {
	if l == r {
		return
	}
	// 防止l+r数据过大
	mid := l + ((r - l) >> 2)
	mergeSortRecursive(arr, l, mid)
	mergeSortRecursive(arr, mid+1, r)
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
