package easy

/**
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。



示例：

s = "leetcode"
返回 0

s = "loveleetcode"
返回 2


提示：你可以假定该字符串只包含小写字母。
*/
func firstUniqChar(s string) int {
	var words [255]int
	length := len(s)
	for i := 0; i < length; i++ {
		words[int(s[i])]++
	}
	for i := 0; i < length; i++ {
		if words[int(s[i])] == 1 {
			return i
		}
	}
	return -1
}
