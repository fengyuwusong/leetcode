package 剑指offer

import (
	"fmt"
	"testing"
)

/**
在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。

示例:

s = "abaccdeff"
返回 "b"

s = ""
返回 " "


限制：

0 <= s 的长度 <= 50000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
// 解法1 双循环遍历 时间复杂度O(n^2) 空间复杂度O(1)
// 解法2 hash 遍历两次s 第一次记录hash状态 第二次遍历得到hash中哪个字符没重复出现 时间复杂度O(n) 空间复杂度O(n)
// 解法3 优化解法2 哈希表可使用有序哈希表(键值对是基于插入顺序排序) 第二次遍历可遍历hash 当重复字符很多时, 可减少遍历得元素个数
// 由于go原生map无序 无法获取第一个出现 故需额外使用空间记录字符出现顺序
func firstUniqChar(s string) byte {
	hashMap := make(map[rune]bool)
	for _, ss := range s {
		if _, ok := hashMap[ss]; !ok {
			hashMap[ss] = true
		} else {
			hashMap[ss] = false
		}
	}

	for _, ss := range s {
		if hashMap[ss] {
			return byte(ss)
		}
	}
	return ' '
}

// 更优解法 使用数组代替map
func firstUniqChar2(s string) byte {
	var words [255]int
	length := len(s)
	for i := 0; i < length; i++ {
		words[int(s[i])]++
	}
	for i := 0; i < length; i++ {
		if words[int(s[i])] == 1 {
			return s[i]
		}
	}
	return ' '
}

func TestFirstUniqChar(t *testing.T) {
	fmt.Println(firstUniqChar("abaccdeff") == 'b')
	fmt.Println(firstUniqChar("") == ' ')
}
