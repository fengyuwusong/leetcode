package main

import "fmt"

type difference struct {
	diffArray []int
}

func NewDifference(nums []int) *difference {
	diffArray := make([]int, len(nums))
	diffArray[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		diffArray[i] = nums[i] - diffArray[i-1]
	}
	return &difference{diffArray: diffArray}
}

func (d *difference) increment(i, j, val int) {
	d.diffArray[i] += val
	if j+1 < len(d.diffArray) {
		d.diffArray[j+1] -= val
	}
}

func (d *difference) result() []int {
	res := make([]int, len(d.diffArray))
	res[0] = d.diffArray[0]
	for i := 1; i < len(d.diffArray); i++ {
		res[i] = d.diffArray[i] + res[i-1]
	}
	return res
}

func corpFlightBookings(bookings [][]int, n int) []int {
	d := NewDifference(make([]int, n))
	for i := 0; i < len(bookings); i++ {
		d.increment(bookings[i][0]-1, bookings[i][1]-1, bookings[i][2])
	}
	return d.result()
}

func main() {
	fmt.Println(corpFlightBookings([][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}, 5))
}
