package recursive

// GetMax: T(N) = 2T(N/2) + O(1) => O(N)
func GetMax(arr []int, left, right int) int {
	if left == right {
		return arr[left]
	}
	mid := (left + right) / 2
	leftMax := GetMax(arr, left, mid)     // 子过程
	rightMax := GetMax(arr, mid+1, right) // 子过程
	if leftMax > rightMax {
		return leftMax
	}
	return rightMax
}
