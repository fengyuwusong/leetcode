package easy

import (
	"fmt"
	"testing"
)

/**
有一幅以二维整数数组表示的图画，每一个整数表示该图画的像素值大小，数值在 0 到 65535 之间。

给你一个坐标 (sr, sc) 表示图像渲染开始的像素值（行 ，列）和一个新的颜色值 newColor，让你重新上色这幅图像。

为了完成上色工作，从初始坐标开始，记录初始坐标的上下左右四个方向上像素值与初始坐标相同的相连像素点，接着再记录这四个方向上符合条件的像素点与他们对应四个方向上像素值与初始坐标相同的相连像素点，……，重复该过程。将所有有记录的像素点的颜色值改为新的颜色值。

最后返回经过上色渲染后的图像。

示例 1:

输入:
image = [[1,1,1],[1,1,0],[1,0,1]]
sr = 1, sc = 1, newColor = 2
输出: [[2,2,2],[2,2,0],[2,0,1]]
解析:
在图像的正中间，(坐标(sr,sc)=(1,1)),
在路径上所有符合条件的像素点的颜色都被更改成2。
注意，右下角的像素没有更改为2，
因为它不是在上下左右四个方向上与初始点相连的像素点。
注意:

image 和 image[0] 的长度在范围 [1, 50] 内。
给出的初始点将满足 0 <= sr < image.length 和 0 <= sc < image[0].length。
image[i][j] 和 newColor 表示的颜色值在范围 [0, 65535]内。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/flood-fill
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 解法1 dfs
// 解法2 bfs
// 由 sr sc 对应的坐标点进行bfs即可 当遇到与 sr sc 相同的颜色替换即可
var dx = [4]int{-1, 1, 0, 0}
var dy = [4]int{0, 0, 1, -1}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	row := len(image)
	if row == 0 {
		return nil
	}
	col := len(image[0])
	queue := []int{sr, sc}
	isVisit := make([][]bool, row)
	for i, _ := range isVisit {
		isVisit[i] = make([]bool, col)
	}
	for len(queue) != 0 {
		sr, sc := queue[0], queue[1]
		queue = queue[2:]
		// 遍历当前坐标点上下左右四个节点 与当前节点颜色相同则加入queue
		for i := 0; i < 4; i++ {
			newSr := sr + dx[i]
			newSc := sc + dy[i]
			// 判断是否越界 没拜访过且是否与当前节点颜色相等
			if newSr < 0 || newSr >= row || newSc < 0 || newSc >= col ||
				isVisit[sr][sc] || image[newSr][newSc] != image[sr][sc] {
				continue
			}
			// 符合条件 入队
			queue = append(queue, newSr, newSc)
		}
		// 当前节点颜色置换
		image[sr][sc] = newColor
		isVisit[sr][sc] = true
	}
	return image
}

func TestFloodFill(t *testing.T) {
	image1 := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	image2 := [][]int{{0, 0, 0}, {0, 0, 0}}
	fmt.Println(floodFill(image1, 1, 1, 2))
	fmt.Println(floodFill(image2, 0, 0, 2))
}
