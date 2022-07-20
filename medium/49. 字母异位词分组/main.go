package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	dict := make(map[string][]string)
	for _, str := range strs {
		key := sortString(str)
		if _, ok := dict[key]; !ok {
			dict[key] = make([]string, 0)
		}
		dict[key] = append(dict[key], str)
	}
	var res [][]string
	for _, v := range dict {
		res = append(res, v)
	}
	return res
}

func sortString(str string) string {
	bytes := []byte(str)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	return string(bytes)
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{""}))
	fmt.Println(groupAnagrams([]string{"", ""}))
	fmt.Println(groupAnagrams([]string{"a"}))
}
