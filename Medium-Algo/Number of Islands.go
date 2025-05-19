// https://leetcode.com/problems/number-of-islands/description/

package MediumAlgo

func numIslands(grid [][]byte) int {
	islands := 0
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				traverse(i, j, grid)
				islands++
			}

		}
	}
	return islands
}

func traverse(x, y int, grid [][]byte) {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[x]) || grid[x][y] == '0' {
		return
	}

	grid[x][y] = '0'

	traverse(x+1, y, grid)
	traverse(x-1, y, grid)
	traverse(x, y+1, grid)
	traverse(x, y-1, grid)
}
