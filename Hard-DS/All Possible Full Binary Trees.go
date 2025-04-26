// https://leetcode.com/problems/all-possible-full-binary-trees

package HardDs

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

// <K,V> = <k, allPossibleFBT(k)>
var mem = map[int][]*TreeNode{}

// n is number of nodes, return the array of root nodes which are full binary trees
func allPossibleFBT(n int) []*TreeNode {
	if n%2 == 0 {
		return nil
	}
	res := make([]*TreeNode, 0)
	if n == 1 {
		res = append(res, &TreeNode{
			Val: 0,
		})
		return res
	}
	if v, found := mem[n]; found {
		return v
	}
	for i := 1; i < n-1; i += 2 {
		var leftList, rightList []*TreeNode
		if v, found := mem[i]; found {
			leftList = v
		} else {
			leftList = allPossibleFBT(i)
		}
		if v, found := mem[n-i-1]; found {
			rightList = v
		} else {
			rightList = allPossibleFBT(n - i - 1)
		}

		for _, l := range leftList {
			for _, r := range rightList {
				res = append(res, &TreeNode{
					Val:   0,
					Left:  l,
					Right: r,
				})
			}
		}
	}
	mem[n] = res
	return res
}
