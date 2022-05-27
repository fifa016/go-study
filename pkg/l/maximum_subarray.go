/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-05-27 20:53:39
 */
package l

import "math"

func maxSubArray(nums []int) int {
	maxSum := math.MinInt
	sum := math.MinInt
	for i := 0; i < len(nums); i++ {

		if nums[i] > sum+nums[i] {
			sum = nums[i]
		} else {
			sum = sum + nums[i]
		}

		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum
}
