/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-15 09:30:53
 */
package l

func ZigzagLevelOrder(root *TreeNode) [][]int {

	res := [][]int{}
	if root == nil {
		return res
	}

	list := []*TreeNode{}

	left := 0
	right := 0
	list = append(list, root)
	isReverse := false
	valList := []int{}

	for left <= right {
		valList = append(valList, list[left].Val)

		if list[left].Left != nil {
			list = append(list, list[left].Left)
		}
		if list[left].Right != nil {
			list = append(list, list[left].Right)
		}

		if left == right {

			right = len(list) - 1

			if !isReverse {
				isReverse = true
			} else {
				i := 0
				j := len(valList) - 1
				for i < j {
					valList[i], valList[j] = valList[j], valList[i]
					i++
					j--
				}
				isReverse = false
			}
			res = append(res, valList)
			valList = []int{}
		}
		left++
	}

	return res

}
