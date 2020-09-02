package main
import "fmt"

const QueueMax int = 10
//循环队列
type Queue struct {
	rear int
	front int
	Data[] int
}
//初始化
func InitQueue(Q *Queue) *Queue {
	Q.rear = 0
	Q.front = 0
	Q.Data = make([]int,QueueMax)
	return &Queue{Q.rear,Q.front,Q.Data}
}
//队长
func Leng(Q *Queue) int {
	return (Q.rear - Q.front + QueueMax)%QueueMax
}
//入队
func EnQueue(Q *Queue,data int) *Queue {
	//队满
	if (Q.rear+1)%QueueMax == Q.front {
		fmt.Println("data:",data)
		fmt.Println("队列已满，返回原队列！")
		return Q
	} else {
		Q.Data[Q.rear] = data
		Q.rear = (Q.rear+1)%QueueMax
		return Q
	}
}
//出队
func OutQueue(Q *Queue) *Queue {
	//队空
	if Q.rear == Q.front {
		return Q
	}else {
		Q.Data[Q.front] = -1
		Q.front = (Q.front+1)%QueueMax
		return Q
	}
}

func main() {
	var myQueue *Queue
	myQueue = new(Queue)
	InitQueue(myQueue)
	for i := 0; i < 11; i++ {
		myQueue = EnQueue(myQueue,i+1)
	}
	fmt.Printf("EnQueue:%v,length:%d\n",myQueue,Leng(myQueue))
	for i := 0; i < 2; i++ {
		myQueue = OutQueue(myQueue)
	}
	fmt.Printf("OutQueue:%v,length:%d\n",myQueue,Leng(myQueue))
	/*
	data: 10
	队列已满，返回原队列！
	data: 11
	队列已满，返回原队列！
	EnQueue:&{9 0 [1 2 3 4 5 6 7 8 9 0]},length:9
	OutQueue:&{9 2 [-1 -1 3 4 5 6 7 8 9 0]},length:7

	Process finished with exit code 0
	*/
}
