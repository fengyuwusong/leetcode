package main

import (
	"container/list"
)

func isValid(s string) bool {
	l := list.New()
	for _, v := range s {
		switch v {
		case '(', '[', '{':
			l.PushBack(v)
		case ')', ']', '}':
			if l.Len() == 0 {
				return false
			}
			switch v {
			case ')':
				if l.Back().Value.(rune) != '(' {
					return false
				}
			case ']':
				if l.Back().Value.(rune) != '[' {
					return false
				}
			case '}':
				if l.Back().Value.(rune) != '{' {
					return false
				}
			}
			l.Remove(l.Back())
		}
	}
	if l.Len() == 0 {
		return true
	}
	return false
}

func main() {
	println(isValid("()"))
	println(isValid("()[]{}"))
	println(isValid("(]"))
	println(isValid("([)]"))
	println(isValid("{[]}"))
}
