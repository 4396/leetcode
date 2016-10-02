// https://leetcode.com/problems/sum-of-left-leaves/

package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var left, right int
	if root.Left != nil {
		if root.Left.Left != nil {
			left = sumOfLeftLeaves(root.Left)
		} else if root.Left.Right == nil {
			left = root.Left.Val
		}
	}
	if root.Right != nil {
		right = sumOfLeftLeaves(root.Right)
	}
	return left + right
}
