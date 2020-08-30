package main

import "fmt"
/*
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
输入: [1,2,3,4,5,6,7] 和 k = 3
输出: [5,6,7,1,2,3,4]
 */
func rotate(num []int,k int) []int {
	k = k%len(num)
	a := num[:len(num)-k]
	b := num[(len(num)-k):len(num)]
	c := make([]int,len(num))
	c = append(b,a...)
	return c
}
func main() {
	x := []int{1,2,3,4,5,6,7}
	y := []int{-1}
	fmt.Println(rotate(y,2))
	fmt.Println(rotate(x,8))
}