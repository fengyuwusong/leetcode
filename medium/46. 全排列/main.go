package main

import "fmt"

func permute(nums []int) [][]int {
	var res [][]int
	var trackBack func(nums, track []int, used []bool)
	trackBack = func(nums, track []int, used []bool) {
		if len(track) == len(nums) {
			// 此处注意go数组为引用传递 不能直接使用原参数 需拷贝
			p := make([]int, len(track))
			copy(p, track)
			res = append(res, p)
			return
		}
		for i := 0; i < len(nums); i++ {
			// 是否已选择
			if used[i] {
				continue
			}
			// 选择
			track = append(track, nums[i])
			// 标记已选择
			used[i] = true
			// 递归下一层
			trackBack(nums, track, used)
			// 撤销选择
			track = track[:len(track)-1]
			used[i] = false
		}
	}
	trackBack(nums, []int{}, make([]bool, len(nums)))
	return res
}

func main() {
	for _, n := range permute([]int{1, 2, 3}) {
		fmt.Println(n)
	}
	for _, n := range permute([]int{5, 4, 6, 2}) {
		fmt.Println(n)
	}
}
