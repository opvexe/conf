package data

import "fmt"

//定义hashTbale
type HashTable struct {
	linkArr [7]Emplink //链表数组  7可以调整
}

//定义Emplink
type Emplink struct { //Emplink 不带表头，第一个节点就存放雇员
	Head *Emp
}

//定义emp雇员
type Emp struct {
	Id   int
	Name string
	Next *Emp
}

//给Hashtable 增加添加的方法
func (this *HashTable) Insert(emp *Emp) {
	//先确定把雇员放在哪个链表  ----> 散列函数
	linkNo := this.HashFun(emp.Id)
	//使用对应的链表添加
	this.linkArr[linkNo].Insert()
}

//添加员工的方法
func (this *Emplink) Insert(emp *Emp) {
	//保证是一个从小到大的顺序，编号 ---->
	cur := this.Head //这是辅助指针
	//定义一个复制指针 ，指向前面一个节点
	var pre *Emp = nil
	//如果当前就是一个空链表
	if cur == nil { //空链表
		this.Head = emp
		return
	}
	//如果不是一个空链表,给emp 找到对应的位置，插入
	//思路是 cur 与emp 比较 ，让pre 保持在cur前面
	for {
		if cur != nil {
			//比较
			if cur.Id >= emp.Id {		//id 重复
				//找到位置
				break
			}
			pre = cur
			cur = cur.Next
		} else {
			break
		}
	}
	//退出
	if cur == nil { //加到最后
		pre.Next = emp
		emp.Next = nil //插入尾部
	} else { //中间添加进去
		pre.Next = emp
		emp.Next = cur
	}
}

//显示所有员工
func (this *HashTable)ShowAll()  {
	for i := 0; i < len(this.linkArr); i++ {
		this.linkArr[i].ShowLink(i)
	}
}

//显示方法
func (this *Emplink)ShowLink(no int)  {
	if this.Head ==nil {
		fmt.Printf("链表%d为空",no)
		return
	}
	//遍历当前链表显示数据
	cur  := this.Head
	for {
		if cur!=nil{
			fmt.Printf("链表%d,雇员%d,雇员名字%s--->",no,cur.Id,cur.Name)
			cur = cur.Next
		}else {
			break
		}
	}
}

//求膜散列
func (this *HashTable) HashFun(id int) int {
	return id % 7 //对应链表的下标
}



func main() {

	var hashTable HashTable
	key  := ""
	id := 0
	name := ""

	for {
		fmt.Println("============= 雇员系统菜单============")
		fmt.Println("input 表示添加雇员")
		fmt.Println("show 表示显示雇员")
		fmt.Println("find 表示查询雇员")
		fmt.Println("exit 表示退出系统")
		fmt.Println("请输入你的选择")

		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("请雇员输入id")
			fmt.Scanln(&id)
			fmt.Println("请雇员名字")
			fmt.Scanln(&name)

			emp :=&Emp{id,name,nil}
			hashTable.Insert(emp)

		case "show":

		case "find":

		case "exit":

		default:
			fmt.Println("输入错误")
		}

	}
}
