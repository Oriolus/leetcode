package main

// https://leetcode.com/problems/uncrossed-lines/

import "fmt"

func uncrossedLinesIteration(nums1 []int, nums1NextPosition int, index map[int][]int, nums2LastPosition int , nums1FirstIntrance map[int]int ) int {

	if nums1NextPosition >= len(nums1) {
		return 0
	}

	curNum1 := nums1[nums1NextPosition]
	uncrossingLines := 0

	if curNumIndex, ok := index[curNum1] ; ok {

		for _, val := range curNumIndex {

			if val <= nums2LastPosition {
				continue
			}

			for j := nums1NextPosition + 1 ; j < len(nums1) ; j++ {

				// if nums1FirstIntrance[nums1[j]] < j {
				// 	continue
				// }

				uncrossingLinesCountInNextIteration := uncrossedLinesIteration(nums1, j, index, val, nums1FirstIntrance)

				if 1 + uncrossingLinesCountInNextIteration > uncrossingLines {
					uncrossingLines = 1 + uncrossingLinesCountInNextIteration
				}
					
			}

			if uncrossingLines == 0 {
				uncrossingLines = 1
			}

			break

		}

	} else {
		return 0
	}

	return uncrossingLines
}

func maxUncrossedLines(nums1 []int, nums2 []int) int {

	maxUncrossingLines := 0

	index := make(map[int][]int)
	nums1FirstIntrance := make(map[int]int)

	for ind, val := range nums2 {

		if list, ok := index[val] ; ok {
			list = append(list, ind)
			index[val] = list
		} else {
			l := make([]int, 0, 1)
			l = append(l, ind)
			index[val] = l
		}

	}

	for ind, val := range nums1 {

		if _, ok := nums1FirstIntrance[val] ; !ok {
			nums1FirstIntrance[val] = ind
		}

	}

	for i := 0 ; i < len(nums1) ; i++ {
		iResult := uncrossedLinesIteration(nums1, i, index, -1, nums1FirstIntrance)
		if iResult > maxUncrossingLines {
			maxUncrossingLines = iResult
		}
	}

	return maxUncrossingLines

}

func test1() {
	nums1 := []int{1,4,2}
	nums2 := []int{1,2,4}

	result := maxUncrossedLines(nums1, nums2)
	fmt.Println(result, 2)
}

func test2() {
	nums1 := []int{2,5,1,2,5}
	nums2 := []int{10,5,2,1,5,2}

	result := maxUncrossedLines(nums1, nums2)
	fmt.Println(result, 3)
}

func test3() {
	nums1 := []int{1}
	nums2 := []int{2}

	result := maxUncrossedLines(nums1, nums2)
	fmt.Println(result, 0)
}

func test4() {
	nums1 := []int{1,3,7,1,7,5}
	nums2 := []int{1,9,2,5,1}

	result := maxUncrossedLines(nums1, nums2)
	fmt.Println(result, 2)
}

func test5() {
	nums1 := []int{1}
	nums2 := []int{1,3}

	result := maxUncrossedLines(nums1, nums2)
	fmt.Println(result, 1)
}


func main() {

	test1()
	test2()
	test3()
	test4()
	test5()

}