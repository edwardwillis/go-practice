// Package graphs contains graph traversal interview problems.
package graphs

// NumIslands counts the number of islands ('1' cells connected
// horizontally/vertically) in the grid. O(rows*cols) time and space.
// grid is mutated (flood-filled) as part of the traversal.
func NumIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])

	var sink func(r, c int)
	sink = func(r, c int) {
		if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] != '1' {
			return
		}
		grid[r][c] = '0'
		sink(r-1, c)
		sink(r+1, c)
		sink(r, c-1)
		sink(r, c+1)
	}

	count := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				count++
				sink(r, c)
			}
		}
	}
	return count
}
