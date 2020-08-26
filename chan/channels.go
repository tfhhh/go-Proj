package main

import "fmt"
//通道(channels) 是连接多个协程的管道。 你可以从一个协程将值发送到通道，然后在另一个协程中接收
func main() {
	msg := make(chan string)	//创建通道变量
	go func() {
		msg <- "hello"	//发送"hello"到通道中
	}()
	myMsg := <- msg	//接收msg通道的信息
	fmt.Println(myMsg)
}
