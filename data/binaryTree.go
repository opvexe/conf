package data

import (
	"fmt"
	"reflect"
)

//二叉树
type BinaryTreeNode struct {
	Data interface{}
	LeftChild *BinaryTreeNode
	RightChild *BinaryTreeNode
}

//创建
func (node *BinaryTreeNode)Create()  {
	if node == nil {
		return
	}
	//根据图形创建二叉树上所有的节点
	node1 := BinaryTreeNode{1,nil,nil}
	node2 := BinaryTreeNode{2,nil,nil}
	node3 := BinaryTreeNode{3,nil,nil}
	node4 := BinaryTreeNode{4,nil,nil}
	node5 := BinaryTreeNode{5,nil,nil}
	node6 := BinaryTreeNode{6,nil,nil}
	node7 := BinaryTreeNode{7,nil,nil}

	//给根节点赋值
	node.Data = 0
	node.LeftChild = &node1
	node.RightChild = &node2

	node1.LeftChild = &node3
	node1.RightChild = &node4

	node2.LeftChild = &node5
	node2.RightChild = &node6

	node3.LeftChild = &node7
}

/*
  所有遍历 都是先左节点 在右节点
 */

//先序遍历   ----> DLR 根 ---左 ---- 右
func (node *BinaryTreeNode)Pre()  {
	if node == nil {
		return
	}
	fmt.Print(node.Data," ")
	node.LeftChild.Pre()
	node.RightChild.Pre()
}

//中序遍历  ----> LDR  左---根 ---- 右
func (node *BinaryTreeNode)Midd()  {
	if node == nil {
		return
	}
	node.LeftChild.Midd()
	fmt.Print(node.Data," ")
	node.RightChild.Midd()
}

//后序遍历   ----> LRD  左 ---- 右  ----根
func (node  *BinaryTreeNode)End()  {
	if node == nil {
		return
	}
	node.LeftChild.Midd()
	node.RightChild.Midd()
	fmt.Print(node.Data," ")
}

//二叉树深度（高度）
func (node *BinaryTreeNode)TreeHeight() int {
	if node == nil {
		return 0
	}
	//获取左节点高度
	leftHeight := node.TreeHeight()
	//获取右节点高度
	rightHeight := node.TreeHeight()
	//比较左子树和右子树的返回值，进行累加
	if leftHeight >rightHeight {
		leftHeight ++
		return leftHeight
	}else {
		rightHeight ++
		return rightHeight
	}
}

//获取叶子节点数
var i int = 0
func (node *BinaryTreeNode)LeafNum()  {
	if node == nil{
		return
	}
	if node.LeftChild ==nil&&node.RightChild==nil {		//叶子节点特性：没有后驱
		i++
	}
	node.LeftChild.LeafNum()
	node.RightChild.LeafNum()
}

//二叉树查找
func (node *BinaryTreeNode)Search(data interface{})  {
	if node == nil  {
		return
	}
	if reflect.DeepEqual(node.Data,data) {
		fmt.Print("二叉树查找:",node)
		return
	}
	node.LeftChild.Search(data)
	node.RightChild.Search(data)
}

//反转二叉树
func (node *BinaryTreeNode)Reverse()  {
	if node == nil {
		return
	}
	node.LeftChild,node.RightChild = node.RightChild,node.LeftChild
	node.LeftChild.Reverse()
	node.RightChild.Reverse()
}

//拷贝二叉树
func (node  *BinaryTreeNode)Copy() *BinaryTreeNode {
	if node == nil {
		return nil
	}
	leftChild := node.LeftChild.Copy()
	rightChild := node.RightChild.Copy()

	newNode := new(BinaryTreeNode)
	newNode.Data = node.Data
	newNode.LeftChild = leftChild
	newNode.RightChild = rightChild
	return newNode
}