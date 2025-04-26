// https://leetcode.com/problems/binary-tree-level-order-traversal/

package MediumDS

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {

		level := len(queue)
		currentLevel := []int{}

		for i := 0; i < level; i++ {
			pop := queue[0]
			queue = queue[1:]

			currentLevel = append(currentLevel, pop.Val)

			if pop.Left != nil {
				queue = append(queue, pop.Left)
			}

			if pop.Right != nil {
				queue = append(queue, pop.Right)
			}
		}

		res = append(res, currentLevel)

	}

	return res
}
