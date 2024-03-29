package l

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		return &ListNode{
			Val:  list1.Val,
			Next: mergeTwoLists(list1.Next, list2),
		}
	} else {
		return &ListNode{
			Val:  list2.Val,
			Next: mergeTwoLists(list2.Next, list1),
		}
	}

}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	head := &ListNode{}
	cur := head

	for list1 != nil && list2 != nil {
		var tmp *ListNode
		if list1.Val < list2.Val {
			tmp = &ListNode{
				Val: list1.Val,
			}
			list1 = list1.Next
		} else {
			tmp = &ListNode{
				Val: list2.Val,
			}
			list2 = list2.Next
		}
		cur.Next = tmp
		cur = cur.Next
	}

	if list1 != nil {
		cur.Next = list1
	}

	if list2 != nil {
		cur.Next = list2
	}

	return head.Next
}
