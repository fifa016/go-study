/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-05-27 15:12:58
 */
package l

import "strings"

func repeatedSubstringPattern(s string) bool {
	s2 := s + s
	s2 = s2[1 : len(s2)-1]

	return strings.Contains(s2, s)
}
