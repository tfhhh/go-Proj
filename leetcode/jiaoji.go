package main

import "fmt"

//给定两个数组，编写一个函数来计算它们的交集。
/*
eg1
输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出: [4,9],可以不考虑输出结果的顺序
eg2
输入: nums1 = [1,2,2,1], nums2 = [2,2]
输出: [2,2]
 */
//map映射，先用map统计数字以及次数，再遍历num2，统计至少减了一次的数字
func InterSet(num1[] int,num2[] int) []int {
	m := map[int]int{}
	//统计num1的数字和次数
	for _, v := range num1 {
		m[v] += 1
	}
	fmt.Println(m)
	//遍历num2，统计相同的数字
	n := make([]int,len(num1))	//n切片，num1或num2长度，交集长度总会小于或等于num1或num2长度
	k := 0
	for _, u := range num2 {
		if m[u] > 0 {	//出现过的数字去匹配map
			m[u] -= 1	//次数减1
			n[k] = u	//将相同的数字存入n切片
			k++
		}
	}
	fmt.Println(n)
	return n[0:k]	//返回k个相同的数字的数组
}
func main() {
	a := []int{4,9,5,9}
	b := []int{9,4,9,8,4}
	fmt.Println(InterSet(a,b))
}
