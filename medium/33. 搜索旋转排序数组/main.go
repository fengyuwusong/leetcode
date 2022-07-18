package main

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		// [l, mid]有序
		if nums[0] <= nums[mid] {
			// target处于[l, mid)
			if nums[0] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			//[mid+1, r]有序
			// target 处于 (mid, r]
			if nums[mid] < target && target <= nums[len(nums)-1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}

func main() {
	println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	println(search([]int{1}, 0))
	println(search([]int{1, 3}, 0))
	println(search([]int{5, 1, 3}, 3))
	println(search([]int{5, 1, 3}, 5))
	println(search([]int{3, 1}, 1))
}
