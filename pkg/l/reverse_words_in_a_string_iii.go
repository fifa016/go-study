/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-05-26 16:42:39
 */
package l

func reverseWords(s string) string {

	runes := []rune(s)

	spaceFound := 0

	for spaceFound < len(s) {
		start := spaceFound
		for spaceFound < len(s) && runes[spaceFound] != ' ' {
			spaceFound++
		}
		reverse(runes, start, spaceFound-1)
		spaceFound ++
	}
	return string(runes)
}

func reverse(array []rune, begin int, end int) {

	for begin < end {
		array[begin], array[end] = array[end], array[begin]
		begin++
		end--
	}

}
