package main

func maxArea(height []int) int {
	var res int
	i, j := 0, len(height)-1
	for i < j {
		// 计算接水面积
		area := min(height[i], height[j]) * (j - i)
		res = max(res, area)
		// 更新指针
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return res
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	println(maxArea([]int{1, 2, 4, 3}))
}
