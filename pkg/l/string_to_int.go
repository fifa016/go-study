/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-02 16:20:49
 */

package l

import "math"

func StrToInt(str string) int {
	sign := 1
	var res int64 = 0
	state := "start"

	for i := 0; i < len(str); i++ {
		state = getState(state, str[i])
		if state == "end" {
			break
		}
		if state == "sign" && str[i] == '-' {
			sign = -1
		}
		if state == "num" {
			res *= 10
			res += int64(str[i] - '0')
			if sign == 1 && res > math.MaxInt32 {
				res = math.MaxInt32
			}
			if sign == -1 &&  -res < math.MinInt32 {
				res = -math.MinInt32
			}
		}
	}

	return (int)(res * int64(sign))
}

var stateMap map[string][]string = map[string][]string{
	"start": []string{"start", "sign", "num"},
	"sign":  []string{"end", "end", "num"},
	"num":   []string{"end", "end", "num"},
	"end":   []string{"end", "end", "end"},
}

func getState(state string, b byte) string {
	if b == ' ' {
		return stateMap[state][0]
	}
	if b == '+' || b == '-' {
		return stateMap[state][1]
	}
	if b >= '0' && b <= '9' {
		return stateMap[state][2]
	}
	return "end"
}
