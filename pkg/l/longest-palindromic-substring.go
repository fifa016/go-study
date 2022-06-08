/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-08 17:44:49
 */
package l

func LongestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s))
	}

	maxI := 0
	maxJ := 0

	for i := 0; i < len(s); i++ {
		dp[i][i] = true
		if i < len(s)-1 {
			if s[i] == s[i+1] {
				dp[i][i+1] = true
				maxI = i
				maxJ = i + 1
			}
		}
	}

	for length := 3; length <= len(s); length++ {
		for i := 0; i < len(s); i++ {
			j := i + length - 1
			if j >= len(s) {
				break
			}
			if j-i == 2 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
			}
			if dp[i][j] && j-i > maxJ-maxI {
				maxI = i
				maxJ = j
			}
		}
	}

	return s[maxI : maxJ+1]
}
