/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-01 19:20:59
 */
package l

func FindKthLargest(nums []int, k int) int {
	// return minHeap(nums, k)
	return quickSelect(nums, k)
}

func quickSelect(nums []int, k int) int {
	left := 0
	right := len(nums) - 1

	for {
		l := left
		r := right

		for l < r {
			for l < r && nums[r] >= nums[left] {
				r--
			}
			for l < r && nums[l] <= nums[left] {
				l++
			}
			if l < r {
				nums[l], nums[r] = nums[r], nums[l]
			}
		}
		nums[l], nums[left] = nums[left], nums[l]
		if l == len(nums)-k {
			return nums[l]
		} else if l < len(nums)-k {
			left = l + 1
		} else {
			right = l - 1
		}
	}
}
func minHeap(nums []int, k int) int {

	// build heap
	for i := k - 1; i >= 0; i-- {
		if 2*i+1 >= k && 2*i+2 >= k {
			continue
		}
		adjust2(nums, i, k-1)
	}

	//adjust
	for i := k; i < len(nums); i++ {
		if nums[i] > nums[0] {
			nums[i], nums[0] = nums[0], nums[i]
			adjust(nums, 0, k-1)
		}
	}

	return nums[0]
}

func adjust2(nums []int, target, last int) {
	if target >= last {
		return
	}

	minIndex := target
	if 2*target+1 <= last && nums[2*target+1] < nums[minIndex] {
		minIndex = 2*target + 1
	}
	if 2*target+2 <= last && nums[2*target+2] < nums[minIndex] {
		minIndex = 2*target + 2
	}
	if minIndex != target {
		nums[minIndex], nums[target] = nums[target], nums[minIndex]
		adjust2(nums, minIndex, last)
	}

}
