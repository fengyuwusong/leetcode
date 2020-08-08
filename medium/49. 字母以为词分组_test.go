package medium

import (
	"fmt"
	"sort"
	"testing"
)

/**
给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

示例:

输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
说明：

所有输入均为小写字母。
不考虑答案输出的顺序。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/group-anagrams
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 可使用质数映射a-z相乘 相同字符得到结果也是唯一的key值
// 根据string从小到大排序生成key
func getMapKey(s string) string {
	r := []int32(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func groupAnagrams(strs []string) [][]string {
	hashMap := make(map[string][]string)
	var ans [][]string
	for _, v := range strs {
		hashMap[getMapKey(v)] = append(hashMap[getMapKey(v)], v)
	}
	for _, v := range hashMap {
		ans = append(ans, v)
	}
	return ans
}
func TestGroupAnagrams(t *testing.T) {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
