package easy

import (
	"fmt"
	"strings"
	"testing"
)

/**
给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。

示例 1:

输入: "Let's take LeetCode contest"
输出: "s'teL ekat edoCteeL tsetnoc"
注意：在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-words-in-a-string-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func reverseWords(s string) string {
	var ans string
	ss := strings.Split(s, " ")
	for _, v := range ss {
		nv := []rune(v)
		for i := 0; i < len(nv)/2; i++ {
			nv[i], nv[len(nv)-i-1] = nv[len(nv)-i-1], nv[i]
		}
		ans += " " + string(nv)
	}
	return ans[1:]
}

func TestReverseWords(t *testing.T) {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}
