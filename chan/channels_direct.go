package main

import "fmt"

//只写通道，发送数据，chan-<string定义为发送数据类型，msg为发送信息
func send(send chan<-string,msg string)  {
	send <- msg	//发射数据
}
//只读通道，接收数据，<-chan string定义接收数据通道类型
func receive(receive <-chan string,send chan<-string)   {
	msg := <- receive	//接收数据
	send <- msg			//发射数据

}
/*
send函数为只能发送数据的通道，receive可以接收数据与发送数据的通道，流程为send发送，receive接收并发送，
myMsg接收receive发送的数据
 */
func main() {
	msg1 := make(chan string,1)
	msg2 := make(chan string,1)

	send(msg1,"ni hao")
	fmt.Println(<-msg1)		//直接接收send函数的数据

	send(msg1,"hello")	//send函数发射数据
	receive(msg1,msg2)		//receive函数接收数据并发射
	myMsg2 := <-msg2		//接收receive函数的数据
	fmt.Println(myMsg2)
}