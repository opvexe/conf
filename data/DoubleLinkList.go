package data

import "golang.org/x/tools/go/ssa/interp/testdata/src/fmt"

//双向链表
type DoubleLinkList struct {
	Data interface{}
	Prev *DoubleLinkList
	Next *DoubleLinkList
}

//创建双向链表
func (node *DoubleLinkList)Create(data ...interface{})  {
	if node == nil|| len(data) == 0 {
		return
	}
	for _,v := range data{
		newNode := new(DoubleLinkList)
		newNode.Data = v
		newNode.Prev = node
		newNode.Next = nil

		node.Next = newNode
		node = node.Next
	}
}

//打印
func (node *DoubleLinkList)Print()  {
	if node == nil {
		return
	}
	if node.Data !=nil {
		fmt.Print(node.Data," ")
	}
	node.Next.Print()
}

//获取双向链表的长度
func (node *DoubleLinkList)Length() int {
	if node == nil {
		return -1
	}
	i:= 0
	for node.Next !=nil{
		i++
		node = node.Next
	}
	return i
}

//按位置插入
func (node *DoubleLinkList)InsertByIndex(index int,data interface{})  {
	if node == nil ||data == nil {
		return
	}
	if index<=0||index>node.Length() {
		return
	}
	//尾部插入
	if index >node.Length() {
		//获取尾结点
		for node.Next !=nil{
			node = node.Next
		}

		newNode := new(DoubleLinkList)
		newNode.Data = data
		newNode.Prev = node
		newNode.Next = nil

		//将尾节点指向头节点
		node.Next = newNode
	}

	//中间或是头部插入
	preNode := node		//保存Index对应的前一个节点
	for i:=0;i<index;i++{
		preNode = node
		node = node.Next
	}

	newNode := new(DoubleLinkList)
	newNode.Data = data
	newNode.Prev = preNode
	newNode.Next = node

	//双向链表 index节点
	node.Prev = newNode
	preNode.Next = newNode
}

//删除节点
func (node *DoubleLinkList)Delete(index int)  {
	if node == nil {
		return
	}
	if index <=0||index >node.Length(){
		return
	}
	preNode := node
	//循环
	for i:=0;i<index;i++{
		preNode = node
		node = node.Next
	}

	if node.Next !=nil {
		//preNode的next 指针，指向node 的右节点
		preNode.Next = node.Next
		//node 的后一个节点的prev 指针，指向prenode
		node.Next.Prev = preNode
	}else {
		preNode = nil
	}
	node.Data = nil
	node.Prev = nil
	node.Next = nil
	node= nil
}