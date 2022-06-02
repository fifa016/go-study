/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-02 17:47:54
 */
package l

func search(arr []int, target int) int {
	if arr[0] == target {
		return 0
	}

	left := 0
	right := len(arr) - 1
	mid :=0
	for left <= right {
		mid = left + (right-left)/2
		if arr[mid] == target {
			for mid > 0 && arr[mid-1] == arr[mid] {
				mid--
			}
			return mid
		}

		if arr[mid] < arr[right] { //右侧是升序区
			if arr[mid] < target && target <= arr[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else if arr[right] < arr[mid] { // 左侧是升序区
			if arr[left] <= target && target < arr[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} 
	}

	return -1

}
