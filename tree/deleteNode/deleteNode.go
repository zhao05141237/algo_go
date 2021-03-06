package deleteNode

import (
	"algo/tree/tree"
)

/**
给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。
一般来说，删除节点可分为两个步骤：
首先找到需要删除的节点；
如果找到了，删除它。
说明： 要求算法时间复杂度为 O(h)，h 为树的高度。
示例:
root = [5,3,6,2,4,null,7]
key = 3
    5
   / \
  3   6
 / \   \
2   4   7
给定需要删除的节点值是 3，所以我们首先找到 3 这个节点，然后删除它。
一个正确的答案是 [5,4,6,2,null,null,7], 如下图所示。
    5
   / \
  4   6
 /     \
2       7
另一个正确答案是 [5,2,6,null,4,null,7]。
    5
   / \
  2   6
   \   \
    4   7

执行用时 :20 ms, 在所有 golang 提交中击败了76.74%的用户
内存消耗 :6.7 MB, 在所有 golang 提交中击败了100.00%的用户
*/
func DeleteNode(root *tree.Node, k int) *tree.Node {
	//node value=key的节点
	//node_parent value=key的父节点
	node, nodeParent := findNode(root, nil, k)
	if node == nil {
		//没有找到 返回整个树
		return root
	}
	if nodeParent == nil {
		//根结点
		if root.Left == nil && root.Right == nil {
			//左右子树都是空
			return nil
		}
		if root.Left == nil {
			//左子树空 返回右子树
			return root.Right
		}
		if root.Right == nil {
			//右子树空 返回左子树
			return root.Left
		}
	}
	moveNode(node, nodeParent)
	return root
}

func findNode(root *tree.Node, parent *tree.Node, k int) (*tree.Node, *tree.Node) {
	if root == nil {
		return nil, parent
	} else if root.Val == k {
		return root, parent
	} else if root.Val < k {
		return findNode(root.Right, root, k)
	} else if root.Val > k {
		return findNode(root.Left, root, k)
	} else {
		return nil, nil
	}
}

func moveNode(root *tree.Node, parent *tree.Node) {
	//空节点
	if root == nil {
		return
	}

	//叶子节点
	if root.Left == nil && root.Right == nil {
		if parent.Val > root.Val {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
		return
	}

	node := root
	mnode := node

	if parent == nil {
		//查找节点正好是根结点 并且根节点的右子树是空
		if node.Left != nil && node.Right == nil {
			root = node.Left
			return
		} else if node.Left == nil && node.Right != nil {
			root = node.Right
			return
		}

	} else {
		//查找节点的左子树或者右子树是空的情况
		if node.Left == nil || node.Right == nil {
			if parent.Left == root {
				if node.Left != nil {
					parent.Left = node.Left
					return
				} else if node.Right != nil {
					parent.Left = node.Right
					return
				}
			} else if parent.Right == root {
				if node.Left != nil {
					parent.Right = node.Left
					return
				} else if node.Right != nil {
					parent.Right = node.Right
					return
				}
			}
		}

	}

	//从查找节点形如遍历 一直遍历到叶子节点
	for node.Left != nil && node.Right != nil {
		mnode = node
		//左右子树都不为空
		if node.Left != nil && node.Right != nil {
			//先用左子树根节点覆盖
			node.Val = node.Left.Val
			node = node.Left
			//右子树不空 可能比根结点大 这时需要移动到根结点的右子树中
			if node.Right != nil {
				topNode := root
				//找到比他小的根结点 topNode
				for topNode.Val > node.Right.Val && topNode.Left != nil {
					topNode = topNode.Left
				}
				//最右结点
				for topNode.Right != nil && topNode.Val < node.Right.Val {
					topNode = topNode.Right
				}
				for topNode.Left != nil {
					topNode = topNode.Left
				}
				topNode.Left = node.Right
				node.Right = nil
			}
			//左子树不空 右子结空 直接替换
		} else if node.Left != nil {
			node.Val = node.Left.Val
			node = node.Left
			//左子树空 右子结不空
		} else if node.Right != nil {
			topNode := root
			//找到比他小的根结点 topNode
			for topNode.Val > node.Right.Val && topNode.Left != nil {
				topNode = topNode.Left
			}
			//最右结点
			for topNode.Right != nil && topNode.Val < node.Right.Val {
				topNode = topNode.Right
			}
			for topNode.Left != nil {
				topNode = topNode.Left
			}
			topNode.Left = node
		}
	}
	if node.Left != nil {
		mnode.Left = node.Left
	} else {
		mnode.Left = nil
	}
}

func InsertNode(root *tree.Node, k int) *tree.Node {
	if root == nil {
		return &tree.Node{
			Val:   k,
			Left:  nil,
			Right: nil,
		}
	}

	if root.Val < k {
		root.Right = insertNodeC(root.Right, k)
	}

	if root.Val >= k {
		root.Left = insertNodeC(root.Left, k)
	}

	return root
}

func insertNodeC(root *tree.Node, k int) *tree.Node {
	if root == nil {
		return &tree.Node{
			Val:   k,
			Left:  nil,
			Right: nil,
		}
	}

	if root.Left == nil && root.Val >= k {
		root.Left = &tree.Node{
			Val:   k,
			Left:  nil,
			Right: nil,
		}
		return root
	}

	if root.Right == nil && root.Val < k {
		root.Right = &tree.Node{
			Val:   k,
			Left:  nil,
			Right: nil,
		}
		return root
	}

	if root.Val < k {
		root.Right = insertNodeC(root.Right, k)
	}

	if root.Val >= k {
		root.Left = insertNodeC(root.Left, k)
	}

	return root
}
