package main

import (
	"fmt"
)

type Tree struct {
	Data  int	//值
	Left  *Tree	//左节点
	Right *Tree	//右节点
}

func Insert(d int) *Tree {
	return &Tree{d,nil,nil}
}
//先序遍历
func  PreOrder(T *Tree) bool  {
	if T == nil {
		return false
	}else {
		fmt.Printf("%d ",T.Data)	//根
		PreOrder(T.Left)				//左
		PreOrder(T.Right)				//右
		return true
	}
}
//中序遍历
func InOrder(T *Tree) bool {
	if T == nil {
		return false
	}else {
		InOrder(T.Left)
		fmt.Printf("%d ",T.Data)
		InOrder(T.Right)
		return true
	}
}
//后序遍历
func PostOrder(T *Tree) bool {
	if T == nil {
		return false
	}else {
		PostOrder(T.Left)
		PostOrder(T.Right)
		fmt.Printf("%d ",T.Data)
		return true
	}
}
//节点数
func count(T *Tree) int {
	if T == nil {
		return 0
	}else if T.Left == nil && T.Right == nil {
		return 1
	}else {
		return 1 + count(T.Left) + count(T.Right)	//根节点+左子树节点+右子树节点
	}
}
//判定二叉排序树
func Order(T *Tree) bool {
	if T ==	nil {
		return false
	}else {
		return T.Left.Data < T.Data && T.Right.Data > T.Data
	}
}
//二叉排序树的查找
func Search(T *Tree,e int) bool {
	if T == nil {
		return false
	}else {
		if e == T.Data {
			return true
		}else {
			if e < T.Data {
				return Search(T.Left,e)
			}else {
				return Search(T.Right,e)
			}
		}
	}
}
//高度
func Higher(T *Tree) int {
	if T == nil {
		return 0
	} else {
		LeftHigh := Higher(T.Left)
		RightHigh := Higher(T.Right)
		//找左右子树中高度大的
		if LeftHigh <= RightHigh {
			return RightHigh+1
		}else {
			return LeftHigh+1
		}
	}
}
//平衡二叉树的判定
func IsBalanced(T *Tree) bool {
	if T == nil {
		return true
	} else {
		lHigh := Higher(T.Left)
		rHigh := Higher(T.Right)
		x := lHigh - rHigh
		if x >= 2 || x <= -2 {
			return false
		}else {
			return IsBalanced(T.Left) && IsBalanced(T.Right)	//左右子树都为平衡二叉树为平衡
		}
	}
}

func main() {
	/*
	 1
   2   3
  4	5
	 */
	var MyNode Tree
	MyNode.Data = 1
	MyNode.Left = Insert(2)
	MyNode.Right = Insert(3)
	MyNode.Left.Left = Insert(4)
	MyNode.Left.Right = Insert(5)
	MyNode.Right.Left = Insert(6)
	count := count(&MyNode)
	fmt.Println("count:",count)
	fmt.Println("MyNode order?",Order(&MyNode))	//false
	fmt.Println("MyNode avl?",IsBalanced(&MyNode))

	PreOrder(&MyNode)	//12453
	fmt.Println()
	InOrder(&MyNode)		//42513
	fmt.Println()
	PostOrder(&MyNode)	//45231
	fmt.Println()
	fmt.Println("-----------------------")

	var MyOrder Tree
	MyOrder.Data = 45
	MyOrder.Left = Insert(12)
	MyOrder.Right= Insert(53)
	MyOrder.Left.Left = Insert(3)
	MyOrder.Left.Right = Insert(37)
	MyOrder.Left.Right.Left = Insert(24)
	fmt.Println("MyOrder order?",Order(&MyOrder))	//true
	InOrder(&MyOrder) //中序遍历MyOrder为递增有序序列[3 12 24 37 45 53]
	fmt.Println()
	fmt.Println("MyOrder avl?",IsBalanced(&MyOrder))	//false
	fmt.Println("MyOrder Higher:",Higher(&MyOrder))	//4
	fmt.Println("search 33:",Search(&MyOrder,37))	//true
	fmt.Println("search 99",Search(&MyOrder,99))	//false
}