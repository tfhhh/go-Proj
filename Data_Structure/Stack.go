package main

import "fmt"

const SIZE = 10
type Stack struct {
	top int
	data [SIZE]int
}
//初始化
func Init(S *Stack) *Stack {
	S.top = 0
	S.data = [SIZE]int{}
	return S
}
//判空
func Empty(S *Stack) bool {
	return S.top == 0
}
//栈长
func Len(S *Stack) int {
	return S.top
}
//进栈
func Push(S *Stack,data int) *Stack {
	if S.top == SIZE {
		fmt.Println("栈满，返回原栈！")
		return S
	}else {
		S.data[S.top] = data	//压入数据
		S.top++	//栈顶+1
		return S
	}
}
//出栈
func Pop(S *Stack) *Stack  {
	if S.top == 0 {
		return S
	}else {
		S.top--	//栈顶减1
		S.data[S.top] = -1	//数据置空
		return S
	}
}
//取栈顶
func GetTop(S *Stack) int {
	if S.top == 0 {
		return -1
	}else {
		return S.data[S.top-1]
	}

}
func main() {
	var mystack *Stack
	mystack = new(Stack)
	Init(mystack)
	fmt.Println("Empty mystack:",Empty(mystack))
	//进栈
	for i := 0; i < 11; i++ {
		mystack = Push(mystack,i+1)
	}
	fmt.Println("Push 1-10:",mystack)
	fmt.Println("mystack len:",Len(mystack))
	fmt.Println("Empty mystack:",Empty(mystack))
	for i := 0; i < 2; i++ {
		mystack = Pop(mystack)
	}
	fmt.Println("Pop 10,9:",mystack)
	fmt.Println("mystack len:",Len(mystack))
	fmt.Println("GetTop:",GetTop(mystack))
	/*
	Empty mystack: true
	栈满，返回原栈！
	Push 1-10: &{10 [1 2 3 4 5 6 7 8 9 10]}
	mystack len: 10
	Empty mystack: false
	Pop 10,9: &{8 [1 2 3 4 5 6 7 8 -1 -1]}
	mystack len: 8
	GetTop: 8

	Process finished with exit code 0
	*/
}
