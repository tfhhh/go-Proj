package main

import (
	"fmt"
)

const StackSize int = 10
//栈，操作受限的线性表
type stack struct {
	top int		//栈顶
	length int	//栈长
	Data[] int  //栈元素
}
//初始化
func(S *stack) Init() bool{
	if S.top < 0 {
		return false
	}else {
		S.top = 0                       //栈顶置0
		S.length = StackSize            //设置栈长
		S.Data = make([]int, StackSize) //创建栈元素数组
		return true
	}
}
//判空
func(S *stack) IsEmpty() bool {
	return S.top == 0
}
//求栈长，S.top-S.base
func(S *stack) Length() int {
	return S.top
}
//进栈
func(S *stack) Push(data int) bool {
	//栈满
	if S.top == S.length {
		return false
	}else {
		S.Data[S.top] = data	//压入数据
		S.top++					//栈顶+1
		return true
	}
}
//出栈
func(S *stack) Pop() bool {
	//栈空
	if S.top == 0 {
		return false
	}else {
		S.top--			  //栈顶-1
		S.Data[S.top] = 0 //出栈元素置空
		return true
	}
}
//取栈顶元素
func(S *stack) GetTop() int{
	if S.top == 0{
		return 0
	}else {
		return S.Data[S.top-1]
	}
}
func main() {
	//创建栈类型
	var myStack stack
	//初始化
	myStack.Init()
	fmt.Println("original stack:",myStack)
	for i := -1; i < 5; i++ {
		myStack.Push(i)
	}
	//初始栈
	fmt.Println("Push stack:",myStack)	//[-1,0，1，2，3，4,0,0,0,0]
	//判空
	fmt.Println("IsEmpty:",myStack.IsEmpty())
	//求栈长
	fmt.Println("Length：",myStack.Length())
	//进栈
	myStack.Push(-2)
	myStack.Push(6)//	[-1,0,1,2,3,4,-2,6,0,0]
	fmt.Printf("push stack:%v,length:%d\n",myStack,myStack.Length())
	//出栈
	myStack.Pop()	//[-1,0,1,2,3,4,-2,0,0,0]
	fmt.Printf("pop stack:%v,length:%d\n",myStack,myStack.Length())
	//取栈顶
	fmt.Println("stack top:",myStack.GetTop())	// top=-2
}