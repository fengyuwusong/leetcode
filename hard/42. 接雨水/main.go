package main

func trap(height []int) int {
	left, right := 0, len(height)-1
	lMax, rMax := 0, 0
	res := 0
	for left < right {
		lMax = max(lMax, height[left])
		rMax = max(rMax, height[right])

		// 说明最低在左边 加左边最大值-当前的差值 左指针+1 否则反之
		if lMax < rMax {
			res += lMax - height[left]
			left++
		} else {
			res += rMax - height[right]
			right--
		}
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
