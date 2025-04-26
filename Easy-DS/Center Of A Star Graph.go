// https://leetcode.com/problems/find-center-of-star-graph/

package EasyDS

func findCenter(edges [][]int) int {
	if edges[1][0] == edges[0][0] || edges[1][0] == edges[0][1] {
		return edges[1][0]
	} else {
		return edges[1][1]
	}
}
