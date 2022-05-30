/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-05-30 15:30:04
 */

package l

func MergeSort(nums []int) {
	doMergeSort(nums, 0, len(nums)-1)
}

func doMergeSort(nums []int, left, right int) {
	if left < right {
		mid := left + (right-left)/2
		doMergeSort(nums, left, mid)
		doMergeSort(nums, mid+1, right)
		mergeInPlace(nums, left, mid, right)
	}
}

func mergeInPlace(nums []int, left, leftEnd, right int) {
	i := left
	j := leftEnd + 1
	for i < j && j <= right {
		for i < j && nums[i] <= nums[j] {
			i++
		}
		index := j
		for j <= right && nums[j] <= nums[i] {
			j++
		}
		swap2(nums, i, index-1)
		swap2(nums, index, j-1)
		swap2(nums, i, j-1)
		i = i + j - index
	}
}

func swap2(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left ++
		right --
	}
}

func BubbleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		swapped := false
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				swapped = true
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
		if !swapped {
			return
		}
	}
}

func SelectSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		minIndex := i
		for j := i; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
}

func InsertSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		tmp := nums[i]
		j := i - 1
		for ; j >= 0; j-- {
			if nums[j] > tmp {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		if j != i-1 {
			nums[j+1] = tmp
		}
	}
}

func ShellSort(nums []int) {
	for gap := len(nums) / 2; gap > 0; gap = gap / 2 {
		for start := 0; start < gap; start++ {
			for i := start + gap; i < len(nums); i = i + gap {
				target := nums[i]
				j := i - gap
				for ; j >= 0; j = j - gap {
					if nums[j] > target {
						nums[j+gap] = nums[j]
					} else {
						break
					}
				}
				if j != i-gap {
					nums[j+gap] = target
				}

			}

		}
	}

}

//因为处理第i个元素的时候，前i-1个元素都已经排好序了，所以在查找插入位置的时候有优化空间
func InsertSortWithBinarySearch(nums []int) {
	for i := 1; i < len(nums); i++ {

		left := 0
		right := i - 1
		for left <= right {
			mid := left + (right-left)/2
			if nums[mid] == nums[i] {
				left = mid + 1
			} else if nums[mid] < nums[i] {
				left = mid + 1
			} else if nums[mid] > nums[i] {
				right = mid - 1
			}
		}

		tmp := nums[i]
		for j := i - 1; j >= left; j-- {
			nums[j], nums[j+1] = nums[j+1], nums[j]
		}
		nums[left] = tmp

	}
}
