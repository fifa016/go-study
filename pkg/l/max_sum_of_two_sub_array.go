/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-08 10:52:14
 */
package l

func MaxSumOfTwoSubArray(nums []int) int {
	dp := make([]int, len(nums))
	maxval := nums[len(nums)-1]
	sumToRight := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			sumToRight = nums[i]
		} else {
			sumToRight = max(nums[i]+sumToRight, nums[i])
		}
		if sumToRight > maxval {
			maxval = sumToRight
		}
		dp[i] = maxval
	}

	maxval = nums[0]
	for i := 0; i < len(nums)-1; i++ {
		if i == 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = max(dp[i-1]+nums[i], nums[i])
		}
		maxval = max(dp[i]+dp[i+1], maxval)
	}

	return maxval
}
