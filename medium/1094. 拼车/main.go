package main

// 差分数组
type Difference struct {
	diffArray []int // 差分数组
}

func NewDifference(nums []int) *Difference {
	// 构造差分数组
	diffArray := make([]int, len(nums))
	diffArray[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		diffArray[i] = nums[i] - diffArray[i-1]
	}
	return &Difference{diffArray: diffArray}
}

func (d *Difference) increment(i, j, val int) {
	d.diffArray[i] += val
	if j+1 < len(d.diffArray) {
		d.diffArray[j+1] -= val
	}
}

func (d *Difference) result() []int {
	// 根据差分数组反推原数组并返回
	res := make([]int, len(d.diffArray))
	res[0] = d.diffArray[0]
	for i := 1; i < len(res); i++ {
		res[i] = res[i-1] + d.diffArray[i]
	}
	return res
}

func carPooling(trips [][]int, capacity int) bool {
	// 创建一个数组 长度1000用于表示每个公里数上的乘客数量
	passengerNums := make([]int, 10000)
	d := NewDifference(passengerNums)
	// 遍历所有旅程 依次加上每个公里的乘客
	for _, trip := range trips {
		d.increment(trip[1], trip[2]-1, trip[0])
	}

	// 遍历所有公里数 查询是否有大于容量的公里
	res := d.result()
	for i := 0; i < len(res); i++ {
		if res[i] > capacity {
			return false
		}
	}

	return true
}

func main() {
	println(carPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 4))
	println(carPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 5))
	println(carPooling([][]int{{2, 1, 5}, {3, 5, 7}}, 3))
}
