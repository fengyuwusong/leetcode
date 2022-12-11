package main

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	index := 0
	switch ruleKey {
	case "type":
		index = 0
	case "color":
		index = 1
	case "name":
		index = 2
	}
	var ans int
	for i := 0; i < len(items); i++ {
		if items[i][index] == ruleValue {
			ans++
		}
	}
	return ans
}

func main() {
	println(countMatches([][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "phone"}, {"phone", "gold", "iphone"}}, "type", "phone"))
}
