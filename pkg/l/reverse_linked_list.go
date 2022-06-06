/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-05-27 21:57:12
 */
package l

func reverseKGroupTest20220606(head *ListNode, k int) *ListNode {

	beforeHead := &ListNode{Next: head}

	pre := beforeHead
	start := pre.Next
	end := start

	for end != nil {

		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}

		next := end.Next

		end.Next = nil

		pre.Next = reverseTest20220606(start)
		start.Next = next

		pre = start
		start = next
		end = start

	}

	return beforeHead.Next

}

func reverseTest20220606(start *ListNode) *ListNode {
	if start == nil || start.Next == nil {
		return start
	}

	var pre *ListNode
	cur := start
	for cur != nil {
		next := cur.Next

		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	beforeHead := &ListNode{Next: head}

	pre := beforeHead
	end := pre

	for end != nil {
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}

		if end == nil {
			break
		}

		start := pre.Next
		next := end.Next

		end.Next = nil
		pre.Next = reverseList(start)
		start.Next = next

		pre = start
		end = pre

	}

	return beforeHead.Next
}

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

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil || left == right {
		return head
	}

	res := &ListNode{
		Next: head,
	}
	cur := head
	var beforeLeft *ListNode = res
	var afterRight *ListNode
	for i := 0; ; i++ {
		if i+1 == left-1 {
			beforeLeft = cur
		}
		if i == right {
			afterRight = cur
			break
		}
		cur = cur.Next
	}

	var pre *ListNode = afterRight
	cur = beforeLeft.Next
	for i := 0; i <= right-left; i++ {
		next := cur.Next

		cur.Next = pre

		pre = cur
		cur = next

	}

	beforeLeft.Next = pre

	return res.Next

}
