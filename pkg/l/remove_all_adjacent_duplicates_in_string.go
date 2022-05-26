/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-05-26 20:54:11
 */
package l

func removeDuplicates2(s string) string {

	bytes := []byte{}

	for i := 0; i < len(s); i++ {
		if len(bytes) == 0 || s[i] != bytes[len(bytes)-1] {
			bytes = append(bytes, s[i])
		} else {
			bytes = bytes[:len(bytes)-1]
		}
	}

	return string(bytes)
}
