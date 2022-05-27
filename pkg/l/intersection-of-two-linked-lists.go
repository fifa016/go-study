/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-05-27 21:56:31
 */
package l

func getIntersectionNode(headA, headB *ListNode) *ListNode {

	a := headA
	b := headB
	countA := 0
	countB := 0

	for a != nil {
		countA++
		a = a.Next
	}
	for b != nil {
		countB++
		b = b.Next
	}

	for countA > countB {
		headA = headA.Next
		countA--
	}
	for countA < countB {
		headB = headB.Next
		countB--
	}
	for countA > 0 {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
		countA--
	}

	return nil
}
