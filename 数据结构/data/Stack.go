package data

import "fmt"

//栈的特性： FILO 先进后
//压栈：头插入
//出栈：头删

type StackNode struct {
	Data interface{}
	Next *StackNode
}

//创建栈
func (node *StackNode)Create(data ...interface{})  {
	if node == nil|| len(data) ==0 {
		return
	}
	//创建结点，保存新结点的下一个节点
	var nextNode *StackNode
	//循环
	for _,v := range data{
		newNode := new(StackNode)
		newNode.Data = v
		newNode.Next = nil

		node = newNode  //将新结点，设置为头节点
		newNode.Next = nextNode //将头节点的next指向下一个节点
		nextNode = node  //更新当前结点为下一个节点
	}
}

//打印
func (node *StackNode)Print()  {
	if node == nil {
		return
	}
	for node!=nil{
		if node.Data !=nil {
			fmt.Println(node.Data)
		}
		node = node.Next
	}
}

//获取长度
func (node *StackNode)Length() int {
	if node == nil {
		return  -1
	}
	i:=0
	for node!=nil{
		node = node.Next
		i++
	}
	return i
}

//压栈
func (node *StackNode)Push(data interface{})  {
	if node == nil ||data == nil {
		return
	}
	newNode := new(StackNode)
	newNode.Data = data
	newNode.Next = nil

	//将新结点置为头节点
	newNode.Next =node
}


//出栈
func (node *StackNode)Pop()  {
	if node !=nil {
		return
	}
	node =node.Next
}

