package main

import (
	"fmt"
)

type Data interface {}

//链表元素结构体
type Node struct {
	Elem Data	//数据域
	Next *Node	//指针域
}
//链表结构体，都从头节点开始操作
type LinkedList struct {
	Head *Node	//头节点
}
//判空
func(L *LinkedList) IsEmpty() bool{
	if L.Head == nil {
		return true
	}
	return false
}
//表长
func (L *LinkedList) Length() int {
	start := L.Head
	count := 0
	for start != nil {
		start = start.Next	//遍历下一个节点
		count++	//计算遍历的次数
	}
	return count
}
//头部添加元素
func (L *LinkedList) Add(e Data) *Node {
	node := &Node{Elem: e}	//定义头部添加数据节点
	node.Next = L.Head	//数据节点的Next指针为原节点的头部指针
	L.Head = node	//数据节点转换为头部节点
	return node
}
//尾部添加元素
func (L *LinkedList) Append(e Data) {
	node := &Node{Elem: e}	//定义尾部添加数据指针
	if L.IsEmpty() == true {
		L.Head = node
	}else {
		cur := L.Head
		for cur.Next != nil {
			cur = cur.Next	//遍历指针域
		}
		cur.Next = node	//cur.Next = nil时将cur.Next连接为尾部添加的数据指针
	}
}
//插入元素,i为位置，e为元素
func (L LinkedList) Insert(i int,e Data)  {
	if i < 0 {
		L.Add(e)	//头部添加
	} else if i > L.Length() {
		L.Append(e)	//尾部添加
	}else {
		pre := L.Head
		for j:=0; j < i-1; j++ {
			pre = pre.Next	//遍历指针域到前一个节点
		}
		node := &Node{Elem: e}	//创建插入数据的指针域
		node.Next = pre.Next	//指针域连接
		pre.Next = node			//指针域连接
	}
}
//查找e值
func (L *LinkedList) Search(e Data) bool {
	start := L.Head
	for start != nil {
		if start.Elem == e {
			return true
		}
		start = start.Next	//遍历节点，直到节点指针域为空
	}
	return false
}

//删除e值
func (L *LinkedList) Delete(e int)  {
	pre := L.Head
	if pre.Elem == e {
		L.Head = pre.Next	//删除值为头节点的值，链表的头节点为pre.Next
	}else {
		for pre.Next != nil {
			if pre.Next.Elem == e {
				pre.Next = pre.Next.Next	//删除所有节点值为e的节点
			}else {
				pre = pre.Next	//查找下一节点
			}
		}
	}
}
//删除链表的第 n 个节点
func(L *LinkedList) removeNthFromEnd(n int) bool{
	var count = 1
	pre := L.Head
	if n <= 0 {
		return false
	}else if n == 1 {
		L.Head = pre.Next
		return true
	}else {
		for pre.Next != nil {
			if count == n-1 {
				pre.Next = pre.Next.Next
				break
			}else {
				pre = pre.Next
				count++
			}
		}
		return true
	}
}
//遍历
func (L *LinkedList) ShowList(){
	if L.IsEmpty() != true {
		start := L.Head
		for {
			fmt.Printf("%v ",start.Elem)
			if start.Next != nil {
				start = start.Next //遍历下一节点
			}else {
				break
			}
		}
		fmt.Println()
	}
}
func main() {
	var MyLinkedList LinkedList
	MyLinkedList.Add(0)
	MyLinkedList.Append(1)
	MyLinkedList.Add(3)
	MyLinkedList.Insert(1,4)
	MyLinkedList.Append(1)
	MyLinkedList.ShowList()	//3 4 0 1 1
	fmt.Println("IsEmpty:",MyLinkedList.IsEmpty())
	fmt.Println("original length:",MyLinkedList.Length())

	mySearch1 := MyLinkedList.Search(1)
	fmt.Println("search 1:", mySearch1)
	mySearch2 := MyLinkedList.Search(2)
	fmt.Println("search 2:",mySearch2)

	MyLinkedList.Delete(1)
	MyLinkedList.ShowList()	//3 4 0
	fmt.Println("deleted length:",MyLinkedList.Length())

}

