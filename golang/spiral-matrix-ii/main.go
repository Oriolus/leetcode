package main

// https://leetcode.com/problems/spiral-matrix-ii/

import "fmt"

// const int NOT_VISITED = 0

func generateMatrix(n int) [][]int {

	m := make([][]int, 0, n)
	for i := 0 ; i < n ; i++ {
		m = append(m, make([]int, n))
	}

	m[0][0] = 1

	
	


	return m
}

func main() {

	n := 3
	matrix := generateMatrix(n)
	fmt.Println(matrix)

}