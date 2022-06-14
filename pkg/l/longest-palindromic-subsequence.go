/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-08 17:44:49
 */
package l

//子序列是不连续的
func longestPalindromeSubseq(s string) int {
	dp := make([][]int, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(s))
	}

	for i := 0; i < len(s); i++ {
		dp[i][i] = 1
		if i < len(s)-1 {
			if s[i] == s[i+1] {
				dp[i][i+1] = 2
			}
		}
	}

	for length := 3; length <= len(s); length++ {
		for i := 0; i < len(s); i++ {
			j := i + length - 1
			if j >= len(s) {
				break
			}
			if s[i] == s[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(s)][len(s)]
}
