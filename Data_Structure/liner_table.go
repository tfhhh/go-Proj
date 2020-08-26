package main

import "fmt"

const LtSize = 10
type List struct {
	length int
	Data[LtSize] int
}
//初始化
func(L *List) Init(data int,index int){
	L.length++
	L.Data[index] = data
}
//插入，把data插在index前面
func(L *List) Insert(data int,index int)bool  {
	if index < 0 || index > LtSize || L.length == LtSize {
		return false
	}else {
		for j := L.length+1; j >= index; j-- {
			L.Data[j] = L.Data[j-1]	//元素后移
		}
		L.Data[index-1] = data	//插入元素
		L.length++	//表长加1
		return true
	}
}
//删除指定位置元素
func(L *List) Delete(index int)bool {
	if index < 1 || index > L.length {
		return false
	}else {
		for j := index-1; j <= L.length+1; j++ {
			L.Data[j] = L.Data[j+1]	//元素前移
		}
		L.length--
		return true
	}
}
//查找指定元素
func (L *List) Search(data int) bool {
	for i := 0; i < L.length; i++ {
		if L.Data[i] == data {
			return true
		}
	}
	return false
}
func main() {
	var myList List
	for i := 0; i < 6; i++ {
		myList.Init(i,i)
	}
	fmt.Println("Init List:",myList)

	myList.Insert(3,1)
	fmt.Println("insert (data:3,index:1):",myList)

	myList.Delete(3)
	fmt.Println("delete (index:3):",myList)
}