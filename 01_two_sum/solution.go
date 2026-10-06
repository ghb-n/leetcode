package main

import "fmt"

func solution(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if idx, ok := seen[complement]; ok {
			return []int{idx, i}
		}
		seen[num] = i
	}
	return nil
}

func main() {
	result := solution([]int{2, 7, 11, 15}, 9)
	fmt.Println(result)
}
