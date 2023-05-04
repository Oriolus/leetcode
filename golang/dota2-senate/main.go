package main

// https://leetcode.com/problems/dota2-senate/

import "fmt"

func getNextBanned(senate string, banned []bool, currentIndex int) (bool, int) {

	currentSenator := senate[currentIndex]

	for i := currentIndex + 1 ; i < len(senate) ; i += 1 {

		if banned[i] == true {
			continue
		}

		if senate[i] != currentSenator {
			return true, i
		}

	}

	if currentIndex == 0 {
		return false, -1
	}

	for i := 0; i < currentIndex ; i += 1 {

		if banned[i] == true {
			continue
		}

		if senate[i] != currentSenator {
			return true, i
		}
			
	}

	return false, -1
}

func predictPartyVictory(senate string) string {
    
	banned := make([]bool, len(senate))
	rCount := 0
	dCount := 0

	for _, v := range senate {
		if string(v) == "D" {
			dCount += 1
		} else {
			rCount += 1
		}
	}

	for rCount > 0 && dCount > 0 {

		for i := 0 ; i < len(senate) ; i += 1 {

			if banned[i] == true {
				continue
			}

			found, banningIndex := getNextBanned(senate, banned, i)

			if found == false {

				if string(senate[i]) == "R" {
					return "Radiant"
				}
				return "Dire"
			}

			if string(senate[banningIndex]) == "D" {
				dCount -= 1
			} else {
				rCount -= 1
			}

			banned[banningIndex] = true

		}

	}

	if dCount == 0 {
		return "Radiant"
	}
	return "Dire"
}

func main() {
	
	fmt.Println(predictPartyVictory("RD"), "Radiant")
	fmt.Println(predictPartyVictory("RDD"), "Dire")
	fmt.Println(predictPartyVictory("DDRRR"), "Dire")
	fmt.Println(predictPartyVictory("DRRD"), "Dire")
	fmt.Println(predictPartyVictory("D"), "Dire")

/*
	1. найти сенатора, который может ходить
	2. для этого сенатора найти сенатора противоположной команды, которого можно забанить (первого справа по кольцу)
		2.1 забанить его
		2.2 если банить некого - вернуть код текущего сенатора
	3. повторять до тех пор, пока количество незабанненых сенаторов одной из команд будет 0
*/

}