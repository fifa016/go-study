/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-05-27 17:19:25
 */
package l

func reverseLeftWords(s string, n int) string {
	bytes := []byte(s)
	leftPart := bytes[0:n]
	rightPart := bytes[n:]
	bytes = append(rightPart, leftPart...)

	return string(bytes)

}
