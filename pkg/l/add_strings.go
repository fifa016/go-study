package l

import "strconv"

func addStrings(num1 string, num2 string) string {
	res := ""
	n1 := len(num1) - 1
	n2 := len(num2) - 1
	jinwei := 0

	for n1 >= 0 || n2 >= 0 || jinwei > 0 {
		val1 := 0
		val2 := 0

		if n1 >= 0 {
			val1 = int(num1[n1] - '0')
		}
		if n2 >= 0 {
			val2 = int(num2[n2] - '0')
		}
		sum := val1 + val2 + jinwei

		if sum >= 10 {
			jinwei = 1
			sum = sum - 10
		} else {
			jinwei = 0
		}
		res = strconv.Itoa(sum) + res
		n1--
		n2--
	}

	return res

}
