package alg

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

// 前序遍历
func preorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.val)
	preorderTraversal(root.left)
	preorderTraversal(root.right)
}

// 中序遍历
func inorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	inorderTraversal(root.left)
	fmt.Println(root.val)
	inorderTraversal(root.right)
}

// 后序遍历
func postorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	postorderTraversal(root.left)
	postorderTraversal(root.right)
	fmt.Println(root.val)
}

// 层序遍历
func levelOrder(root *TreeNode) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		fmt.Println(node.val)

		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
}

// 奇数层从左到右遍历，偶数层从右到左遍历
func levelOrderReverse(root *TreeNode) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}
	level := 1
	result := make([][]int, 0)

	for len(queue) > 0 {
		size := len(queue)
		levelValues := make([]int, size)

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			if level%2 == 1 {
				levelValues[i] = node.val
			} else {
				levelValues[size-1-i] = node.val
			}

			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}
		result = append(result, levelValues)

		level++
	}
	fmt.Println(result)
}

// 计算二叉树两节点的最长距离
func maxDistance(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxLen := 0
	dfsDistance(root, &maxLen)
	return maxLen
}

func dfsDistance(node *TreeNode, maxLen *int) int {
	if node == nil {
		return 0
	}

	lHeight := dfsDistance(node.left, maxLen)
	rHeight := dfsDistance(node.right, maxLen)

	if lHeight+node.val+rHeight > *maxLen {
		*maxLen = lHeight + node.val + rHeight
	}
	if lHeight+node.val > *maxLen {
		*maxLen = lHeight + node.val
	}
	if node.val+rHeight > *maxLen {
		*maxLen = node.val + rHeight
	}
	if node.val > *maxLen {
		*maxLen = node.val
	}
	return max(node.val, max(node.val+lHeight, node.val+rHeight))
}

func TestTreeNode(t *testing.T) {
	// 构建一个二叉树示例
	root := &TreeNode{val: 1}
	root.left = &TreeNode{val: 2}
	root.right = &TreeNode{val: 3}
	root.left.left = &TreeNode{val: 4}
	root.left.right = &TreeNode{val: 5}
	root.right.left = &TreeNode{val: 6}
	root.right.right = &TreeNode{val: 7}
	root.left.left.left = &TreeNode{val: 8}
	root.left.left.right = &TreeNode{val: 9}
	root.left.right.left = &TreeNode{val: 10}
	root.left.right.right = &TreeNode{val: 11}
	root.right.left.left = &TreeNode{val: 12}
	root.right.left.right = &TreeNode{val: 13}
	root.right.right.left = &TreeNode{val: 14}
	root.right.right.right = &TreeNode{val: 15}

	fmt.Println("前序遍历结果：")
	preorderTraversal(root)

	fmt.Println("中序遍历结果：")
	inorderTraversal(root)

	fmt.Println("后序遍历结果：")
	postorderTraversal(root)

	fmt.Println("层序遍历结果：")
	levelOrder(root)

	fmt.Println("层序遍历每层翻转结果：")
	levelOrderReverse(root)

	maxDistance(root)
}
