/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-05-27 16:42:45
 */
package l

import (
	"bytes"
	"strconv"
)

func compressString(S string) string {
	begin := 0

	var buffer bytes.Buffer

	for i := 0; i < len(S); i++ {
		if S[i] != S[begin]{

			buffer.WriteByte(S[begin])
			buffer.WriteString(strconv.Itoa(i - begin))
			begin = i
		}
	}

	buffer.WriteByte(S[begin])
	buffer.WriteString(strconv.Itoa(len(S) - begin))


	res := buffer.String()
	if len(res) > len(S) {
		return S
	}
	return res
}
