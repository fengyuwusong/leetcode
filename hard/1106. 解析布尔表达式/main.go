package main

import "fmt"

func parseBoolExpr(expression string) bool {
	var stack []rune
	for _, s := range expression {
		if s != ',' && s != ')' {
			stack = append(stack, s)
			continue
		}
		if s == ')' {
			tNum, fNum := 0, 0
			for {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if top == 't' {
					tNum++
				} else if top == 'f' {
					fNum++
				} else if top == '(' {
					break
				}
			}
			op := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			switch op {
			case '&':
				if fNum > 0 {
					stack = append(stack, 'f')
				} else {
					stack = append(stack, 't')
				}
			case '|':
				if tNum > 0 {
					stack = append(stack, 't')
				} else {
					stack = append(stack, 'f')
				}
			case '!':
				if tNum > 0 {
					stack = append(stack, 'f')
				} else {
					stack = append(stack, 't')
				}
			}
		}
	}
	return stack[len(stack)-1] == 't'
}
func main() {
	fmt.Println(parseBoolExpr("!(f)"))
	fmt.Println(parseBoolExpr("|(f, t)"))
	fmt.Println(parseBoolExpr("&(t,f)"))
	fmt.Println(parseBoolExpr("|(&(t,f,t),!(t))"))
}
