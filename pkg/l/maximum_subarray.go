/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-05-27 20:53:39
 */
package l

import "math"

func maxSubArrayWithDP(nums []int) int {
	dp := make([]int, len(nums))
	maxval := nums[0]
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = max(dp[i-1]+nums[i], nums[i])

		}
		if dp[i] > maxval {
			maxval = dp[i]
		}
	}
	return maxval
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

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
