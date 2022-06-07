package l

func canJump(nums []int) bool {
	dp := make([]bool, len(nums))

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			dp[i] = true
		} else {
			for j := 0; j < i; j++ {
				if dp[j] && nums[j]+j >= i {
					dp[i] = true
					break
				}
			}
		}
	}

	return dp[len(dp)-1]
}
