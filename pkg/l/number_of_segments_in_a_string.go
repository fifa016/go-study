/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-05-27 14:51:56
 */
package l

func countSegments(s string) int {
	begin := 0
	end := len(s) - 1
	for begin < len(s) && s[begin] == ' ' {
		begin++
	}
	for end >= 0 && s[end] == ' ' {
		end--
	}

    if begin > end {
        return 0
    }

	last := begin 
	res := 1

	for begin <= end {
		if s[begin] == ' ' && begin != last && s[last] != s[begin] {
			res++
		}
		last = begin
		begin++
	}
     
	return res 

}
