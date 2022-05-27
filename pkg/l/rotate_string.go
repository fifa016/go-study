/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-05-27 17:24:01
 */

package l

import "strings"

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	s2 := s + s

	return strings.Contains(s2, goal)
}
