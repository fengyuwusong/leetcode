package easy

import "testing"

/**
Given a non-empty array of digits representing a non-negative integer, plus one to the integer.

The digits are stored such that the most significant digit is at the head of the queue, and each element in the array contain a single digit.

You may assume the integer does not contain any leading zero, except the number 0 itself.

Example 1:

Input: [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Example 2:

Input: [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
66
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/plus-one
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] == 10 {
			digits[i] = 0
			continue
		}
		break
	}
	// 开头为0 需进位
	if digits[0] == 0 {
		digits = append(digits, 0)
		digits[0] = 1
	}
	return digits
}

func TestPlusOne(t *testing.T) {
	for _, i := range plusOne([]int{1, 2, 3}) {
		println(i)
	}
	for _, i := range plusOne([]int{9}) {
		println(i)
	}
}
