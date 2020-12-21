package main

import (
	"fmt"
	"strings"
)

func main() {
	//input := `.#.
	//..#
	//###`

	input := `....###.
#...####
##.#.###
..#.#...
##.#.#.#
#.######
..#..#.#
######.#`

	grid := initGrid(input)

	for i := 0; i < 6; i++ {
		grid = simulateCycle(grid)
	}

	fmt.Println(countActive(grid))
}

func countActive(grid [][][][]bool) int {
	active := 0
	for _, layer := range grid {
		for _, row := range layer {
			for _, hypercube := range row {
				for _, cube := range hypercube {
					if cube {
						active++
					}
				}
			}
		}
	}
	return active
}

func simulateCycle(grid [][][][]bool) [][][][]bool {
	extendedGrid := extendGrid(grid)
	newGeneration := make4DArray(getSizes(extendedGrid))
	xSize, ySize, zSize, wSize := getSizes(extendedGrid)
	for x, layer := range extendedGrid {
		for y, row := range layer {
			for z, hypercube := range row {
				for w, cube := range hypercube {
					newGeneration[x][y][z][w] = getNewCube(cube, x, y, z, w, xSize, ySize, zSize, wSize, extendedGrid)
				}
			}
		}
	}
	return newGeneration
}

func extendGrid(grid [][][][]bool) [][][][]bool {
	xSize, ySize, zSize, wSize := getSizes(grid)
	extendedGrid := make4DArray(xSize+2, ySize+2, zSize+2, wSize+2)
	for i := 1; i < xSize+1; i++ {
		for j := 1; j < ySize+1; j++ {
			for k := 1; k < zSize+1; k++ {
				for l := 1; l < wSize+1; l++ {
					extendedGrid[i][j][k][l] = grid[i-1][j-1][k-1][l-1]
				}
			}
		}
	}
	return extendedGrid
}

func getSizes(grid [][][][]bool) (int, int, int, int) {
	return len(grid), len(grid[0]), len(grid[0][0]), len(grid[0][0][0])
}

func getNewCube(cube bool, x, y, z, w, xSize, ySize, zSize, wSize int, grid [][][][]bool) bool {
	activeNeighbours := 0
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			for k := z - 1; k <= z+1; k++ {
				for l := w - 1; l <= w+1; l++ {
					if i < 0 || j < 0 || k < 0 || l < 0 || i >= xSize || j >= ySize || k >= zSize || l >= wSize || (i == x && j == y && k == z && l == w) {
						continue
					}
					if grid[i][j][k][l] {
						activeNeighbours++
					}
				}
			}
		}
	}
	if cube && (activeNeighbours != 2 && activeNeighbours != 3) {
		return false
	}
	if !cube && (activeNeighbours != 3) {
		return false
	}
	return true
}

func initGrid(input string) [][][][]bool {
	var grid [][][][]bool

	var layer [][]bool
	for _, rowRaw := range strings.Split(input, "\n") {
		var row []bool
		for _, cube := range rowRaw {
			element := true
			if cube == '.' {
				element = false
			}
			row = append(row, element)
		}
		layer = append(layer, row)
	}

	grid = append(grid, [][][]bool{layer})

	return grid
}

func make4DArray(x, y, z, w int) [][][][]bool {
	result := make([][][][]bool, x, x)
	for i := 0; i < x; i++ {
		result[i] = make([][][]bool, y, y)
		for j := 0; j < y; j++ {
			result[i][j] = make([][]bool, z, z)
			for k := 0; k < z; k++ {
				result[i][j][k] = make([]bool, w, w)
			}
		}
	}
	return result
}
