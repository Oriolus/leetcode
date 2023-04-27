package main

import (
	"fmt"
	"math"
)

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxOfInts(a []int) int {
	max := a[0]
	for _, v := range(a) {
		if max < v {
			max = v
		}
	}
	return max
}

func minimizeArrayValue(nums []int) int {
	sum := 0
	for _, val := range(nums) {
		sum += val
	}
	optimum := int(math.Ceil(float64(sum) / float64(len(nums))))
	fmt.Println(optimum)
	fmt.Println(nums	)

	for i := 1; i < len(nums); i += 1 {
		l, r := nums[i - 1], nums[i]
		
		if l >= r || r == 1 || l >= optimum {
			continue
		}

		dif := min(optimum - l, r - 1)
		nums[i - 1], nums[i] = l + dif, r - dif
		fmt.Println(nums)
	}

	return maxOfInts(nums)
}

func main() {

	// nums := []int{3,7,1,6}
	// nums := []int{10, 1}
	nums := []int{13,13,20,0,8,9,9}

	result := minimizeArrayValue(nums)

	fmt.Println(result)

}