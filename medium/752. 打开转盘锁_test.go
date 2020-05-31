package medium

import (
	"fmt"
	"testing"
)

/**
你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' 。
每个拨轮可以自由旋转：例如把 '9' 变为  '0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。

锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。

列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。

字符串 target 代表可以解锁的数字，你需要给出最小的旋转次数，如果无论如何不能解锁，返回 -1。



示例 1:

输入：deadends = ["0201","0101","0102","1212","2002"], target = "0202"
输出：6
解释：
可能的移动序列为 "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202"。
注意 "0000" -> "0001" -> "0002" -> "0102" -> "0202" 这样的序列是不能解锁的，
因为当拨动到 "0102" 时这个锁就会被锁定。
示例 2:

输入: deadends = ["8888"], target = "0009"
输出：1
解释：
把最后一位反向旋转一次即可 "0000" -> "0009"。
示例 3:

输入: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"], target = "8888"
输出：-1
解释：
无法旋转到目标数字且不被锁定。
示例 4:

输入: deadends = ["0000"], target = "8888"
输出：-1


提示：

死亡列表 deadends 的长度范围为 [1, 500]。
目标数字 target 不会在 deadends 之中。
每个 deadends 和 target 中的字符串的数字会在 10,000 个可能的情况 '0000' 到 '9999' 中产生。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/open-the-lock
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func openLock(deadends []string, target string) int {
	// 制作deadends hash
	deadendsHash := make(map[string]struct{})
	for _, value := range deadends {
		deadendsHash[value] = struct{}{}
	}
	// 判断起始节点是否存在deadends中
	if _, exists := deadendsHash["0000"]; exists {
		return -1
	}
	// 判断目标节点是否在死亡数组中 存在则直接返回-1
	if _, exists := deadendsHash[target]; exists {
		return -1
	}

	// 是否遍历过 手动实现set结构 记录到达当前节点需要的步数 10^4次方 使用uint即可
	isVisited := make(map[string]uint)
	// 用于记录下一步可到达的节点
	// 根节点加入队列
	queue := []string{"0000"}
	// 根节点步数为0
	isVisited["0000"] = 0
	// 开始遍历队列
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		// 依次遍历每一位 记录下一步可到达的节点
		var nexts []string
		for i := 0; i < len(node); i++ {
			// 分别将其 +1 -1 若没遍历过加入队列 否则跳过
			upNode := up(node, i)
			downNode := down(node, i)
			nexts = append(nexts, upNode, downNode)
		}
		// 遍历下一步可到达的节点 去除 已访问 或 在死亡数组中的节点
		for _, next := range nexts {
			_, ok1 := isVisited[next]
			_, ok2 := deadendsHash[next]
			if !ok1 && !ok2 {
				// 记录到达当前节点需要的步数
				isVisited[next] = isVisited[node] + 1
				// 是否目标节点
				if target == next {
					return int(isVisited[next])
				}
				// 加入队列进行下一轮的bfs
				queue = append(queue, next)
			}
		}
	}
	return -1
}

// 实现对应位数+1方法
func up(node string, index int) string {
	// 将node转为rune
	nodeRune := []rune(node)
	// 重新计算对应位
	nodeRune[index] = (nodeRune[index]-'0'+1)%10 + '0'
	return string(nodeRune)
}

// 实现对应位数-1方法
func down(node string, index int) string {
	// 将node转为rune
	nodeRune := []rune(node)
	// 重新计算对应位 若-1 可能出现负数结果错误 (n+9) % 10 同样起到-1效果
	nodeRune[index] = (nodeRune[index]-'0'+9)%10 + '0'
	return string(nodeRune)
}

func TestOpenLock(t *testing.T) {
	fmt.Println(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"))
	fmt.Println(openLock([]string{"0000"}, "8888"))
}
