/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-03 10:57:25
 */
package l

func LinkedListMergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head
	fast := head
	var brk *ListNode

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next

		if fast == nil || fast.Next == nil { //就是在最后一步之前确定brk的位置
			brk = slow
		}
		slow = slow.Next
	}
	brk.Next = nil

	head1 := LinkedListMergeSort(head)
	head2 := LinkedListMergeSort(slow)
	
	beforeHead := &ListNode{}
	cur := beforeHead
	for head1 !=  nil || head2 != nil {
		if head1 == nil || (head2 != nil && head1.Val >= head2.Val  ) {
			cur.Next = head2
			head2 = head2.Next
		} else {
			cur.Next = head1
			head1 = head1.Next
		}
		cur = cur.Next
	}

	return beforeHead.Next
}

func LinkedListBubbleSort(head *ListNode) *ListNode {
	count := 0
	cur := head
	for cur != nil {
		count++
		cur = cur.Next
	}

	if count <= 1 {
		return head
	}

	beforeHead := &ListNode{Next: head}
	for i := 0; i < count-1; i++ {
		var pre *ListNode = beforeHead
		cur1 := beforeHead.Next
		cur2 := cur1.Next
		for cur2 != nil {
			for i < count-1 && cur1.Val <= cur2.Val { //注意先判断i越界
				pre = cur1
				cur1 = cur2
				cur2 = cur2.Next
				i++
			}
			if i >= count-1 {
				break
			}
			next := cur2.Next

			if pre != nil {
				pre.Next = cur2
			}
			cur1.Next = next
			cur2.Next = cur1

			pre = cur2
			cur2 = next
			i++ //注意这里别忘了
		}
		count--
	}

	return beforeHead.Next
}

func InsertSortLinkedList(head *ListNode) *ListNode {

	beforeHead := &ListNode{
		Next: head,
	}

	cur := head

	for cur.Next != nil {
		insertPos := beforeHead
		for insertPos.Next != nil && insertPos != cur && insertPos.Next.Val <= cur.Next.Val {
			insertPos = insertPos.Next
		}
		if insertPos == cur {
			cur = cur.Next
			continue
		}

		tmp := cur.Next
		cur.Next = tmp.Next

		tmp.Next = insertPos.Next
		insertPos.Next = tmp

	}

	return beforeHead.Next
}
