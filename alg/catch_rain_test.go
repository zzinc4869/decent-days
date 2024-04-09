package alg

import (
	"fmt"
	"testing"
)

// leetcode 42
func TestCatchRainWater(t *testing.T) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	res := trap(height)
	fmt.Println(res)
}

func trap(height []int) int {
	lens := len(height)
	leftMax := make([]int, lens)
	rightMax := make([]int, lens)

	left, right := 0, 0
	for i := 0; i < lens; i++ {
		if height[i] > left {
			left = height[i]
		}
		leftMax[i] = left
	}

	for i := lens - 1; i >= 0; i-- {
		if height[i] > right {
			right = height[i]
		}
		rightMax[i] = right
	}

	fmt.Println(leftMax)
	fmt.Println(rightMax)

	total := 0
	for i := 0; i < lens; i++ {
		total = total + min(leftMax[i], rightMax[i]) - height[i]
	}
	return total
}
