package main

import (
	"fmt"
	"math"
)

func bestCoordinate(towers [][]int, radius int) []int {
	size := 50
	max := 0
	ans := make([]int, 2)
	for i := 0; i <= size; i++ {
		for j := 0; j <= size; j++ {
			var sum int
			for _, tower := range towers {
				dx := tower[0] - i
				dy := tower[1] - j
				d := math.Sqrt(float64(dx*dx + dy*dy))
				if d <= float64(radius) {
					sum += int(float64(tower[2]) / (1 + d))
				}
			}
			if sum > max {
				max = sum
				ans[0] = i
				ans[1] = j
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(bestCoordinate([][]int{{0, 1, 2}, {2, 1, 2}, {1, 0, 2}, {1, 2, 2}}, 1))
}
