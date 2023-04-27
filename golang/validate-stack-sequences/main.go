package main

import "fmt"

// https://leetcode.com/problems/validate-stack-sequences/

/*
 current := result[last]
 if current == pope
*/

func validateStackSequences(pushed []int, popped []int) bool {

	var index_poped, index_pushed, index_result int
	result := make([]int, len(pushed))

	result[index_result] = pushed[index_pushed]
	index_result, index_pushed = 1, 1
	for ; index_poped < len(popped) && index_pushed < len(pushed) && index_result < len(result) ; {

		if index_result > 0 && popped[index_poped] == result[index_result - 1] {
			index_result -= 1;
			index_poped += 1;
		} else {
			result[index_result] = pushed[index_pushed]
			index_result += 1
			index_pushed += 1
		}

	}

	for ; index_poped < len(popped) ; {

		if result[index_result - 1] == popped[index_poped] {
			index_result--
			index_poped++
		} else {
			return false
		}

	}

	return true
}


func main() {

	// pushed := []int{1,2,3,4,5}
	// poped := []int{4,5,3,2,1}
	// pushed := []int{1,2,3,4,5}
	// poped := []int{4,3,5,1,2}
	pushed := []int{1}
	poped := []int{2}

	result := validateStackSequences(pushed, poped)
	fmt.Println(result)
}