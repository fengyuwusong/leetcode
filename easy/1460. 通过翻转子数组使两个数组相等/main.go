package main

import "fmt"

//给你两个长度相同的整数数组 target 和 arr 。每一步中，你可以选择 arr 的任意 非空子数组 并将它翻转。你可以执行此过程任意次。
//
// 如果你能让 arr 变得与 target 相同，返回 True；否则，返回 False 。
//
//
//
// 示例 1：
//
//
//输入：target = [1,2,3,4], arr = [2,4,1,3]
//输出：true
//解释：你可以按照如下步骤使 arr 变成 target：
//1- 翻转子数组 [2,4,1] ，arr 变成 [1,4,2,3]
//2- 翻转子数组 [4,2] ，arr 变成 [1,2,4,3]
//3- 翻转子数组 [4,3] ，arr 变成 [1,2,3,4]
//上述方法并不是唯一的，还存在多种将 arr 变成 target 的方法。
//
//
// 示例 2：
//
//
//输入：target = [7], arr = [7]
//输出：true
//解释：arr 不需要做任何翻转已经与 target 相等。
//
//
// 示例 3：
//
//
//输入：target = [3,7,9], arr = [3,7,11]
//输出：false
//解释：arr 没有数字 9 ，所以无论如何也无法变成 target 。
//
//
//
//
// 提示：
//
//
// target.length == arr.length
// 1 <= target.length <= 1000
// 1 <= target[i] <= 1000
// 1 <= arr[i] <= 1000
//
// Related Topics数组 | 哈希表 | 排序
//
// 👍 37, 👎 0
//
//
//
//

// fengyuwusong 2022-08-24 00:44:59
//leetcode submit region begin(Prohibit modification and deletion)
// 由于能通过翻转的方式使数组每个元素互换位置 故元素出现个数一样 则说明能通过翻转的方式使其相等
func canBeEqual(target []int, arr []int) bool {
	dict := make(map[int]int)
	for i := 0; i < len(target); i++ {
		dict[target[i]]++
		dict[arr[i]]--
	}
	for _, v := range dict {
		if v != 0 {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(canBeEqual([]int{3, 4, 5}, []int{3, 5, 4}))
	fmt.Println(canBeEqual([]int{3, 4, 511}, []int{3, 5, 4}))
}
