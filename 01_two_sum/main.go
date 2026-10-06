package main

import "fmt"

func twoSum(nums []int, target int) []int {
    // seen is a map. Key: a number we've passed. Value: its index in nums.
    // We use it to look up "have I seen this number before?" in O(1).
    seen := make(map[int]int)

    for i, num := range nums {
        // complement is the number we need to find to reach target.
        // If num + complement = target, then complement = target - num.
        complement := target - num

        // If we've already seen the complement, we have our pair.
        // idx is where we saw it. i is where we are now.
        if idx, ok := seen[complement]; ok {
            return []int{idx, i}
        }

        // We haven't seen the complement. Store this number and its index
        // so future iterations can find it.
        seen[num] = i
    }

    return nil
}

func main() {
	fmt.Println(twoSum([]int{2,7,11,15}, 9))
	fmt.Println(twoSum([]int{3,2,4}, 6))
	fmt.Println(twoSum([]int{3,3}, 6))
}
