package maxDepth

import (
	"algo/tree/tree"
	"math"
)

/**
给定一个二叉树，找出其最大深度。
二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
说明: 叶子节点是指没有子节点的节点。
示例：
给定二叉树 [3,9,20,null,null,15,7]，
    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。
*/

func MaxDepth(root *tree.Node) int {
	if root == nil {
		return 0
	}
	return maxDepthCompute(root)
}

func maxDepthCompute(root *tree.Node) int {
	if root == nil {
		return 0
	} else {
		leftDepth := maxDepthCompute(root.Left) + 1
		rightDepth := maxDepthCompute(root.Right) + 1
		return int(math.Max(float64(leftDepth), float64(rightDepth)))
	}
}
