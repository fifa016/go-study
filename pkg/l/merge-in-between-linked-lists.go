/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-03 15:07:05
 */
package l

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {

	tail2 := list2
	for tail2.Next != nil {
		tail2 = tail2.Next
	}

	beforeA := list1
	afterB := list1
	cur := list1
	for i := 0; ; i++ {
		if i+1 == a {
			beforeA = cur
		}
		if i == b + 1 {
			afterB = cur
            break
		}
		cur = cur.Next
	}

	beforeA.Next = list2
	tail2.Next = afterB

	return list1
}
