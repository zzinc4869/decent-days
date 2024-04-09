package alg

import (
	"fmt"
	"math"
	"testing"
)

// leetcode 41
func TestFirstMissingPositive(t *testing.T) {
	missArr := []int{7, 8, 9, 11, 12}
	missV := firstMissingPositive(missArr)
	fmt.Println(missV)
}

func firstMissingPositive(nums []int) int {
	valueSet := make(map[int]bool)
	missV := math.MaxInt
	for i := 0; i < len(nums); i++ {
		valueSet[nums[i]] = true
	}

	for i := 0; i < len(nums); i++ {
		min1 := nums[i] - 1
		if _, ok := valueSet[min1]; !ok && min1 < missV && min1 > 0 {
			missV = min1
		}

		max1 := nums[i] + 1
		if _, ok := valueSet[max1]; !ok && max1 < missV && max1 > 0 {
			missV = max1
		}
	}

	if _, ok := valueSet[1]; !ok {
		missV = min(1, missV)
	}

	return missV
}
