package main

import (
	"fmt"
	"strings"
)

func main() {
	input := `....###.
#...####
##.#.###
..#.#...
##.#.#.#
#.######
..#..#.#
######.#`

	grid := initGrid(input)

	fmt.Println("Before any cycles:")
	printGrid(grid)

	for i := 0; i < 6; i++ {
		grid = simulateCycle(grid)
		fmt.Printf("After %d cycles:\n", i+1)
		printGrid(grid)
	}

	fmt.Println(countActive(grid))
}

func countActive(grid [][][]bool) int {
	active := 0
	for _, layer := range grid {
		for _, row := range layer {
			for _, cube := range row {
				if cube {
					active++
				}
			}
		}
	}
	return active
}

func simulateCycle(grid [][][]bool) [][][]bool {
	extendedGrid := extendGrid(grid)
	newGeneration := make3DArray(getSizes(extendedGrid))
	xSize, ySize, zSize := getSizes(extendedGrid)
	for x, layer := range extendedGrid {
		for y, row := range layer {
			for z, cube := range row {
				newGeneration[x][y][z] = getNewCube(cube, x, y, z, xSize, ySize, zSize, extendedGrid)
			}
		}
	}
	return newGeneration
}

func extendGrid(grid [][][]bool) [][][]bool {
	xSize, ySize, zSize := getSizes(grid)
	extendedGrid := make3DArray(xSize+2, ySize+2, zSize+2)
	for i := 1; i < xSize+1; i++ {
		for j := 1; j < ySize+1; j++ {
			for k := 1; k < zSize+1; k++ {
				extendedGrid[i][j][k] = grid[i-1][j-1][k-1]
			}
		}
	}
	return extendedGrid
}

func getSizes(grid [][][]bool) (int, int, int) {
	return len(grid), len(grid[0]), len(grid[0][0])
}

func getNewCube(cube bool, x, y, z, xSize, ySize, zSize int, grid [][][]bool) bool {
	activeNeighbours := 0
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			for k := z - 1; k <= z+1; k++ {
				if i < 0 || j < 0 || k < 0 || i >= xSize || j >= ySize || k >= zSize || (i == x && j == y && k == z) {
					continue
				}
				if grid[i][j][k] {
					activeNeighbours++
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

func printGrid(grid [][][]bool) {
	for _, layer := range grid {
		for _, row := range layer {
			for _, cube := range row {
				if cube {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func initGrid(input string) [][][]bool {
	var grid [][][]bool

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

	grid = append(grid, layer)

	return grid
}

func make3DArray(x, y, z int) [][][]bool {
	var result [][][]bool
	for i := 0; i < x; i++ {
		var layer [][]bool
		for j := 0; j < y; j++ {
			layer = append(layer, make([]bool, z, z))
		}
		result = append(result, layer)
	}
	return result
}
