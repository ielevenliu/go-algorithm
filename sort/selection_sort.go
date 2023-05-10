package sort

// 基础款
func BasicSelectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minPos := i
		for j := i + 1; j < len(arr); j++ {
			if arr[minPos] > arr[j] {
				minPos = j
			}
		}
		if minPos != i {
			arr[minPos], arr[i] = arr[i], arr[minPos]
		}
	}
}

// 优化款(同时找到最大值和最小值)
func OptimizeSelectionSort(arr []int) {
	for h, t := 0, len(arr)-1; t-h > 2; {
		minPos, maxPos := h, t
		for j := h + 1; j < t-1; j++ {
			if arr[minPos] > arr[j] {
				minPos = j
			}
			if arr[maxPos] <= arr[j] {
				maxPos = j
			}
		}
		if minPos != h {
			arr[minPos], arr[h] = arr[h], arr[minPos]
		}
		if maxPos != t {
			arr[maxPos], arr[t] = arr[t], arr[maxPos]
		}
		h++
		t--
	}
}
