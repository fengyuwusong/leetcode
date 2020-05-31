package medium

import (
	"fmt"
	"sort"
	"testing"
)

// 暴力解
//func threeSum(nums []int) [][]int {
//	sort.Ints(nums)
//	var x, y, z int
//	var res [][]int
//	for i := 0; i < len(nums); i++ {
//		if i != 0 && x == nums[i] {
//			continue
//		}
//		x = nums[i]
//		for j := i + 1; j < len(nums); j++ {
//			if j > i+1 && nums[j] == nums[j-1] {
//				continue
//			}
//			y = nums[j]
//			z = 0 - x - y
//			for k := j + 1; k < len(nums); k++ {
//				if z == nums[k] {
//					res = append(res, []int{x, y, z})
//					break
//				}
//			}
//		}
//	}
//	return res
//}

//// hash表解法 不可ac 特殊情况未编写 如重复值
//func threeSum(nums []int) [][]int {
//	var res [][]int
//	hash := make(map[int][]int)
//	for i := 0; i < len(nums)-2; i++ {
//		for j := i + 1; j < len(nums)-1; j++ {
//			if hash[nums[j]] != nil {
//				res = append(res, append(hash[nums[j]], nums[j]))
//				continue
//			}
//			// hash表不存在则两个人x, y 组队
//			hash[0-nums[j]-nums[i]] = []int{nums[j], nums[i]}
//		}
//	}
//	return res
//}

// 双指针
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums)-1; i++ {
		// x 值不可能大于0
		if nums[i] > 0 {
			break
		}
		left := i + 1
		right := len(nums) - 1
		for {
			// left >= right  x z 不可能符号相同 时退出
			if left >= right || nums[i]*nums[right] > 0 {
				break
			}
			data := nums[i] + nums[left] + nums[right]
			if data == 0 {
				res = append(res, []int{nums[left], nums[i], nums[right]})
				// 同一x值可能有不同结果 left + 1取其他情况
				left++
				for left < right && nums[left] == nums[left-1] {
					left++
				}
			} else if data > 0 {
				right--
				for right > left && nums[right] == nums[right+1] {
					right--
				}
			} else {
				left++
				for left < right && nums[left] == nums[left-1] {
					left++
				}
			}
		}
		//去除x值重复值情况
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}

func TestThreeSum(t *testing.T) {
	//fmt.Printf("%v\n", threeSum([]int{-1, 0, 1, 2, -1, -4}))
	//fmt.Printf("%v\n", threeSum([]int{0, 0, 0, 0}))
	//fmt.Printf("%v\n", threeSum([]int{1, 0, -1}))
	//fmt.Printf("%v\n", threeSum([]int{-2, 0, 1, 1, 2}))
	fmt.Printf("%v\n", threeSum([]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}))
}
