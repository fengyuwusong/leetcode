package main

import "fmt"

func removeStars(s string) string {
	var ans []byte
	for i := range s {
		if s[i] == '*' {
			ans = ans[:len(ans)-1]
		} else {
			ans = append(ans, s[i])
		}
	}
	return string(ans)
}

func main() {
	fmt.Println(removeStars("leet**cod*e"))
	fmt.Println(removeStars("erase*****"))
}
