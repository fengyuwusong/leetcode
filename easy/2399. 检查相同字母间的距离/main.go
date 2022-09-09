package main

import "fmt"

func checkDistances(s string, distance []int) bool {
	memo := make(map[rune]int)
	for i, b := range s {
		if start, ok := memo[b]; !ok {
			memo[b] = i
		} else if i-start-1 != distance[b-'a'] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(checkDistances("abaccb", []int{1, 3, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
	fmt.Println(checkDistances("aa", []int{1, 3, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
}
