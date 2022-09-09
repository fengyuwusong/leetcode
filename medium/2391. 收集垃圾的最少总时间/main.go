package main

import (
	"fmt"
	"strings"
)

func garbageCollection(garbage []string, travel []int) int {
	var ans int
	var mTemp, pTemp, gTemp int
	for i := range garbage {
		mNum := strings.Count(garbage[i], "M")
		pNum := strings.Count(garbage[i], "P")
		gNum := strings.Count(garbage[i], "G")
		if mNum > 0 {
			ans += mNum + mTemp
			if i > 0 {
				ans += travel[i-1]
			}
			mTemp = 0
		} else if i > 0 {
			mTemp += travel[i-1]
		}
		if pNum > 0 {
			ans += pNum + pTemp
			if i > 0 {
				ans += travel[i-1]
			}
			pTemp = 0
		} else if i > 0 {
			pTemp += travel[i-1]
		}
		if gNum > 0 {
			ans += gNum + gTemp
			if i > 0 {
				ans += travel[i-1]
			}
			gTemp = 0
		} else if i > 0 {
			gTemp += travel[i-1]
		}
	}
	return ans
}

func main() {
	fmt.Println(garbageCollection([]string{"G", "P", "GP", "GG"}, []int{2, 4, 3}))
	fmt.Println(garbageCollection([]string{"MMM", "PGM", "GP"}, []int{3, 10}))
}
