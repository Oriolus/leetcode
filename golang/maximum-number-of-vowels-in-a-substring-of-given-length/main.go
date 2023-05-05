package main

import "fmt"

func isVowel(s string) int {

	if s == "a" || s == "e" || s == "i" || s == "o" || s == "u" {
		return 1
	}

	return 0
}

func maxVowels(s string, k int) int {
    
	vowelsCount := 0

	// init sliding windows
	for i := 0 ; i < k && i < len(s) ; i += 1 {
		vowelsCount += isVowel(string(s[i]))
	}

	maxNumbersOfVowels := vowelsCount

	// go through string

	for i := k ; i < len(s) ; i += 1 {

		vowelsCount -= isVowel(string(s[i - k]))
		vowelsCount += isVowel(string(s[i]))

		if vowelsCount > maxNumbersOfVowels {
			maxNumbersOfVowels = vowelsCount
		}

	}

	return maxNumbersOfVowels
}

func main() {

	fmt.Println(maxVowels("abciiidef", 3), 3)
	fmt.Println(maxVowels("aeiou", 2), 2)
	fmt.Println(maxVowels("leetcode", 3), 2)
	fmt.Println(maxVowels("e", 3), 1)

}