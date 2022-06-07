/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-08 01:34:58
 */
package l

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for index := range dp {
		dp[index] = make([]int, len(grid[0]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if i == 0 {
				dp[i][j] = grid[i][j] + dp[i][j-1]
			} else if j == 0 {
				dp[i][j] = grid[i][j] + dp[i-1][j]
			} else {
				dp[i][j] = min(grid[i][j]+dp[i-1][j], grid[i][j]+dp[i][j-1])
			}
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
