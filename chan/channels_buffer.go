package main

import (
	"fmt"
)

func main() {
	msg := make(chan string,2)	//最多缓存两个值
	//由于此通道是有缓冲的， 因此我们可以将这些值发送到通道中，而无需并发的接收。
	msg <- "msg1"
	msg <- "msg2"
	fmt.Println(<-msg)
	fmt.Println(<-msg)
}
