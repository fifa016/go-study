/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-08 02:23:59
 */
package l

import "math"

func jump(nums []int) int {
	dp := make([]int, len(nums))

	for i := 1; i < len(nums); i++ {
		min := math.MaxInt
		for j := 0; j < i; j++ {
			canJump := nums[j]+j >= i
			if dp[j] >= 0 && canJump && dp[j]+1 < min {
				min = dp[j] + 1
			}
		}
		dp[i] = min
	}

	return dp[len(dp)-1]
}
