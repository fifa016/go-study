/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-05-30 18:38:37
 */
package l

// func main() {
//     nums := []int{3, 2, 7, 1, 4, 6, 8, 5, 10, 9}
//     fmt.Println("before sort")
//     fmt.Println(nums)
//     mergeSort(nums)
//     fmt.Println("after sort")
//     fmt.Println(nums)
// }

func mergeSort2(nums []int) {
	doMergeSort2(nums, 0, len(nums)-1)
}

func doMergeSort2(nums []int, start int, end int) {
	if start >= end {
		return
	}
	middle := start + (end-start)/2
	doMergeSort2(nums, start, middle)
	doMergeSort2(nums, middle+1, end)
	merge2(nums, start, middle, middle+1, end)
}

func merge2(nums []int, leftStart int, leftEnd int, rightStart int, rightEnd int) {
	tmpSlice := make([]int, len(nums))
	tmpLeft := leftStart
	tmpRight := rightStart
	index := 0
	for tmpLeft <= leftEnd && tmpRight <= rightEnd {
		cur := tmpLeft
		if nums[tmpLeft] > nums[tmpRight] {
			cur = tmpRight
			tmpRight++
		} else {
			tmpLeft++
		}
		tmpSlice[index] = nums[cur]
		index++
	}
	copySlice(nums, tmpLeft, leftEnd, tmpSlice, index)
	copySlice(nums, tmpRight, rightEnd, tmpSlice, index)
	index = 0
	for leftStart <= rightEnd {
		nums[leftStart] = tmpSlice[index]
		leftStart++
		index++
	}
}

func copySlice(from []int, start int, end int, to []int, toStart int) {
	for start <= end {
		to[toStart] = from[start]
		start++
		toStart++
	}
}
