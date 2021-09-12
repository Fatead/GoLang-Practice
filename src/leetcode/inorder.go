package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorder(node *TreeNode, nodeList *list.List) {
	if node == nil {
		return
	}
	inorder(node.Left, nodeList)
	nodeList.PushBack(node.Val)
	inorder(node.Right, nodeList)
}

func inorderTraversal(root *TreeNode) []int {
	nodeList := list.New()
	inorder(root, nodeList)
	nums := make([]int, nodeList.Len())
	index := nodeList.Front()
	for i := 0; i < len(nums); i++ {
		nums[i] = index.Value.(int)
		index = index.Next()
	}
	return nums
}

func main() {

}
