package alg

import (
	"fmt"
	"sort"
	"testing"
)

// leetcode40
func TestCombinationSum2(t *testing.T) {
	//candidates := []int{10, 1, 2, 7, 6, 1, 5}
	candidates := []int{3, 1, 3, 5, 1, 1}
	target := 8
	res := combinationSum2(candidates, target)
	fmt.Println(res)
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	path := make([]int, 0)
	res := make([][]int, 0)
	res = dfsSum(candidates, len(candidates), 0, target, path, res)
	return res
}

func dfsSum(candidates []int, length, begin, target int, path []int, res [][]int) [][]int {
	if target == 0 {
		res = append(res, append([]int(nil), path...))
		return res
	}
	for i := begin; i < length; i++ {
		curr := candidates[i]
		// 大剪枝：减去 candidates[i] 小于 0，减去后面的 candidates[i + 1]、candidates[i + 2] 肯定也小于 0，因此用 break
		if target-curr < 0 {
			break
		}

		// 小剪枝：同一层相同数值的结点，从第 2 个开始，候选数更少，结果一定发生重复，因此跳过，用 continue
		if i > begin && curr == candidates[i-1] {
			continue
		}

		path = append(path, curr)
		fmt.Printf("递归之前 => %v，剩余 = %v\n", path, target-candidates[i])

		res = dfsSum(candidates, length, i+1, target-curr, path, res)

		path = path[0 : len(path)-1]
		fmt.Printf("递归之后 => %v，剩余 = %v\n", path, target-candidates[i])
	}
	return res
}
