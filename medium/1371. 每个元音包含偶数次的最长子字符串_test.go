package medium

import (
	"fmt"
	"testing"
)

/**
给你一个字符串 s ，请你返回满足以下条件的最长子字符串的长度：每个元音字母，即 'a'，'e'，'i'，'o'，'u' ，在子字符串中都恰好出现了偶数次。



示例 1：

输入：s = "eleetminicoworoep"
输出：13
解释：最长子字符串是 "leetminicowor" ，它包含 e，i，o 各 2 个，以及 0 个 a，u 。
示例 2：

输入：s = "leetcodeisgreat"
输出：5
解释：最长子字符串是 "leetc" ，其中包含 2 个 e 。
示例 3：

输入：s = "bcbcbc"
输出：6
解释：这个示例中，字符串 "bcbcbc" 本身就是最长的，因为所有的元音 a，e，i，o，u 都出现了 0 次。


提示：

1 <= s.length <= 5 x 10^5
s 只包含小写英文字母。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-the-longest-substring-containing-vowels-in-even-counts
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func findTheLongestSubstring(s string) int {

	// 哈希表 hashData[i] = j  =>
	// 索引 i[0=>(00000)2, 1<<5=>(10000)2] 每一位分别表示 a e i o u 五个字符 如 a => 2^0 => 1 e => 2^1 => 10 ...
	// 索引对应状态维护为: 0=>偶数 1=>基数
	// j 表示当前遍历s的位置
	hashData := make([]int, 1<<5)

	max, status := 0, 0 // max 记录最大值 status 记录到遍历i节点为止的元音奇偶状态
	// 初始状态为-1 即未开始遍历
	for i, _ := range hashData {
		hashData[i] = -1
	}
	// i=>0 表示没有元音字符 当没选择字符时 j=>0 即还没开始遍历
	hashData[0] = 0
	// 遍历数组
	for i, ss := range s {
		switch ss {
		case 'a':
			// 计算当前状态 使用异或操作刚好对应奇偶数相同抵消的规则 奇偶状态不同则取偶 不同取奇
			// 例如 在 i 之前的status状态为 10001 a、u出现了奇数次 那么与 00001 a 做异或操作
			// 位相同取0 不通取1 得出结果 10000 即加上s[i]后 a的状态是偶数 0
			status ^= 1 << 0
		case 'e':
			status ^= 1 << 1
		case 'i':
			status ^= 1 << 2
		case 'o':
			status ^= 1 << 3
		case 'u':
			status ^= 1 << 4
		}
		// 查看该status状态下是否有值 存在值说明可以两两抵消, 元音个数之和为偶数 有则取最大值
		if hashData[status] != -1 {
			// i+1 - hashData[status] 当前位置 - 上一个相同状态位置 = 符合元音偶数的长度
			if max < i+1-hashData[status] {
				max = i + 1 - hashData[status]
			}
		} else {
			// 否则更新当前状态下的值 注意当前位置为 i+1
			// 此处else才进行 否则会覆盖符合状态的最小值 无法计算出最大长度
			hashData[status] = i + 1
		}
	}
	return max
}

func TestFindTheLongestSubstring(t *testing.T) {
	fmt.Println(findTheLongestSubstring("eleetminicoworoep"))
}
