/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-05-31 22:07:34
 */
package l

func getKthFromEnd(head *ListNode, k int) *ListNode {
	count := 0
	cur := head
	for cur != nil {
		count++
		cur = cur.Next
	}

	cur = head 
	for i := 0; i < count - k; i++ {
		cur  = cur.Next
	}
	return cur
}
