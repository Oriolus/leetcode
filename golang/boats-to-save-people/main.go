package main

import (
	"fmt"
	// "sort"
) 

func addOrInc(weights map[int]int, value int) {
	if _, v_is := weights[value]; v_is == true {
		weights[value] += 1
	} else {
		weights[value] = 1
	}
}

func delOrDec(weights map[int]int, value int) map[int]int {
	value += 1
	return weights
}

func numRescueBoats(people []int, limit int) int {

	weights := make(map[int]int)
	boats := 0

	for _, w := range(people) {

		target := limit - w

		if target == 0 {
			boats += 1
			continue
		}

		val, exists := weights[target]

		if exists == false {
			addOrInc(weights, w)
			continue
		}

		if val == 1 {
			delete(weights, target)
		} else {
			weights[target] -= 1
		}

		boats += 1

	}

	for key, val := range(weights) {
		fmt.Println(key, " ", val)
	}

    return boats
}

func main() {

	people := []int{3,5,3,4,2}
	limit := 5

	result := numRescueBoats(people, limit)

	fmt.Println(result)

}