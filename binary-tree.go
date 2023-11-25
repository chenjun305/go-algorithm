package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树的中序遍历
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	arr := inorderTraversal(root.Left)
	arr = append(arr, root.Val)
	return append(arr, inorderTraversal(root.Right)...)
}

// 二叉树的 最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left >= right {
		return left + 1
	} else {
		return right + 1
	}
}

// 翻转二叉树
// 给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。
func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	}
	return root
}

// 对称二叉树
// 给你一个二叉树的根节点 root ， 检查它是否轴对称。
//
// 该问题可以转化为：两个树在什么情况下互为镜像？
// 如果同时满足下面的条件，两个树互为镜像：
// 1, 它们的两个根结点具有相同的值
// 2, 每个树的右子树都与另一个树的左子树镜像对称
func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

func check(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && check(left.Left, right.Right) && check(left.Right, right.Left)
}

// 将有序数组转换为二叉搜索树
func sortedArrayToBST(nums []int) *TreeNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	root := new(TreeNode)
	mid := len(nums) / 2
	root.Val = nums[mid]
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}

/* 二叉树的直径
* 给你一棵二叉树的根节点，返回该树的 直径 。
* 二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。
*
* 递归搜索每个节点的深度，并设一个全局变量的直径，直径与左右子树高度和进行比较
 */
var diameter int = 0

func diameterOfBinaryTree(root *TreeNode) int {
	diameter = 0
	depth(root)
	return diameter
}

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := depth(node.Left)
	right := depth(node.Right)
	if (left + right) > diameter {
		diameter = left + right
	}
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}
