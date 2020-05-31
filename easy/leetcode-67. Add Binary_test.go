package easy

import (
	"testing"
)

/**
Given two binary strings, return their sum (also a binary string).

The input strings are both non-empty and contains only characters 1 or 0.

Example 1:

Input: a = "11", b = "1"
Output: "100"
Example 2:

Input: a = "1010", b = "1011"
Output: "10101"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-binary
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func addBinary(a string, b string) string {
	al := len(a) - 1
	bl := len(b) - 1
	needAdd := 0
	var res string
	for ; al >= 0 || bl >= 0; al, bl = al-1, bl-1 {
		if al >= 0 && bl >= 0 {
			aRes := rune2int(a[al])
			bRes := rune2int(b[bl])
			na, tRes := cal(aRes + bRes + needAdd)
			needAdd = na
			res = tRes + res
			continue
		}
		if al < 0 {
			bRes := rune2int(b[bl])
			na, tRes := cal(bRes + needAdd)
			needAdd = na
			res = tRes + res
			continue
		}
		aRes := rune2int(a[al])
		na, tRes := cal(aRes + needAdd)
		needAdd = na
		res = tRes + res
	}
	if needAdd == 1 {
		res = "1" + res
	}
	return res
}

func rune2int(r uint8) int {
	if r == '1' {
		return 1
	}
	return 0
}

func cal(i int) (int, string) {
	needAdd := 0
	res := "0"
	switch i {
	case 0:
	case 1:
		res = "1"
	case 2:
		needAdd = 1
	case 3:
		needAdd = 1
		res = "1"
	}
	return needAdd, res
}

func TestAddBinary(t *testing.T) {
	println(addBinary("11", "1"))
	println(addBinary("1010", "10110"))
}
