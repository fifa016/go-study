/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-17 20:17:14
 */
package l

func mergeKLists(lists []*ListNode) *ListNode {
	beforeHead := &ListNode{}

	k := []*ListNode{}

	last := -1

	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			k = append(k, lists[i])
			last++
		}
	}

	//建堆
	for i := last; i >= 0; i-- {
		if 2*i+1 < len(k) || 2*i+2 < len(k) {
			adjust3(k, i, last)
		}
	}

	cur := beforeHead

	for len(k) > 0 && k[0] != nil {
		cur.Next = &ListNode{
			Val: k[0].Val,
		}
		cur = cur.Next

		k[0] = k[0].Next
		if k[0] == nil {
			k[0] = k[last]
			last--
		}
		adjust3(k, 0, last)
	}

	return beforeHead.Next
}

func adjust3(lists []*ListNode, current int, last int) {
	if current >= last {
		return
	}

	min := current
	if 2*current+1 <= last && lists[2*current+1].Val < lists[min].Val {
		min = 2*current + 1
	}
	if 2*current+2 <= last && lists[2*current+2].Val < lists[min].Val {
		min = 2*current + 2
	}
	if min == current {
		return
	}
	lists[min], lists[current] = lists[current], lists[min]
	adjust3(lists, min, last)
}
