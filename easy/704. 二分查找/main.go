package main

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		}
	}
	return -1
}

func main() {
	println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
	println(search([]int{-1, 0, 3, 5, 9, 12}, 2))
}
