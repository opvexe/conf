package data

import "fmt"

//循环列表
//尾节点指向头节点

type CircularNode struct {
	Data interface{}
	Next *CircularNode
}

//创建循环列表
func (node *CircularNode)Create(data ...interface{})  {
	if node == nil|| len(data) == 0 {
		return
	}
	//保存头节点
	head := node
	//循环
	for _,v := range data{
		newNode := new(CircularNode)
		newNode.Data = v
		newNode.Next = nil

		node.Next = newNode
		node = node.Next
	}

	//循环结束时，将尾节点指向头节点
	node.Next = head.Next
}

//打印
func (node *CircularNode)Print()  {
	if node == nil {
		return
	}
	head := node //标记位
	for {
		node = node.Next
		if node.Data!=nil {
			fmt.Println(node.Data)
		}
		if node.Next == head.Next {
			break
		}
	}
}

//获取长度
func (node *CircularNode)Length() int {
	if node ==nil {
		return -1
	}
	start := node.Next
	i:= 0
	for {
		node = node.Next
		i++
		if node.Next == start {
			break
		}
	}
	return i
}
