package main

import "fmt"

// 假设有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花不能种植在相邻的地块上，它们会争夺水源，两者都会死去。
//
// 给你一个整数数组 flowerbed 表示花坛，由若干 0 和 1 组成，其中 0 表示没种植花，1 表示种植了花。另有一个数 n ，能否在不打破种植规则
// 的情况下种入 n 朵花？能则返回 true ，不能则返回 false 。
//
//
//
// 示例 1：
//
//
// 输入：flowerbed = [1,0,0,0,1], n = 1
// 输出：true
//
//
// 示例 2：
//
//
// 输入：flowerbed = [1,0,0,0,1], n = 2
// 输出：false
//
//
//
//
// 提示：
//
//
// 1 <= flowerbed.length <= 2 * 10⁴
// flowerbed[i] 为 0 或 1
// flowerbed 中不存在相邻的两朵花
// 0 <= n <= flowerbed.length
//
//
// Related Topics 贪心 数组 👍 680 👎 0

// author: fengyuwusong
// create time: 2023-10-23 18:42:57
// leetcode submit region begin(Prohibit modification and deletion)
func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	var add = 1
	for _, v := range flowerbed {
		if v == 0 {
			add++
		} else {
			add = 0
		}
		if add == 3 {
			n--
			add = 1
		}
		if n <= 0 {
			return true
		}
	}
	if add+1 == 3 {
		n--
	}
	return n <= 0
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 0, 1}, 2))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1, 0, 0}, 2))
}
