package main

import (
	"fmt"
	"sort"
)

func longestCommonPrefix(strs []string) string {
	sort.Strings(strs)
	first := strs[0]
	last := strs[len(strs)-1]

	i := 0
	for i < len(first) && first[i] == last[i] {
		i++
	}
	return first[:i]
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}
