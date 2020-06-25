package easy

import (
	"fmt"
	"testing"
)

/**
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-common-prefix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 解法1 横向扫描 以第一个str为模板 依次遍历strs 得到公共前缀
// 时间复杂度O(nm) 空间复杂度O(1)
// 解法2 纵向扫描 以每一个字符为模板 依次遍历每个strs每个串的每个字符得到公共前缀
// 时间复杂度O(nm) 空间复杂度O(1)
// 解法3 分治法 以第一种解法为基础，我们可以将strs分为两部分
// 然后依次以第一种解法得到两部分的解在合并
// 时间复杂度T(n) = 2*T(n/2) + O(m) => O(nm)
// 空间复杂度O(mlogn)空间复杂度主要取决于递归调用的层数，层数最大为logn，每层需要m的空间存储返回结果
// 层数最大为logn，每层需要 m 的空间存储返回结果。
// 解法4 二分查找 查找字符中的公共前缀并不需要从头到尾进行遍历，可以使用二分查找的方法。
// 先判断[0, minLen) 中mid = midLen/2 是否相等，定位最后一个相等字符的位置。
// 时间复杂度O(mnlogm) 空间复杂度O(1)
// 解法3
func longestCommonPrefix3(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var lcp func(int, int) string
	lcp = func(start int, end int) string {
		// 递归结束条件
		if start == end {
			return strs[start]
		}
		// 分别计算左右两部分的公共前缀
		mid := (start + end) / 2
		lcpLeft, lcpRight := lcp(start, mid), lcp(mid+1, end)
		// 取左右两部分公共前缀的最长前缀
		minLen := min(len(lcpLeft), len(lcpRight))
		for i := 0; i < minLen; i++ {
			if lcpLeft[i] != lcpRight[i] {
				return lcpLeft[:i]
			}
		}
		return lcpLeft[:minLen]
	}
	return lcp(0, len(strs)-1)
}

// 解法4
func longestCommonPrefix4(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	// 判断当前索引前字符是否服务公共前缀
	isCommonPrefix := func(index int) bool {
		str0, count := strs[0][:index], len(strs)
		for i := 1; i < count; i++ {
			if strs[i][:index] != str0 {
				return false
			}
		}
		return true
	}
	// 取得最小长度字符串长度
	minLen := len(strs[0])
	for _, s := range strs {
		if len(s) < minLen {
			minLen = len(s)
		}
	}
	startIndex, endIndex := 0, minLen
	// 开始二分查找
	for startIndex < endIndex {
		midIndex := (endIndex-startIndex+1)/2 + startIndex
		if isCommonPrefix(midIndex) {
			startIndex = midIndex
		} else {
			endIndex = midIndex - 1
		}
	}
	return strs[0][:endIndex]
}
func TestLongestCommonPrefix(t *testing.T) {
	fmt.Println(longestCommonPrefix3([]string{"123", "1211", "1233", "1111"}))
	fmt.Println(longestCommonPrefix4([]string{"123", "1211", "1233", "1111"}))
}
