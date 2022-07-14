package sort

import (
	"fmt"
	"testing"
)

// Bubble 冒泡排序
func Bubble(arr []int) []int {
	for i := len(arr) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func TestBubble(t *testing.T) {
	fmt.Printf("%v\n", Bubble([]int{3, 2, 5, 1, 7, 4}))
}
