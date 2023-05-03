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

		foundGroupIndex := -1
		// the simplest way to get merged groups is to create new slice of groups
		newGroups := make([][]string, 0, len(strs))

		for i, group := range groups {

			if isInGroup(group, str) == false {
				newGroups = append(newGroups, group)
				continue
			}

			if foundGroupIndex == -1 {

				group = append(group, str)
				newGroups = append(newGroups, group)
				foundGroupIndex = i

			} else {

				mergedGroup := append(groups[foundGroupIndex], groups[i]...)
				newGroups[foundGroupIndex] = mergedGroup

			}

		}

		if foundGroupIndex == -1 {

			currentGroup := make([]string, 1)
			currentGroup[0] = str
			newGroups = append(newGroups, currentGroup)
		}

		groups = newGroups

	}

	return len(groups)
}

func main() {


	// strs := []string{"kccomwcgcs","socgcmcwkc","sgckwcmcoc","coswcmcgkc","cowkccmsgc","cosgmccwkc","sgmkwcccoc","coswmccgkc","kowcccmsgc","kgcomwcccs"}
	// result := numSimilarGroups(strs)
	// fmt.Println(result, 5)

	strs := []string{"tars","rats","arts","star"}
	result := numSimilarGroups(strs)
	fmt.Println(result, 2)

	// strs := []string{"omv","ovm"}
	// result := numSimilarGroups(strs)
	// fmt.Println(result, 1)

}