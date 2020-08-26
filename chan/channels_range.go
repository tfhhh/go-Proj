package main

import "fmt"

func main()  {
	msg := make(chan int)
	go func() {
		for i := 0; i < 4; i++ {
			msg <- i
		}
		close(msg)
	}()
	fmt.Println(<-msg)	//0
	fmt.Println(<-msg)	//1
	for v := range msg {
		fmt.Println("data:",v)	//通道里还剩2，3两个数据未被接收
	}
	_,ok := <- msg
	if !ok{
		fmt.Println("closed channels")
	}
}
