package main

/**
你是产品经理，目前正在带领一个团队开发新的产品。不幸的是，你的产品的最新版本没有通过质量检测。由于每个版本都是基于之前的版本开发的，所以错误的版本之后的所有版本都是错的。

假设你有 n 个版本 [1, 2, ..., n]，你想找出导致之后所有版本出错的第一个错误的版本。

你可以通过调用 bool isBadVersion(version) 接口来判断版本号 version 是否在单元测试中出错。实现一个函数来查找第一个错误的版本。你应该尽量减少对调用 API 的次数。

 
示例 1：

输入：n = 5, bad = 4
输出：4
解释：
调用 isBadVersion(3) -> false
调用 isBadVersion(5) -> true
调用 isBadVersion(4) -> true
所以，4 是第一个错误的版本。
示例 2：

输入：n = 1, bad = 1
输出：1
 

提示：

1 <= bad <= n <= 231 - 1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/first-bad-version
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

// 二分查找
func firstBadVersion(n int) int {
	min := 1
	max := n
	for min < max { // 循环到左右指针相同
		// 这样计算中间值存在溢出风险
		//mid := (min + max ) / 2
		mid := min + (max - min) / 2
		if isBadVersion(mid) {
			max = mid // 答案在 [min, mid] 中
		}else {
			min = mid + 1 // 答案在 [mid + 1, max] 中
		}
	}
	// 当min == max 时, 则为正确答案
	return min
}

func isBadVersion(n int) bool {
	if n >= 4 {
		return true
	}
	return false
}

func main() {
	print(firstBadVersion(5))
}