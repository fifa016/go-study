/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-14 20:23:19
 */
package l

func lengthOfLongestSubstring(s string) int {

	res := 0
	indexMap := map[byte]int{}

	left := 0
	for i := 0; i < len(s); i++ {
		if index, exist := indexMap[s[i]]; exist {
			left = max(left, index+1)
		}
		indexMap[s[i]] = i

		res = max(res, i-left+1)

	}

	return res
}
