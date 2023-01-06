package main

import "fmt"

//给你一个正整数 num ，请你统计并返回 小于或等于 num 且各位数字之和为 偶数 的正整数的数目。
//
// 正整数的 各位数字之和 是其所有位上的对应数字相加的结果。
//
//
//
// 示例 1：
//
//
//输入：num = 4
//输出：2
//解释：
//只有 2 和 4 满足小于等于 4 且各位数字之和为偶数。
//
//
// 示例 2：
//
//
//输入：num = 30
//输出：14
//解释：
//只有 14 个整数满足小于等于 30 且各位数字之和为偶数，分别是：
//2、4、6、8、11、13、15、17、19、20、22、24、26 和 28 。
//
//
//
//
// 提示：
//
//
// 1 <= num <= 1000
//
//
// Related Topics 数学 模拟 👍 66 👎 0

// author: fengyuwusong
// create time: 2023-01-06 23:25:07
// leetcode submit region begin(Prohibit modification and deletion)
func countEven(num int) int {
	var ans int
	for i := 2; i <= num; i++ {
		if checkSum(i) {
			ans++
		}
	}
	return ans
}

func checkSum(num int) bool {
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum%2 == 0
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(countEven(4))
	fmt.Println(countEven(30))
}
