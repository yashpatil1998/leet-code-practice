// https://leetcode.com/problems/rotting-oranges/

package HardAlgo

func orangesRotting(grid [][]int) int {
	steps := 0

	m := len(grid)
	n := len(grid[0])

	queue := make([][2]int, 0)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	dx := [4]int{0, 0, -1, 1}
	dy := [4]int{1, -1, 0, 0}

	for len(queue) > 0 {
		startingPoints := len(queue)
		stepsIncrement := 0
		for startingPoints > 0 {
			startingPoints--
			pop := queue[0]
			queue = queue[1:]

			for k := 0; k < 4; k++ {
				nx, ny := pop[0]+dx[k], pop[1]+dy[k]
				if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] == 1 {
					grid[nx][ny] = 2
					queue = append(queue, [2]int{nx, ny})
					stepsIncrement = 1
				}
			}

		}
		steps += stepsIncrement
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}

	return steps
}
