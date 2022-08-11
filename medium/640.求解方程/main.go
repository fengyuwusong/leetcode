package main

import (
	"fmt"
	"strconv"
	"strings"
)

// author: fengyuwusong date: 2022-08-10 16:27:31

// 640. 求解方程
//求解一个给定的方程，将x以字符串 "x=#value" 的形式返回。该方程仅包含 '+' ， '-' 操作，变量 x 和其对应系数。
//
// 如果方程没有解，请返回 "No solution" 。如果方程有无限解，则返回 “Infinite solutions” 。
//
// 题目保证，如果方程中只有一个解，则 'x' 的值是一个整数。
//
//
//
// 示例 1：
//
//
//输入: equation = "x+5-3+x=6+x-2"
//输出: "x=2"
//
//
// 示例 2:
//
//
//输入: equation = "x=x"
//输出: "Infinite solutions"
//
//
// 示例 3:
//
//
//输入: equation = "2x=x"
//输出: "x=0"
//
//
//
//
// 提示:
//
//
// 3 <= equation.length <= 1000
// equation 只有一个 '='.
// equation 方程由整数组成，其绝对值在 [0, 100] 范围内，不含前导零和变量 'x' 。
//
// Related Topics数学 | 字符串 | 模拟
//
// 👍 152, 👎 0
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
func getCol(r rune) int {
	switch r {
	case 'x':
		return 0
	case '+', '-':
		return 1
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return 2
	default:
		return 3
	}
}
func solveEquation(equation string) string {
	expressions := strings.Split(equation, "=")
	if len(expressions) != 2 {
		return "No solution"
	}
	lX, lConst, err := getXAndConst(expressions[0])
	if err != nil {
		return "No solution"
	}
	rX, rConst, err := getXAndConst(expressions[1])
	if err != nil {
		return "No solution"
	}
	if lX == rX && lConst == rConst {
		return "Infinite solutions"
	} else if lX == rX && lConst != rConst {
		return "No solution"
	} else {
		x := (rConst - lConst) / (lX - rX)
		return "x=" + strconv.Itoa(x)
	}
}

func getXAndConst(equation string) (x int, constNum int, err error) {
	dict := map[string][]string{
		"begin": {"x", "sign", "nums", "error"},
		"nums":  {"x", "sign", "nums", "other"},
		"x":     {"error", "sign", "error", "other"},
		"sign":  {"x", "error", "nums", "other"},
		"other": {"error", "error", "error", "error"},
	}
	state := "begin"
	temp := 0
	tempSet := false
	negative := false
	for _, c := range equation {
		state = dict[state][getCol(c)]
		switch state {
		case "nums":
			temp = temp*10 + int(c-'0')
			tempSet = true
		case "sign":
			if negative {
				constNum -= temp
			} else {
				constNum += temp
			}
			if c == '-' {
				negative = true
			} else {
				negative = false
			}
			temp = 0
			tempSet = false
		case "x":
			if !tempSet && temp == 0 {
				temp = 1
			}
			if negative {
				x -= temp
			} else {
				x += temp
			}
			tempSet = false
			negative = false
			temp = 0
		default:
			return 0, 0, fmt.Errorf("error")
		}
	}
	if temp != 0 {
		if negative {
			constNum -= temp
		} else {
			constNum += temp
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(solveEquation("x+5-3+x=6+x-2"))
	fmt.Println(solveEquation("x=x"))
	fmt.Println(solveEquation("2x=x"))
	fmt.Println(solveEquation("-x=6"))
	fmt.Println(solveEquation("0x=0"))
	fmt.Println(solveEquation("0x=6"))
}
