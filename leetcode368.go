package main

import (
	"fmt"
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	sort.Ints(nums)
	n := len(nums)
	dp := make([]int, n)
	prev := make([]int, n)
	maxIdx := 0

	for i := range nums {
		dp[i] = 1
		prev[i] = -1
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
				prev[i] = j
			}
		}
		if dp[i] > dp[maxIdx] {
			maxIdx = i
		}
	}

	fmt.Println(dp)
	fmt.Println(prev)

	result := []int{}
	for i := maxIdx; i >= 0; i = prev[i] {
		result = append([]int{nums[i]}, result...)
	}

	return result
}
