package medium

import (
	"fmt"
	"testing"
)

// dfs
func canVisitAllRooms(rooms [][]int) bool {
	if len(rooms) == 0 {
		return true
	}
	visit := make([]bool, len(rooms))
	stack := []int{0}
	for len(stack) != 0 {
		roomNum := stack[0]
		stack = stack[1:]
		visit[roomNum] = true
		for _, rn := range rooms[roomNum] {
			if !visit[rn] {
				stack = append(stack, rn)
			}
		}
	}
	for _, v := range visit {
		if v == false {
			return false
		}
	}
	return true
}
func TestCanVisitAllRooms(t *testing.T) {
	fmt.Println(canVisitAllRooms([][]int{{1}, {2}, {3}, {}}))
	fmt.Println(canVisitAllRooms([][]int{{1, 3}, {3, 0, 1}, {2}, {0}}))
}
