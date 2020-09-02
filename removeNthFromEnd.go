package main

import (
	"fmt"
)

type Node struct {
	Val int
	Next *Node
}

func add(head *Node,data int) *Node {
	node := &Node{}
	node.Val = data
	node.Next = head
	return node
}
func Show(head *Node) []int {
	if head == nil {
		return nil
	}
	var res []int
	for head.Next != nil {
		res = append(res,head.Val)
		head = head.Next
	}
	res = append(res,head.Val)
	return res
}
//删除倒数第k个节点
func removeNthFromEnd(head *Node,n int) *Node {
	node := &Node{}
	node.Next = head
	//head：[1,2,3,4,5]
	//node：[0,1,2,3,4,5]
	/*增加一个哨兵节点，当head走到正数n时，node开始遍历，head走到nil，node走到删除节点倒数n的前一位置
	, 比如删除倒数第4个节点（2），当head走到4时，node开始遍历，head走1步后为空，
	node走一步后所在位置为删除节点前一位置
	 */
	var cur *Node
	cur = node
	i := 1

	for head.Next != nil { 	//head开始遍历
		head = head.Next
		if i >= n {			//当走到n时node开始遍历
			node = node.Next
		}
		i++
	}
	//head.Next == nil,删除node所在位置的下一个节点
	node.Next = node.Next.Next
	return cur.Next	//去掉头节点
}
func main() {
	var node *Node
	for i := 7; i > 1; i-- {
		node = add(node,i-1)
	}
	fmt.Println(Show(node))				//[1,2,3,4,5,6]
	node = removeNthFromEnd(node,5) 	//[1,3,4,5,6]
	node = removeNthFromEnd(node,2)	//[1,3,4,6]
	fmt.Println(Show(node))				//[1,3,4,6]

}