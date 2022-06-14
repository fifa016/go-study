/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-14 14:06:12
 */
package l

func minDistance(word1 string, word2 string) int {

	dp := make([][]int, len(word1)+1)

	for i := 0; i < len(word1)+1; i++ {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 0; i < len(word1)+1; i++ {
		for j := 0; j < len(word2)+1; j++ {
			if i == 0 {
				dp[i][j] = j
				continue
			}
			if j == 0 {
				dp[i][j] = i
				continue
			}

			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				//替换
				swap := dp[i-1][j-1] + 1

				//插入left
				insert := dp[i][j-1] + 1

				//删除left
				del := dp[i-1][j] + 1

				dp[i][j] = min(swap, min(insert, del))
			}
		}
	}

	return dp[len(word1)][len(word2)]

}
