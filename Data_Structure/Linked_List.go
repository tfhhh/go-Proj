package main

import (
	"fmt"
)

type ListNode struct {
  	Val int
    Next *ListNode
}
//判空
func IsEmpty(head *ListNode) bool {
	if head == nil {
		return true
	}
	return false
}
//链表长度
func Length(head *ListNode) int {
	if head == nil {
		return 0
	}
	count := 0
	for head.Next != nil {
		head = head.Next
		count++
	}
	return count + 1
}
//头部添加
func Add(head *ListNode,data int) *ListNode {
	node := &ListNode{}
	node.Val = data
	node.Next = head
	return node
}
//尾部添加
func Append(head *ListNode,data int) *ListNode  {
	if head == nil {
		head = Add(head,data)
		return head
	}
	node := &ListNode{}
	node.Val = data
	node.Next =  nil

	cur := head
	for cur.Next != nil  {
		cur = cur.Next
	}
	cur.Next = node
	return head
}
//指定位置插入
func Insert(head *ListNode,index int,data int) *ListNode {
	if index <= 0 {
		return head
	}else if index == 1{
		head = Add(head,data)
		return head
	}else if index == Length(head) {
		head = Append(head,data)
		return head
	}else {
		cur := head
		count := 1
		for count < index - 1 {
			cur = cur.Next
			count++
		}
		node := &ListNode{}
		node.Val = data
		node.Next = cur.Next
		cur.Next = node
		return head
	}
}
//查找
func Search(head *ListNode,data int) (bool,int) {
	if head == nil {
		return false,-1
	}
	cur := head
	count := 1
	result := false
	for cur.Next != nil {
		if cur.Val != data {
			cur = cur.Next
			count++
		}else {
			result =  true
			break
		}
	}
	if result == false {
		return false,-1
	}else {
		return true,count
	}

}
//删除指定位置元素
func Delete(head *ListNode,index int) *ListNode {
	if index == 1 {
		head = head.Next
		return head
	}
	cur := head
	count := 1
	for cur.Next != nil {
		if count < index - 1 {
			cur = cur.Next
		}
		count++
		if count == index-1 {
			break
		}
	}
	cur.Next = cur.Next.Next
	return head
}
//遍历
func ShowList(head *ListNode)  {
	if head == nil {
		return
	}
	for  {
		fmt.Printf("%d ",head.Val)
		if head.Next != nil {
			head = head.Next
		}else {
			break
		}
	}
	fmt.Println()
}
func main() {
	var MyList *ListNode
	for i := 5; i > 0; i-- {
		MyList = Add(MyList,i)
	}
	for i := 6; i < 11; i++ {
		MyList = Append(MyList,i)
	}
	//show
	MyList = Insert(MyList,5,11)
	fmt.Println("Length:",Length(MyList))
	fmt.Println("IsEmpty:",IsEmpty(MyList))
	ShowList(MyList)
	//search
	s1,ok1 := Search(MyList,11)
	s2,ok2 := Search(MyList,12)
	fmt.Printf("search 11:%v,index:%v\n",s1,ok1)
	fmt.Printf("search 12:%v,index:%v\n",s2,ok2)
	//delete
	MyList = Delete(MyList,1)
	ShowList(MyList)
	MyList = Delete(MyList,4)
	ShowList(MyList)
	/*
	Length: 11
	IsEmpty: false
	1 2 3 4 11 5 6 7 8 9 10
	search 11:true,index:5
	search 12:false,index:-1
	2 3 4 11 5 6 7 8 9 10
	2 3 4 5 6 7 8 9 10
	*/
}
