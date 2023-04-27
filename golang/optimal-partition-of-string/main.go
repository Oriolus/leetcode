package main

import "fmt"

func partitionString(s string) int {
	var result int = 0

	m := make(map[string]bool)
	for _, val := range(s) {
		
		var c string = string(val)
		_, exists := m[c]
		if exists == true {
			result += 1
			m = make(map[string]bool)
			m[c] = true
		} else {
			m[c] = true
		}

	}

	if len(m) > 0 {
		result += 1
	}

    return result
}

func main() {

	result := partitionString("aa")
	fmt.Println(result)
}