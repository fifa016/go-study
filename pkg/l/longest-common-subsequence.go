/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-09 14:45:11
 */
package l

func longestCommonSubsequence(text1 string, text2 string) int {

	dp := make([][]int, len(text1)+1)
	for i := 0; i <= len(text1); i++ {
		dp[i] = make([]int, len(text2)+1)
	}

	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	return dp[len(text1)][len(text2)]
}

func LongestCommonSubstring(text1 string, text2 string) int {

	dp := make([][]int, len(text1)+1)
	for i := 0; i <= len(text1); i++ {
		dp[i] = make([]int, len(text2)+1)
	}

	maxVal := 0
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				maxVal = max(dp[i][j], maxVal)
			}
		}
	}

	return maxVal
}
