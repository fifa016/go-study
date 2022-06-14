/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-14 23:07:17
 */
package l

func numIslands(grid [][]byte) int {
	res := 0

	var dfs func(grid [][]byte, x, y int)

	dfs = func(grid [][]byte, x, y int) {

		if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == '0' {
			return
		}
		grid[x][y] = '0'

		dfs(grid, x-1, y)
		dfs(grid, x+1, y)
		dfs(grid, x, y-1)
		dfs(grid, x, y+1)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				res++
				dfs(grid, i, j)
			}
		}
	}

	return res
}
