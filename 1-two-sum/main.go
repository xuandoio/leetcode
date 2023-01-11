package main

import "fmt"

func main() {
	i := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(i)
}

func twoSum(nums []int, target int) []int {
	var tmp = make(map[int]int)
	for i, v := range nums {
		if _, ok := tmp[target-v]; ok {
			return []int{tmp[target-v], i}
		}
		tmp[v] = i
	}

	return []int{}
}
