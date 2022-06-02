/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-02 16:59:14
 */
package l

func levelOrder(root *NaryTreeNode) [][]int {
	res := [][]int{}
	left := 0
	cache := []*NaryTreeNode{}
	cache = append(cache, root)
	right := 1
	for left < right {
		level := []int{}
		for left < right {
			level = append(level, cache[left].Val)
			if cache[left].children != nil {
				for _, node := range cache[left].children {
					cache = append(cache, node)
				}
			}

			left++

		}

		right = len(cache)
		res = append(res, level)
	}

	return res
}
