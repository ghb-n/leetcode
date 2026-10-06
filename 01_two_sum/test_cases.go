package main

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{"example 1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"example 2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"example 3", []int{3, 3}, 6, []int{0, 1}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := solution(c.nums, c.target)
			if !reflect.DeepEqual(got, c.expected) {
				t.Errorf("got %v, want %v", got, c.expected)
			}
		})
	}
}
