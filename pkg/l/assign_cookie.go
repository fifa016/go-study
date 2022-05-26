package l

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	gIndex := 0
	sIndex := 0

	for gIndex < len(g) && sIndex < len(s) {
		if g[gIndex] < s[sIndex] {
			gIndex++
		}
		sIndex++
	}

	return gIndex
}

func IndContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	gIndex := 0
	sIndex := 0
	for gIndex < len(g) && sIndex < len(s) {
		if g[gIndex] <= s[sIndex] {
			gIndex++
		}
		sIndex++
	}
	return gIndex

}
