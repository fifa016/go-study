/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-17 21:40:57
 */
package l

func detectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}

	slow = head
	for slow != fast {
		fast = fast.Next
		slow = slow.Next
	}

	return fast
}
