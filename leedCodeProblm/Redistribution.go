package leedcodeproblm

import "fmt"

func RedistributionCharacter(words []string) bool {
	frequency := make([]int, 26)
	for _, word := range words {
		for _, c := range word {
			frequency[c-'a']++
		}
	}

	for _, value := range frequency {
		fmt.Println(frequency[i])
		if frequency[i]%len(words) != 0 {
			return false
		}
	}

	return true
}
