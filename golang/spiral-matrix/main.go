package main

// https://leetcode.com/problems/spiral-matrix/

import "fmt"

const VISITED int = -101

func spiralOrder(matrix [][]int) []int {

	wasMoved := true
	iPointer, jPointer := 0, 0
	spiral := make([]int, 0, len(matrix) * len(matrix[0]))
	
	spiral = append(spiral, matrix[0][0])
	matrix[0][0] = VISITED

	for ; wasMoved ; {

		wasMoved = false

		if jPointer < len(matrix[iPointer]) - 1 && matrix[iPointer][jPointer + 1] != VISITED {

			jPointer++

			for ; jPointer < len(matrix[iPointer]) && matrix[iPointer][jPointer] != VISITED ; jPointer++ {
				spiral = append(spiral, matrix[iPointer][jPointer])
				wasMoved = true
				matrix[iPointer][jPointer] = VISITED
			}

			jPointer--

		}

		if iPointer < len(matrix) - 1 && matrix[iPointer + 1][jPointer] != VISITED {

			iPointer++

			for ; iPointer < len(matrix) && matrix[iPointer][jPointer] != VISITED ; iPointer++ {
				spiral = append(spiral, matrix[iPointer][jPointer])
				matrix[iPointer][jPointer] = VISITED
				wasMoved = true
			}

			iPointer--

		}

		if jPointer > 0 && matrix[iPointer][jPointer - 1] != VISITED {

			jPointer--

			for ; jPointer >= 0 && matrix[iPointer][jPointer] != VISITED ; jPointer-- {
				spiral = append(spiral, matrix[iPointer][jPointer])
				matrix[iPointer][jPointer] = VISITED
				wasMoved = true
			}

			jPointer++

		}

		if iPointer > 0 && matrix[iPointer - 1][jPointer] != VISITED {

			iPointer--

			for ; iPointer >= 0 && matrix[iPointer][jPointer] != VISITED ; iPointer-- {
				spiral = append(spiral, matrix[iPointer][jPointer])
				matrix[iPointer][jPointer] = VISITED
				wasMoved = true
			}

			iPointer++

		}

	}

	return spiral
}

func main() {
	
	// matrix := [][]int {
	// 	{1,2,3},
	// 	{4,5,6},
	// 	{7,8,9},
	// }

	// matrix := [][]int{{1},{2},{3}}
	matrix := [][]int{{1}}

	// matrix := [][]int{{1,2,3,4},{5,6,7,8},{9,10,11,12}}

	result := spiralOrder(matrix)
	fmt.Println(result)

}