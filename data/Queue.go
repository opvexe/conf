package data

import "fmt"

//队列特性FIFO
// 入队：尾部插入
// 出队：头部删除

type QueueNode struct {
	Data interface{}
	Next *QueueNode
}

//创建队列
func (node *QueueNode)Create(data ...interface{})  {
	if node == nil|| len(data) == 0 {
		return
	}
	//循环添加数据
	for _,v := range data{
		newNode := new(QueueNode)
		newNode.Data = v
		newNode.Next = nil

		node.Next = newNode
		node = node.Next
	}
}

//打印
func (node *QueueNode)Print()  {
	if node == nil {
		return
	}
	if node.Data!=nil {
		fmt.Println(node.Data)
	}
	node.Next.Print()
}


//入队
func (node *QueueNode)EnQueue(data interface{})  {
	if node== nil||data==nil {
		return
	}
	newNode := new(QueueNode)
	newNode.Data = data
	newNode.Next = nil

	//循环找到尾结点
	for node.Next!=nil{
		node = node.Next
	}
	node.Next = newNode
}

//出队
func (node *QueueNode)DeQueue()  {
	if node ==nil {
		return
	}
	node.Next = node.Next.Next //将数据节点1更新为数据节点的第二个
}




