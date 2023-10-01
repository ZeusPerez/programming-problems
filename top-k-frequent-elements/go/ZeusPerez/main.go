package main

import (
	"fmt"
)

func main() {
	inputs := []struct {
		nums []int
		k    int
	}{
		{nums: []int{1, 1, 1, 2, 2, 3}, k: 2},
		{nums: []int{1}, k: 1},
	}

	for _, input := range inputs {
		fmt.Printf("%+v\n", topKFrequent(input.nums, input.k))
	}

	return
}

func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int)
	for _, num := range nums {
		if count, ok := countMap[num]; ok {
			countMap[num] = count + 1
		} else {
			countMap[num] = 1
		}
	}

	countSlice := make([][]int, len(nums)+1)

	for num, count := range countMap {
		countSlice[count] = append(countSlice[count], num)
	}

	result := make([]int, 0)
	for i := len(countSlice) - 1; i >= 0; i-- {
		if len(result) >= k {
			return result
		}
		result = append(result, countSlice[i]...)
	}

	return result
}
