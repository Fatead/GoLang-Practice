package binaryTreePath

import "strconv"

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func dfs(node *TreeNode, path string, allPath *[]string) {
	if node != nil {
		path += strconv.Itoa(node.Val)
	} else {
		return
	}
	if node.Left == nil && node.Right == nil {
		*allPath = append(*allPath, path)
		return
	}
	path += "->"
	dfs(node.Left, path, allPath)
	dfs(node.Right, path, allPath)
}

func binaryTreePaths(root *TreeNode) []string {
	result := make([]string, 0)
	dfs(root, "", &result)
	return result
}
