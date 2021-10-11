package main

import "fmt"

/**
给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。

 

示例：

输入："Let's take LeetCode contest"
输出："s'teL ekat edoCteeL tsetnoc"
 

提示：

在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-words-in-a-string-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func reverseWords(s string) string {
	bs := []byte(s)
	j := 0
	for i, v := range bs {
		if v == ' '{
			// 记录下次j起始位置
			temp := i+1
			// 退回空格前最后一个字母
			i--
			for ; j<i;j, i= j+1, i-1{
				bs[i], bs[j] = bs[j], bs[i]
			}
			j = temp
		}
	}

	// 最后一个单词也需要反转
	for i:= len(bs) -1;j<i;i,j=i-1, j+1 {
		bs[i], bs[j] = bs[j], bs[i]
	}
	return string(bs)
}

func main() {
	s := "i am jone"
	fmt.Println(reverseWords(s))
}
