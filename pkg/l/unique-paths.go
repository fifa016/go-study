/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-07 23:30:57
 */
package l

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for index := range dp {
		dp[index] = make([]int, n)
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 || j == n-1 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i+1][j] + dp[i][j+1]
			}
		}
	}
	return dp[0][0]
}
