package 二叉树

import (
	"bufio"
	"os"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isCompleted(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 检测原理 : 按照BFS, 找到第一个不饱和的节点
	var queue []*TreeNode
	queue = append(queue, root)
	flag := false
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		// 从这之后就是不能有孩子节点
		if flag {
			if node.Left != nil || node.Right != nil {
				return false
			}
		}
		if node.Left != nil && node.Right != nil {
			// 1. 如果左右孩子节点都有那么就都保存
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}else if node.Left != nil {
			// 2. 左孩子没有右孩子第一个不饱和的节点
			queue = append(queue, node.Left)
			flag = true
		}else if node.Right != nil {
			// 3. 只有右没有左
			return false
		}else {
			flag = true
		}
	}
	input := bufio.NewScanner(os.Stdin)
	input.Bytes()
	return true
}
