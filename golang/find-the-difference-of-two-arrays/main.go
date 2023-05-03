package main

import "fmt"

func findDifference(nums1 []int, nums2 []int) [][]int {

	map1 := make(map[int]bool, len(nums1))
	for _, val := range nums1 {
		map1[val] = true
	}

	map2 := make(map[int]bool, len(nums2))
	for _, val := range nums2 {
		map2[val] = true
	}

	diff1 := make([]int, 0, len(nums1))
	for k := range map1 {
		
		if _, ok := map2[k]; ok == false {
			diff1 = append(diff1, k)
		}

	}

	diff2 := make([]int, 0, len(nums2))
	for k := range map2 {

		if _, ok := map1[k]; ok == false {
			diff2 = append(diff2, k)
		}

	}

	res := make([][]int, 2)
	res[0] = diff1
	res[1] = diff2

	return res
}

func main() {

	nums1 := []int{1,2,3}
	nums2 := []int{2,4,6}
	result := findDifference(nums1, nums2)
	fmt.Println(result)

}