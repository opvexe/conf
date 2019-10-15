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

	// 在循环之外,形成一个闭环
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

//插入
func (node *CircularNode)Insert(index int,data interface{})  {
	if node == nil ||data == nil {
		return
	}
	if index<=0||index>node.Length() {
		return
	}
	//标记位
	start := node.Next
	//定义前一个节点
	preNode := node
	//循环
	for i:=0;i<index;i++{
		preNode = node
		node = node.Next
	}
	newNode := new(CircularNode)
	newNode.Data = data
	newNode.Next = node

	preNode.Next = newNode

	if index == 1 {
		for {
			if start == node.Next {		//找到尾节点
				break
			}
			node = node.Next
		}
		node.Next = newNode
	}
}

//删除
func (node  *CircularNode)Delete(index  int)  {
	if node == nil {
		return
	}
	if index<=0||index>node.Length() {
		return
	}
	// 定义标记位置
	start := node.Next

	// 定义preNode, 用来标记index 对应结点的前一个结点
	preNode := node
	// 循环找到index 对应结点
	for i := 0; i < index; i++ {
		preNode = node
		node = node.Next
	}
	if index == 1 {
		// 定义临时变量, 用来找寻尾结点
		temp := node
		for {
			if start == temp.Next {
				break
			}
			temp = temp.Next
		}							// 当循环结束时,temp中保存尾结点

		// 将原链表的尾结点,指向删除后的第一个数据结点.
		temp.Next = node.Next
	}

	preNode.Next = node.Next

	// 置空
	node.Data = nil
	node.Next = nil
	node = nil
}