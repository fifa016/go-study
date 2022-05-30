/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-05-30 15:30:36
 */
package l

func getLengthOfLastWord(s string) int {
	res := 0
	trim := true
	for i := len(s) - 1; i >= 0; i-- {
		if s[i:i+1] == " " && trim {
			continue
		} else if s[i:i+1] != " " {
			trim = false
			res++
		} else {
			break
		}
	}

	return res

}
