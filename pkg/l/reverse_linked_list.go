/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-05-27 21:57:12
 */
package l

func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode = nil

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}
