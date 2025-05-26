// https://leetcode.com/problems/largest-color-value-in-a-directed-graph/description/

package HardDs

func largestPathValue(colors string, edges [][]int) int {
	n := len(colors)
	graph := make(map[int][]int)
	inDegree := make(map[int]int)

	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		inDegree[edge[1]]++
	}

	queue := make([]int, 0)

	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	if len(queue) == 0 {
		return -1
	}

	colorMem := make(map[int][26]int)

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	ans := 0
	nodesProcessed := 0

	for len(queue) > 0 {
		nodesProcessed++
		pop := queue[0]
		queue = queue[1:]
		currentPopColorCounts := colorMem[pop]
		indexToIncrement := colors[pop] - 'a'
		currentPopColorCounts[indexToIncrement]++
		colorMem[pop] = currentPopColorCounts

		ans = max(ans, colorMem[pop][colors[pop]-'a'])

		for _, neighbour := range graph[pop] {
			inDegree[neighbour]--
			if inDegree[neighbour] == 0 {
				queue = append(queue, neighbour)
			}

			neighbourColorCounts := colorMem[neighbour]
			for i := 0; i < 26; i++ {
				neighbourColorCounts[i] = max(neighbourColorCounts[i], colorMem[pop][i])
			}
			colorMem[neighbour] = neighbourColorCounts

		}

	}

	if nodesProcessed != n {
		return -1
	}

	return ans

}
