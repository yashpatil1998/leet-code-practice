// https://leetcode.com/problems/unique-paths/description/

package MediumAlgo

func uniquePaths(m int, n int) int {
	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		grid[i][0] = 1
	}
	for j := 0; j < n; j++ {
		grid[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
	}

	return grid[m-1][n-1]
}
