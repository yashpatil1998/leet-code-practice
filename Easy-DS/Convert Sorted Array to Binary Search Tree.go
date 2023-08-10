// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree

package EasyDS

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    return utilFunc(nums, 0, len(nums) - 1)
}

func utilFunc(nums []int, start, end int) *TreeNode {
    if start > end {
        return nil
    }
    mid := (start + end) / 2
    node := &TreeNode{
        Val: nums[mid],
    }
    node.Left = utilFunc(nums, start, mid - 1)
    node.Right = utilFunc(nums, mid + 1, end);
    return node;
}