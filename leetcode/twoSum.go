package main

import "fmt"
/*
给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
 */
func twoSum(nums []int, target int) []int {
	res := make([]int,2)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i]+nums[j] == target && i < j {
				res[0] = i
				res[1] = j
			}
		}
	}
	return res
}
func main() {
	a := []int{3,3}
	fmt.Println(twoSum(a,6))
}