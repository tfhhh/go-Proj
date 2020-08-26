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
func (Q *Queue) Init() Queue {
	Q.rear = 0
	Q.front = 0
	Q.Data = make([]int, QueueMax)
	return Queue{Q.rear,Q.front,Q.Data}
}
//队长
func(Q *Queue) Length() int {
	return (Q.rear - Q.front + QueueMax)%QueueMax
}
//入队
func(Q *Queue) EnQueue(data int) bool {
	//队满
	if (Q.rear+1)%QueueMax == Q.front {
		return false
	} else {
		Q.Data[Q.rear] = data
		Q.rear = (Q.rear+1)%QueueMax
		return true
	}
}
//出队
func(Q *Queue) OutQueue() bool {
	//队空
	if Q.rear == Q.front {
		return false
	}else {
		Q.Data[Q.front] = 0
		Q.front = (Q.front+1)% QueueMax
		return true
	}
}

func main() {
	var myQueue Queue
	myQueue.Init()
	//入队，[1,2,3,4,0,0,0,0,0]
	for i := 1; i < 5; i++ {
		myQueue.EnQueue(i)
	}
	fmt.Printf("init queue:%v,length:%d\n",myQueue,myQueue.Length())
	myQueue.OutQueue()	//[0,2,3,4,0,0,0,0,0]
	myQueue.OutQueue()	//[0,0,3,4,0,0,0,0,0]
	fmt.Printf("my queue:%v,length:%d\n",myQueue,myQueue.Length())
}