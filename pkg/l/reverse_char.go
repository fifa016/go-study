/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-05-26 16:14:11
 */
package l

func minFlipsMonoIncr(s string) int {
	firstOneFromLeft := -1
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			firstOneFromLeft = i
			break
		}
	}
	if firstOneFromLeft == -1 {
		return 0
	}
	toOneCount := 0
	for i := len(s) - 1; i >= firstOneFromLeft; i-- {
		if s[i] == '0' {
			toOneCount ++
		}
	}


	firstZeroFromRight := len(s)
	toZeroCount := 0
	for i := len(s) - 1; i >=0 ; i-- {
		if s[i] == '0'{
			firstZeroFromRight = i
			break
		}
	}
	if firstZeroFromRight == len(s) {
		return 0
	}

	for i := 0; i <= firstZeroFromRight; i++ {
		if s[i] == '1' {
			toZeroCount ++
		}
	}

	if toZeroCount < toOneCount {
		return toZeroCount
	}else {
		return toOneCount
	}


}
