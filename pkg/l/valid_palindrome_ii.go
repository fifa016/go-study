/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-05-27 16:18:47
 */

package l

func validPalindrome(s string) bool {

	bytes := []byte(s)

	begin := 0
	end := len(s) - 1

	for begin < end {
		if bytes[begin] != bytes[end] {
			return isPalin(bytes, begin+1, end) || isPalin(bytes, begin, end-1)
		}
		begin++
		end--
	}

	return true

}

func isPalin(charArray []byte, begin, end int) bool {
	for begin < end {
		if charArray[begin] != charArray[end] {
			return false
		}
		begin++
		end--
	}
	return true
}
