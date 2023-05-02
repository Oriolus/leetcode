package main

// https://leetcode.com/problems/longest-substring-without-repeating-characters/

import "fmt"

/*
	Given a string s, find the length of the longest substring without repeating characters.

	1. Using hashset to save last carachter position
	2. When finding caracheter that exists in hashset
		2.1 save result if it higher of last result
		2.2 remove values in hashset with indexes lower or equal than current carachter last interence


	DP:
		using
			- lasts = hashset for saving last carachter index
			- dp = array for saving current max non repeated symbols substring
	
		dp saves current max non-repeated symbols substring
		i - current symbol index
			lasts[s[i]] < i - dp[i - 1]: dp[i] = dp[i] + 1
			else: dp[i] = i - lasts[s[i]]
			save lasts[s[i]] = i

*/

func lengthOfLongestSubstring(s string) int {

	if s == "" {
		return 0
	}

	result := 1

	lasts := make(map[string]int, 128)
	dp := make([]int, len(s))
	dp[0] = 1
	lasts[string(s[0])] = 0

	for i := 1 ; i < len(s) ; i += 1 {

		current := string(s[i])
		currentLastIndex, exists := lasts[current]

		if exists == true && i - dp[i - 1] <= currentLastIndex {
			dp[i] = i - currentLastIndex
		} else {
			dp[i] = dp[i - 1] + 1
		}

		if result < dp[i] {
			result = dp[i]
		}
		lasts[current] = i

	}

	return result
}

func main() {

	result := lengthOfLongestSubstring("abcabcbb")
	fmt.Println(result)

}