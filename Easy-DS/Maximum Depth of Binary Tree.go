// https://leetcode.com/problems/maximum-depth-of-binary-tree

package EasyDS
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    leftH := maxDepth(root.Left)
    rightH := maxDepth(root.Right)
    if leftH > rightH {
        return leftH + 1
    } else {
        return rightH + 1
    }
}
