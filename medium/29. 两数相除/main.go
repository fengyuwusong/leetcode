package main

import (
	"fmt"
	"math"
)

/**
给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。

返回被除数 dividend 除以除数 divisor 得到的商。

整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2



示例 1:

输入: dividend = 10, divisor = 3
输出: 3
解释: 10/3 = truncate(3.33333..) = truncate(3) = 3
示例 2:

输入: dividend = 7, divisor = -3
输出: -2
解释: 7/-3 = truncate(-2.33333..) = -2


提示：

被除数和除数均为 32 位有符号整数。
除数不为 0。
假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231,  231 − 1]。本题中，如果除法结果溢出，则返回 231 − 1。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/divide-two-integers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 解法 被除数 > 除数 * 商 => (除数+除数+xx)  > 除数  则从除数和被除数里面一直加即可找到对应的商
// 例如： 60/8 = (60-32) / 8 + 4 = (60-32-16)/8+4+2 = (60-32-16-8)/8 +4+2+1 = 7
// 被除数和除数相同符号，则商为正数 防止为负数
func divide(dividend int, divisor int) int {
	switch divisor {
	case 0:
		return 0
	case 1:
		return dividend
	case -1:
		//当除数为-1且被除数为Integer.MIN_VALUE时，将会溢出，返回Integer.MAX_VALUE
		if dividend > math.MinInt32 {
			return -dividend
		}
		return math.MaxInt32
	default:
		// 把被除数与除数调整为正数,为防止被除数Integer.MIN_VALUE转换为正数会溢出
		if dividend < 0 && divisor < 0 {
			return _divide(-int64(dividend), -int64(divisor))
		} else if dividend < 0 || divisor < 0 {
			// 除数与被输出符号不一致 使用绝对值运算 商为负数
			return -_divide(int64(math.Abs(float64(dividend))), int64(math.Abs(float64(divisor))))
		} else {
			return _divide(int64(dividend), int64(divisor))
		}
	}
}

func _divide(dividend, divisor int64) int {
	// 被除数小于除数 结果等于0
	if dividend < divisor {
		return 0
	}
	var sum int64 = divisor // 记录用了count个divisor的和
	count := 1              // 使用了多少个divisor
	for dividend >= sum {
		// 每次翻倍
		sum <<= 1
		count <<= 1
	}

	// 此时dividend < sum 回退一次
	sum >>= 1
	count >>= 1

	return count + _divide(dividend-sum, divisor)
}

func main() {
	fmt.Println(divide(10, 3))
}
