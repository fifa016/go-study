/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-05-27 17:12:31
 */
package l

func reverseVowels(s string) string {

	bytes := []byte(s)

	left := 0
	right := len(s)

	for left < right {
		for left < right && left < len(s) && !isYuan(bytes[left]) {
			left++
		}
		for left < right && right >= 0 && !isYuan(bytes[right]) {
			right--
		}
		bytes[left], bytes[right] = bytes[right], bytes[left]
	}

	return string(bytes)
}

func isYuan(char byte) bool {
	return char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u'
}
