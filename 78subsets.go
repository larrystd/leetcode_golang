package main

import (
	"fmt"
)

func subsets(nums []int) [][]int {
	subsets_ := [][]int{[]int{}}
	for i := range nums {
		getsets(nums, &subsets_, i)
	}
	return subsets_
}

func getsets(nums []int, subsets *[][]int, idx int) {
	n := len(*subsets)
	for i := 0; i < n; i++ {
		news := make([]int, len((*subsets)[i]))
		copy(news, (*subsets)[i]) // deep copy, dest : src
		news = append(news, nums[idx])
		*subsets = append(*subsets, news)
	}
}

func main2() {
	nums := []int{1, 2, 3}
	subsets := subsets(nums)
	for _, subset := range subsets {
		for _, e := range subset {
			fmt.Printf("%d ", e)
		}
		fmt.Print("\n")
	}
}
