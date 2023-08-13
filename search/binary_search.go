package search

// 寻找一个数
func BinarySearchTarget(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// 寻找左边界
// 循环终止条件是 left = right + 1，即 right = mid-1 后来到了target值的左边left位置为target
func BinarySearchLeftEdge(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			right = mid - 1
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if len(arr) == 0 || arr[left] != target {
		return -1
	}
	return left
}

// 寻找右边界
func BinarySearchRightEdge(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			left = mid + 1
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if len(arr) == 0 || arr[right] != target {
		return -1
	}
	return right
}
