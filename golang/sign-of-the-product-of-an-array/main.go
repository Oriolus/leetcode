package main

// https://leetcode.com/problems/sign-of-the-product-of-an-array/

import "fmt"

func arraySign(nums []int) int {

	negCount := 0
    for _, val := range nums {
        if val == 0 {
            return 0
        }

        if val < 0 {
            negCount += 1
        }

    }

    if negCount % 2 == 1 {
        return -1
    }
    return 1
}

func main() {
	
	nums := []int{-1,-2,-3,-4,3,2,1}
	result := arraySign([]int{})
	fmt.Println(result)

}