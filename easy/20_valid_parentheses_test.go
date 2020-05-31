package easy

import (
	"testing"
)

/**
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

Example 1:
Input: "()"
Output: true

Example 2:
Input: "()[]{}"
Output: true

Example 3:
Input: "(]"
Output: false

Example 4:
Input: "([)]"
Output: false

Example 5:
Input: "{[]}"
Output: true
*/

// 使用栈的思路进行解答
func isValid(s string) bool {
	var stack []string // 自定义栈
	for _, str := range s {
		switch str {
		case '(':
			fallthrough
		case '{':
			fallthrough
		case '[':
			// 入栈
			stack = append(stack, string(str))
			break
		case ')':
			// 查看是否相等并出栈
			if len(stack) == 0 {
				return false
			}
			ch := stack[len(stack)-1]
			if ch == "(" {
				// 出栈
				if len(stack) > 1 {
					stack = stack[0 : len(stack)-1]
				} else {
					stack = stack[0:0]
				}
				break
			}
			return false
		case '}':
			// 查看是否相等并出栈
			if len(stack) == 0 {
				return false
			}
			ch := stack[len(stack)-1]
			if ch == "{" {
				// 出栈
				if len(stack) > 1 {
					stack = stack[0 : len(stack)-1]
				} else {
					stack = stack[0:0]
				}
				break
			}
			return false
		case ']':
			// 查看是否相等并出栈
			if len(stack) == 0 {
				return false
			}
			ch := stack[len(stack)-1]
			if ch == "[" {
				// 出栈
				if len(stack) > 1 {
					stack = stack[0 : len(stack)-1]
				} else {
					stack = stack[0:0]
				}
				break
			}
			return false
		}
	}
	return len(stack) == 0
}

func TestIsValid(t *testing.T) {
	println(isValid("[{{}}]"))
}
