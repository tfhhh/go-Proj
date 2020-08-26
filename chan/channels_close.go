package main

import "fmt"

//如果Goroutine正在等待从通道接收数据，那么另一些Goroutine将会在该通道上写入数据，否则程序将会死锁。
func main() {
	msg := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			msg <- i
		}
		close(msg)
	}()
	fmt.Println(<-msg)	//接收0
	fmt.Println(<-msg)	//接收1
	for {
		v,ok := <- msg
		if !ok{
			fmt.Println("closed channels")	//输出data：2，true后输出关闭通道
			break
		}
		fmt.Println("data:",v,",",ok)		//剩余2未被接收，将在这里输出，data：2，true
	}
}