package main

import (
	"fmt"
	"strings"
)

//G -> G
//() -> o
//(al) -> al
func interpret(command string) string {
	var ans strings.Builder
	for i := 0; i < len(command); i++ {
		switch command[i] {
		case 'G':
			ans.WriteRune('G')
		case ')':
			ans.WriteRune('o')
		case 'a':
			ans.WriteString("al")
			i += 2
		}
	}
	return ans.String()
}

func main() {
	fmt.Println(interpret("G()(al)"))
	fmt.Println(interpret("G()()()()(al)"))
	fmt.Println(interpret("(al)G(al)()()G"))

}
