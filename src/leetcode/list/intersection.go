package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pt1 := headA
	pt2 := headB
	stack1 := make([]*ListNode, 0)
	stack2 := make([]*ListNode, 0)
	for pt1 != nil {
		stack1 = append(stack1, pt1)
		pt1 = pt1.Next
	}
	for pt2 != nil {
		stack2 = append(stack2, pt2)
		pt2 = pt2.Next
	}
	index1 := len(stack1) - 1
	index2 := len(stack2) - 1
	var pre *ListNode
	for index1 >= 0 && index2 >= 0 {
		if stack1[index1] == stack2[index2] {
			pre = stack1[index1]
			index1--
			index2--
		} else {
			return pre
		}
	}
	return nil
}
