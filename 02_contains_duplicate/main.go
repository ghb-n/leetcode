package main

import "fmt"

func containsDuplicate(nums []int) bool {
	seen := make(map[int]int)

	for i, num := range nums {
		if _, ok := seen[num]; ok {
			return true
		}
		seen[num] = i
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{2, 7, 2, 11}))
}
