package main

// https://leetcode.com/problems/similar-string-groups/

import "fmt"

/*
	todo:

	это слова из одной группы. Но если в массиве сначала проверяем 1 и 3, а затем 2 и 1, получается 3 не входит в группу
	coswcmcgkc
	coswmccgkc
	cosgmccwkc
*/

func areSimilar(a string, b string) bool {

	firstDifferenceIndex := -1
	secondDifferenceIndex := -1

	for i := 0 ; i < len(a) ; i++ {

		if a[i] == b[i] {
			continue
		}

		if firstDifferenceIndex == -1 {
			firstDifferenceIndex = i
		} else if secondDifferenceIndex == -1 {
			secondDifferenceIndex = i
		} else {
			return false
		}

	}

	if firstDifferenceIndex == -1 {
		return true
	}

	if secondDifferenceIndex == -1 {
		return false
	}

	return a[firstDifferenceIndex] == b[secondDifferenceIndex] && a[secondDifferenceIndex] == b[firstDifferenceIndex]

}

func isInGroup(group []string, a string) bool {

	for _, val := range group {
		if areSimilar(val, a) == true {
			return true
		}
	}

	return false
}

func numSimilarGroups(strs []string) int {

	groups := make([][]string, 0, len(strs))

	for _, str := range strs {

		groupFound := false

		for i, group := range groups {

			if isInGroup(group, str) {

				group = append(group, str)
				groups[i] = group
				groupFound = true
				break

			}

		}

		if groupFound == false {

			currentGroup := make([]string, 1)
			currentGroup[0] = str
			groups = append(groups, currentGroup)
		}

	}

	fmt.Println(groups)

	return len(groups)
}

func main() {

	fmt.Println(areSimilar("tars", "rats"), true)
	fmt.Println(areSimilar("rats", "arts"), true)
	fmt.Println(areSimilar("rats", "star"), false)
	fmt.Println(areSimilar("coswmccgkc", "cosgmccwkc"), true)

	strs := []string{"kccomwcgcs","socgcmcwkc","sgckwcmcoc","coswcmcgkc","cowkccmsgc","cosgmccwkc","sgmkwcccoc","coswmccgkc","kowcccmsgc","kgcomwcccs"}
	result := numSimilarGroups(strs)
	fmt.Println(result, 2)

	

	// strs := []string{"tars","rats","arts","star"}
	// result := numSimilarGroups(strs)
	// fmt.Println(result, 2)

	// strs = []string{"omv","ovm"}
	// result = numSimilarGroups(strs)
	// fmt.Println(result, 1)

}