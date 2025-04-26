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
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	queue := []*TreeNode{root}

	res := []int{}
	traversal := [][]int{}

	for len(queue) > 0 {
		lenCurrLevel := len(queue)

		currT := []int{}

		for i := 0; i < lenCurrLevel; i++ {
			pop := queue[0]
			queue = queue[1:]
			currT = append(currT, pop.Val)
			if pop.Left != nil {
				queue = append(queue, pop.Left)
			}
			if pop.Right != nil {
				queue = append(queue, pop.Right)
			}
		}
		traversal = append(traversal, currT)
	}

	for _, row := range traversal {
		res = append(res, row[len(row)-1])
	}

	return res
}
