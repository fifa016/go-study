/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-01 16:49:38
 */
package l

//一个数组，有正有负，把正的移到右边，负的移到左边。
func Move(nums []int) {

	for i, count := 0, 0; i < len(nums) && count < len(nums); i++ {
		if nums[i] == 0 {
			continue
		} else if nums[i] < 0 {
			tmp := nums[i]
			for j := i; j >= 1; j-- {
				nums[j] = nums[j-1]
			}
			nums[0] = tmp
		} else {
			tmp := nums[i]
			for j := i; j < len(nums)-1; j++ {
				nums[j] = nums[j+1]
			}
			nums[len(nums)-1] = tmp
			i--
		}
		count++
	}

}
