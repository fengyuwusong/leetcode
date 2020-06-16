package medium

import (
	"fmt"
	"strconv"
	"testing"
)

/**
给定一个经过编码的字符串，返回它解码后的字符串。

编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。



示例 1：

输入：s = "3[a]2[bc]"
输出："aaabcbc"
示例 2：

输入：s = "3[a2[c]]"
输出："accaccacc"
示例 3：

输入：s = "2[abc]3[cd]ef"
输出："abcabccdcdcdef"
示例 4：

输入：s = "abc3[cd]xyz"
输出："abccdcdcdxyz"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/decode-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 解法1 使用栈 dfs
// 问题点 当数值为双位数以上时需做判断兼容

// 解法2 使用两个栈 一个数字栈 一个字母栈
func decodeString(s string) string {
	var ss, stack []rune
	for _, v := range s {
		stack = append(stack, v)
		if stack[len(stack)-1] == ']' {
			stack = stack[:len(stack)-1]
			// 出栈直至栈顶为 '['
			for stack[len(stack)-1] != '[' {
				ss = append([]rune{stack[len(stack)-1]}, ss...)
				stack = stack[:len(stack)-1]
			}
			// 继续出栈获取前置数字
			i := len(stack) - 2
			for ; i >= 0; i-- {
				if ok := isNum(stack[i]); !ok {
					break
				}
			}
			num, _ := strconv.ParseInt(string(stack[i+1:len(stack)-1]), 10, 64)
			stack = stack[:i+1]
			// 将结果再次进栈
			for ; num > 0; num-- {
				stack = append(stack, ss...)
			}
			ss = []rune{}
		}
	}
	return string(stack)
}

func isNum(r rune) bool {
	val := r - '0'
	if val >= 0 && val < 10 {
		return true
	}
	return false
}

func TestDecodeString(t *testing.T) {
	fmt.Println(decodeString("100[ac]"))
	fmt.Println(decodeString("2[a1[c]]"))
	fmt.Println(decodeString("3[a]2[bc]"))
	fmt.Println(decodeString("3[a2[c]]"))
	fmt.Println(decodeString("2[abc]3[cd]ef"))
	fmt.Println(decodeString("abc3[cd]xyz"))
}
