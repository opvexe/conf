package main

import (
	"fmt"
)

func main() {

	slice := []int{9, 1, 5, 6, 10, 8, 3, 7, 2, 4, 11}

	//buleSort(slice)

	//selectSort(slice)

	//insertSort(slice)

	//headSort(slice)

	//qiucklySort(slice,0,len(slice)-1)

	fmt.Println(slice)
}


/*
 冒泡排序  [相邻两元素之间比较]
*/
func buleSort(slice []int) {

	for i := 0; i < len(slice)-1; i++ {

		for j := 0; j < len(slice)-1-i; j++ {

			if slice[j] > slice[j+1] {

				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
		fmt.Println("排序顺序:", slice)
	}
	fmt.Println("冒泡排序:", slice)
}

/*
 选择排序
*/

func selectSort(slice []int) {

	for i := 0; i < len(slice); i++ {

		p := i

		for j := i + 1; j < len(slice); j++ {

			if slice[p] > slice[j] {

				p = j
			}
		}

		if p != i {

			slice[p], slice[i] = slice[i], slice[p]
		}
		fmt.Println("排序顺序:", slice)
	}

	fmt.Println("选择排序:", slice)
}

/*
 插入排序
*/
func insertSort(slice []int) {

	for i := 1; i < len(slice); i++ {
		for j := i; j > 0; j-- {
			if slice[j] > slice[j-1] {
				break
			}
			slice[j], slice[j-1] = slice[j-1], slice[j]
		}
	}

	fmt.Println("插入排序:", slice)
}

/*
 快速排序
*/

func qiucklySort(s []int,left,right int){

	if left >=right {
		return
	}

	value := s[left]
	k:= left

	for i := left+1;i<=right;i++{
		if s[i] < value{
			s[k] = s[i]
			s[i] = s[k+1]
			k++
		}
	}
	s[k] = value
	qiucklySort(s,left,k-1)
	qiucklySort(s,k+1,right)
}

/*
 堆排序
*/

func headSort(tree []int) {

	length := len(tree)

	for i := length/2 - 1; i >= 0; i-- {
		nodeSort(tree, i, length-1)
	}

	// 次数tree已经是个大根堆了。只需每次交换根节点和最后一个节点，并减少一个比较范围。再进行一轮比较
	for i := length - 1; i > 0; i-- {
		// 如果只剩根节点和左孩子节点，就可以提前结束了
		if i == 1 && tree[0] <= tree[i] {
			break
		}
		// 交换根节点和比较范围内最后一个节点的数值
		tree[0], tree[i] = tree[i], tree[0]
		// 这里递归的把较大值一层层提上来
		nodeSort(tree, 0, i-1)
	}
}

func nodeSort(tree []int, startNode, latestNode int) {
	var largerChild int
	leftChild := startNode*2 + 1
	rightChild := leftChild + 1

	// 子节点超过比较范围就跳出递归
	if leftChild >= latestNode {
		return
	}

	// 左右孩子节点中找到较大的，右孩子不能超出比较的范围
	if rightChild <= latestNode && tree[rightChild] > tree[leftChild] {
		largerChild = rightChild
	} else {
		largerChild = leftChild
	}

	// 此时startNode节点数值已经最大了，就不用再比下去了
	if tree[largerChild] <= tree[startNode] {
		return
	}
	fmt.Println(tree)
	// 到这里发现孩子节点数值比父节点大，所以交换位置，并继续比较子孙节点，直到把大鱼捞上来
	tree[startNode], tree[largerChild] = tree[largerChild], tree[startNode]
	nodeSort(tree, largerChild, latestNode)
}
