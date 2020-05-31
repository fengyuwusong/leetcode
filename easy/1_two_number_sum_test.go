package easy

import (
	"fmt"
	"testing"
)

//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
//示例:
//给定 nums = [2, 7, 11, 15], target = 9
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/two-sum
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func twoSum(nums []int, target int) []int {
	// 以 value, index 建立哈希表
	// 初始化时建立可避免重复值问题 如 [3, 3] 6
	hashMap := make(map[int]int)
	// 循环哈希表查看是否存在对应索引
	for index, value := range nums {
		otherKey := target - value
		if otherIndex, has := hashMap[otherKey]; has {
			return []int{otherIndex, index}
		}
		hashMap[value] = index
	}
	return []int{}
}

func Test_twoSum(t *testing.T) {

	fmt.Printf("%v\n", twoSum([]int{3, 3}, 6))
}
