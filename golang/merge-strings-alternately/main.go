package main

import (
	"fmt"
)

func mergeAlternately(word1 string, word2 string) string {

	res := make([]byte, len(word1) + len(word2))
	res_ind := 0
	words_index := 0

	for  ; words_index < len(word1) && words_index < len(word2) ; words_index += 1 {
		res[res_ind] = word1[words_index]
		res_ind += 1
		res[res_ind] = word2[words_index]
		res_ind += 1
	}

	for i := words_index ; i < len(word1) ; i += 1 {
		res[res_ind] = word1[i]
		res_ind += 1
	}

	for i := words_index ; i < len(word2) ; i += 1 {
		res[res_ind] = word2[i]
		res_ind += 1
	}

    return string(res[:])
}

func main() {


	// word1, word2 := "abc", "pqr"
	// word1, word2 := "ab", "pqrs"
	word1, word2 := "abcd", "pq"

	result := mergeAlternately(word1, word2)

	fmt.Println(result)

}