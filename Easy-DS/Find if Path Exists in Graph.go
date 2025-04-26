// https://leetcode.com/problems/find-if-path-exists-in-graph/

package EasyDS

func validPath(n int, edges [][]int, source int, destination int) bool {
	graph := map[int][]int{}

	for _, edge := range edges {
		src := edge[0]
		dst := edge[1]
		graph[src] = append(graph[src], dst)
		graph[dst] = append(graph[dst], src)
	}

	visited := map[int]bool{}

	queue := []int{source}
	for len(queue) > 0 {
		curNode := queue[0]
		queue = queue[1:]

		if curNode == destination {
			return true
		}

		if !visited[curNode] {
			queue = append(queue, graph[curNode]...)
		}

		visited[curNode] = true

	}
	return false
}
