package data

import (
	"fmt"
	"reflect"
)

//单链表
type LinkNode struct {
	Data interface{}
	Next *LinkNode
}

//创建
func (node  *LinkNode)Create(data ...interface{})  {
	if node == nil|| len(data) == 0 {
		return
	}
	//循环赋值
	for _,v := range data{
		newNode := new(LinkNode)
		newNode.Data = v
		newNode.Next = nil

		node.Next = newNode
		node = node.Next
	}
}

//打印
func (node *LinkNode)Print()  {
	if node == nil {
		return
	}
	if node.Data!=nil {
		fmt.Print(node.Data," ")
	}
	node.Next.Print()
}

//长度
func (node *LinkNode)Length() int {
	if node ==nil {
		return -1
	}
	i:=0
	for node.Next!=nil{
		node = node.Next
		i++
	}
	return i
}

//头部插入法
func (node *LinkNode)InsertByHead(data interface{})  {
	if node == nil ||data == nil {
		return
	}
	newNode := new(LinkNode)
	newNode.Data = data
	newNode.Next = nil

	//先右后左
	newNode.Next = node.Next
	node.Next  = newNode
}

//尾部插入法
func (node *LinkNode)InsertByTail(data interface{})  {
	if node == nil || data == nil {
		return
	}
	newNode := new(LinkNode)
	newNode.Data = data
	newNode.Next = nil
	//循环获取尾节点
	for node.Next!=nil{
		node = node.Next
	}
	node.Next = newNode
}

//中间插入法
func (node *LinkNode)InsertByMiddle(index int,data interface{})  {
	if node == nil ||data == nil {
		return
	}
	if index<=0||index>node.Length() {
		return
	}
	preNode := node
	//循环
	for i:=0;i<index;i++{
		preNode = node
		node = node.Next
	}

	newNode := new(LinkNode)
	newNode.Data = data
	newNode.Next = nil

	newNode.Next = node
	preNode.Next = newNode
}

//按位置删除
func (node *LinkNode)DeleteByIndex(index  int)  {
	if node == nil {
		return
	}
	if index<=0||index>node.Length() {
		return
	}
	preNode := node
	//循环
	for i:=0;i<index;i++{
		preNode = node
		node = node.Next
	}
	preNode.Next = node.Next
	//销毁
	node.Data = nil
	node.Next = nil
	node = nil
}

//查询链表节点
func (node *LinkNode)Search(data interface{}) int {
	if node == nil {
		return -1
	}
	len := node.Length()
	for i:=1;i<=len;i++{
		node = node.Next
		if reflect.DeepEqual(node.Data, data) {
			return i
		}
	}
	return -1
}