package main

// https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/solutions/

import "fmt"

const (
	aSymbol = 97
	eSymbol = 101
	iSymbol = 105
	oSymbol = 111
	uSymbol = 117
)

func isVowel(s byte) int {

	if s == aSymbol || s == eSymbol || s == iSymbol || s == oSymbol || s == uSymbol {
		return 1
	}

	return 0
}

func maxVowels(s string, k int) int {
    
	vowelsCount := 0

	// init sliding windows
	for i := 0 ; i < k && i < len(s) ; i += 1 {
		vowelsCount += isVowel(s[i])
	}

	maxNumbersOfVowels := vowelsCount

	// go through string

	for i := k ; i < len(s) ; i += 1 {

		vowelsCount -= isVowel(s[i - k])
		vowelsCount += isVowel(s[i])

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