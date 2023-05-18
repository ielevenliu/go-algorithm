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

/*
归并排序的扩展
Question：
 1. 小和问题，在一个数组中，每一个数左边比当前数小的数累加起来，叫做这个数组的小和
    例子：[1, 3, 4, 2, 5]
         1左边比1小的数，没有
         3左边比3小的数，1
         4左边比4小的数，1,3
         2左边比2小的数，1
         5左边比5小的数，1,3,4,2
        小和为 1+(1+3)+1+(1+3+4+2)
 2. 逆序对问题，在一个数组中，左边的数如果比右边的数大，则这两个数构成一个逆序对
*/
