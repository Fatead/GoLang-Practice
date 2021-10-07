package rightSideView

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func rightSideView(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	Queue := make([]*TreeNode, 0)
	Queue = append(Queue, root)
	currNum := 1
	nextNum := 0
	for len(Queue) > 0 {
		node := Queue[0]
		Queue = Queue[1:]
		if node.Left != nil {
			Queue = append(Queue, node.Left)
			nextNum++
		}
		if node.Right != nil {
			Queue = append(Queue, node.Right)
			nextNum++
		}
		currNum--
		if currNum == 0 {
			result = append(result, node.Val)
			currNum = nextNum
			nextNum = 0
		}
	}
	return result
}
