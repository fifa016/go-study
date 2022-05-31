/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-05-30 15:30:04
 */

package l


func CountSort(nums []int){
	max := 0
	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	tmpArr := make([]int , max + 1)

	for i := 0; i < len(nums); i++ {
		tmpArr[nums[i]] ++
	}
	index := 0

	for i := 0; i < len(tmpArr); i++ {
		for tmpArr[i] > 0 {
			nums[index] = i
			tmpArr[i] --
			index ++
		}
	}
}

func HeapSort(nums []int) {

	//build heap
	for i := len(nums) - 1; i >= 0; i-- {
		if 2*i+1 >= len(nums) && 2*i+2 >= len(nums) {
			continue
		}

		adjust(nums, i, len(nums)-1)
	}

	// swap and adjust
	for i := len(nums) - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		adjust(nums, 0, i-1)
	}

}

func adjust(nums []int, target, last int) {

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

	if target != minIndex {
		nums[target], nums[minIndex] = nums[minIndex], nums[target]
		adjust(nums, minIndex, last)
	}

}

func MergeSortNoInPlace(nums []int) {
	doMergeSortNoInPlace(nums, 0, len(nums)-1)
}
func doMergeSortNoInPlace(nums []int, left, right int) {
	if left >= right {
		return
	}
	mid := left + (right-left)/2
	doMergeSortNoInPlace(nums, left, mid)
	doMergeSortNoInPlace(nums, mid+1, right)

	mergeNoInPlace(nums, left, mid, right)
}
func mergeNoInPlace(nums []int, left, mid, right int) {
	tmpNums := make([]int, len(nums))
	tmpIndex := 0
	l := left
	r := mid + 1

	for l <= mid && r <= right {
		if nums[l] < nums[r] {
			tmpNums[tmpIndex] = nums[l]
			l++
		} else {
			tmpNums[tmpIndex] = nums[r]
			r++
		}
		tmpIndex++
	}
	for l <= mid {
		tmpNums[tmpIndex] = nums[l]
		l++
		tmpIndex++
	}
	for r <= right {
		tmpNums[tmpIndex] = nums[r]
		r++
		tmpIndex++
	}
	tmpIndex = 0
	length := right - left + 1
	for tmpIndex < length {
		nums[left] = tmpNums[tmpIndex]
		tmpIndex++
		left++
	}
}

func QuickSort(nums []int) {
	doQuickSort(nums, 0, len(nums)-1)
}

func doQuickSort(nums []int, left, right int) {
	if left >= right {
		return
	}

	flag := left
	l := left
	r := right

	for l < r {
		for l < r && nums[r] > nums[flag] {
			r--
		}
		for l < r && nums[l] < nums[flag] {
			l++
		}

		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
		nums[l] = nums[flag]
	}
	doQuickSort(nums, left, l)
	doQuickSort(nums, l+1, right)

}

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
		left++
		right--
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
