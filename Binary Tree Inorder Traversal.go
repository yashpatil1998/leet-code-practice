// https://leetcode.com/problems/binary-tree-inorder-traversal/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func inorderTraversal(root *TreeNode) []int {
    sol := []int{}
    if root == nil {
        return sol
    }
    sol = append(sol, inorderTraversal(root.Left)...)
    sol = append(sol, root.Val)
    sol = append(sol, inorderTraversal(root.Right)...)

    return sol
}