package main

import (
	"fmt"
)

func countConsistentStrings(allowed string, words []string) int {
	var ans int
	memo := make(map[rune]bool)
	for _, r := range allowed {
		memo[r] = true
	}
	for _, s := range words {
		flag := false
		for _, r := range s {
			if !memo[r] {
				flag = true
				break
			}
		}
		if !flag {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(countConsistentStrings("ab", []string{"ad", "db", "baa", "badab"}))
}
