package main

import "fmt"

func maxRepeating(sequence string, word string) int {
	var j, ans int
	for i := 0; i < len(sequence); i++ {
		curr := i
		temp := 0
		j = 0
		for curr < len(sequence) && sequence[curr] == word[j] {
			curr++
			j++
			if j == len(word) {
				j = 0
				temp++
				if temp > ans {
					ans = temp
				}
			}
		}

		// 当前循环到尽头 则无需再遍历了
		if curr == len(sequence) {
			break
		}
	}
	return ans
}

func main() {
	fmt.Println(maxRepeating("ababc", "ab"))
	fmt.Println(maxRepeating("ababc", "bc"))
	fmt.Println(maxRepeating("ababc", "ba"))
	fmt.Println(maxRepeating("ababc", "ac"))
}
