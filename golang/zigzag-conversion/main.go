package main

import "fmt"

// https://leetcode.com/problems/zigzag-conversion/

func convert(s string, numRows int) string {

	period := 2 * numRows - 2

	if period == 0 || period == 1 {
		return s
	}

	r := len(s) / period

	var res string = ""

	for row := 0 ; row < numRows ; row += 1 {

		for col := 0 ; col < r + 1 ; col += 1 {

			if ind := row + col * period ; ind < len(s) {
				res += string(s[ind])
			}

			if row >= 1 && row != numRows - 1 {
				
				if ind := (col + 1) * period - row ; ind < len(s) {
					res += string(s[ind])
				}

			}

		}

	}

	return res
}

func main() {

	result := convert("A", 1)
	fmt.Println(result)
}