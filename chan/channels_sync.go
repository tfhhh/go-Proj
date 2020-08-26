package main

import (
	"fmt"
	"time"
)

func worker(done chan bool)  {
	fmt.Println("working")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true	//发射一个值通知完成
}
func main() {
	done := make(chan bool)	//通道，布尔变量
	go worker(done)	//创建协程
	<-done	//程序一直阻塞，直到收到worker的发送完成通知
	//x := <- done	//程序一直阻塞，直到收到worker的发送完成通知
	//fmt.Println(x)	//true
}
