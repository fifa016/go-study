/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-08 03:07:34
 */
package l

func maxProfitII(prices []int) int {
	dp := make([]int, len(prices))

	dp[0] = 0
	for i := 1; i < len(prices); i++ {
		max := 0
		for j := 0; j < i; j++ {
			profit := dp[j]
            if prices[i] - prices[j] > 0 {
                profit += prices[i] - prices[j]
            }
			if profit > max {
				max = profit 
			}
		}
		dp[i] = max
	}

	return dp[len(prices)-1]
}
