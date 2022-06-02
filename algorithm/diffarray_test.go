package algorithm

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
