package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
//插入
func Insert(data int) *TreeNode {
	return &TreeNode{data,nil,nil}
}

//先序遍历
func PreOrder(root *TreeNode) []int {
	var res1 []int
	node := root
	if node == nil{
		return nil
	}else {
		res1 = append(res1,node.Val)	//根的遍历
		res1 = append(res1,PreOrder(node.Left)...)	//左子树所有节点
		res1 = append(res1,PreOrder(node.Right)...)	//右子树所有节点
		return  res1
	}
}
//中序遍历
func InOrder(root *TreeNode) []int {
	var res2 []int
	node := root
	if root == nil{
		return nil
	}else {
		res2 = append(res2,InOrder(node.Left)...)
		res2 = append(res2,node.Val)
		res2 = append(res2,InOrder(node.Right)...)
		return  res2
	}
}
//后序遍历
func PostOrder(root *TreeNode) []int {
	var res3 []int
	node := root
	if node == nil{
		return nil
	}else {
		res3 = append(res3,PostOrder(node.Left)...)
		res3 = append(res3,PostOrder(node.Right)...)
		res3 = append(res3,node.Val)
		return  res3
	}
}
//节点数
func Count(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + Count(root.Left) + Count(root.Right)	//根节点+左子树节点+右子树节点
}
//判定二叉排序树,中序遍历为递增序列
func SortTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var myres []int
	var x = true
	myres = InOrder(root)
	for i := 0; i < Count(root)-1; i++ {	//判断中序遍历是否递增
		if  myres[i] > myres[i+1] {
			x = false
			break
		}
	}
	return x
}
//二叉排序树的查找
func Search(root *TreeNode,data int) bool {
	if root == nil {
		return false
	}
	if root.Val == data {	//搜索数据成功
		return true
	}else if root.Val > data {
		return Search(root.Left,data)	//数据小于根节点，往左子树搜索
	}else {
		return Search(root.Right,data)	//数据大于根节点，往右子树搜索
	}
	return false
}
//高度
func Height(root *TreeNode) int {
	if root == nil {
		return 0
	}else {
		LH := Height(root.Left)
		RH := Height(root.Right)
		if LH >= RH {
			return 1 + LH	//左子树高则为LH+1（根节点）
		}else {
			return 1 + RH
		}
	}
}
//平衡二叉树的判定
func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}else {
		//平衡节点的左右子树为平衡二叉树
		LH := Height(root.Left)
		RH := Height(root.Right)
		if LH-RH > 1 || LH-RH < -1{
			return false
		}
		return IsBalanced(root.Left) && IsBalanced(root.Right)
	}
}
func main() {
	/*
	1
  2	  3
4	5
	 */
	var my_tree *TreeNode
	fmt.Println(PreOrder(my_tree))
	my_tree = Insert(1)
	my_tree.Left = Insert(2)
	my_tree.Right= Insert(3)
	my_tree.Left.Left = Insert(4)
	my_tree.Left.Right= Insert(5)
	fmt.Println("------------二叉树⬇------------")
	fmt.Println("PreOrder my_tree:",PreOrder(my_tree))
	fmt.Println("InOrder my_tree:",InOrder(my_tree))
	fmt.Println("PostOrder my_tree:",PostOrder(my_tree))
	fmt.Println("my_tree node count：",Count(my_tree))
	fmt.Println("my_tree Height:",Height(my_tree))
	fmt.Println("my_tree IsBalanced:",IsBalanced(my_tree))
	fmt.Println("Sort my_tree:",SortTree(my_tree))
	fmt.Println("------------排序二叉树⬇------------")
	/*
	   45
   12	  53
 3    37
   24 	 42
	 */
	var MyOrder *TreeNode
	MyOrder = Insert(45)
	MyOrder.Left = Insert(12)
	MyOrder.Right= Insert(53)
	MyOrder.Left.Left = Insert(3)
	MyOrder.Left.Right = Insert(37)
	MyOrder.Left.Right.Left = Insert(24)
	MyOrder.Left.Right.Right = Insert(42)
	fmt.Println("InOrder MyOrder:",InOrder(MyOrder))
	fmt.Println("sort MyOrder:",SortTree(MyOrder))
	fmt.Println("MyOrder Height:",Height(MyOrder))
	fmt.Println("MyOrder node count：",Count(MyOrder))
	fmt.Println("MyOrder IsBalanced:",IsBalanced(MyOrder))
	fmt.Println("Search 42:",Search(MyOrder,42))
	fmt.Println("Search 88:",Search(MyOrder,88))
	MyOrder.Left.Right.Left = Insert(40)
	fmt.Println("24->40,InOrder MyOrder:",InOrder(MyOrder))
	fmt.Println("24->40,sort MyOrder:",SortTree(MyOrder))
	fmt.Println("------------------------")
	/*
	------------二叉树⬇------------
	PreOrder my_tree: [1 2 4 5 3]
	InOrder my_tree: [4 2 5 1 3]
	PostOrder my_tree: [4 5 2 3 1]
	my_tree node count： 5
	my_tree Height: 3
	my_tree IsBalanced: true
	Sort my_tree: false
	------------排序二叉树⬇------------
	InOrder MyOrder: [3 12 24 37 42 45 53]
	sort MyOrder: true
	MyOrder Height: 4
	MyOrder node count： 7
	MyOrder IsBalanced: false
	Search 42: true
	Search 88: false
	24->40,InOrder MyOrder: [3 12 40 37 42 45 53]
	24->40,sort MyOrder: false
	------------------------

	Process finished with exit code 0

	 */
}
