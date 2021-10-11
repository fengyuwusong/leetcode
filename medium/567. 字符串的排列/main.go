package main

import "fmt"

/**
给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。

换句话说，s1 的排列之一是 s2 的 子串 。



示例 1：

输入：s1 = "ab" s2 = "eidbaooo"
输出：true
解释：s2 包含 s1 的排列之一 ("ba").
示例 2：

输入：s1= "ab" s2 = "eidboaoo"
输出：false


提示：

1 <= s1.length, s2.length <= 104
s1 和 s2 仅包含小写字母

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutation-in-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 滑动窗口 + hash法判断是否所有s2中是否存在s1的所有字符即可
func checkInclusion(s1 string, s2 string) bool {
	hash := make([]int, 256)
	// 先构建s1的hash表
	for _, s := range s1 {
		hash[s]++
	}
	for i, j := 0, 0; i < len(s2) && j < len(s2); {
		// j指针向前走 按照hash记录的字符查找s1是否均存在
		for ; j < len(s2) && hash[s2[j]] > 0; j++ {
			hash[s2[j]]--
		}
		// 如长度等于s1则说明存在s1的所有字符
		if j-i == len(s1) {
			return true
		}

		// i指针往前走 开始下一次遍历
		hash[s2[i]]++
		i++
	}
	return false
}

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
	fmt.Println(checkInclusion("ab", "eidboaoo"))
}
