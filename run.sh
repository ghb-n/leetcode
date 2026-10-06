#!/bin/bash

# Usage: ./run.sh two_sum
PROBLEM=$1

if [ -z "$PROBLEM" ]; then
    echo "Usage: ./run.sh <problem_name>"
    exit 1
fi

cd "$(dirname "$0")"

# Create folder
mkdir -p "$PROBLEM"
cd "$PROBLEM"

# Create solution.go if it doesn't exist
if [ ! -f solution.go ]; then
    cat > solution.go <<'EOF'
package main

import "fmt"

func solution(nums []int, target int) []int {
    // Write your solution here
    return nil
}

func main() {
    result := solution([]int{2, 7, 11, 15}, 9)
    fmt.Println(result)
}
EOF
    echo "Created solution.go — write your solution."
fi

# Create problem.md if it doesn't exist
if [ ! -f problem.md ]; then
    cat > problem.md <<'EOF'
# Problem

Describe the problem here.

## Example
Input: ...
Output: ...
EOF
    echo "Created problem.md — write the problem description."
fi

# Create test_cases.go if it doesn't exist
if [ ! -f test_cases.go ]; then
    cat > test_cases.go <<'EOF'
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
EOF
    echo "Created test_cases.go — write tests here."
fi

echo ""
echo "Problem: $PROBLEM"
echo ""
echo "Commands:"
echo "  go run solution.go     # run your solution"
echo "  go test                # run tests"
echo "  code problem.md        # edit the problem"