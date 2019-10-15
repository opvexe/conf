package data

/*
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)


//自定义切片
type Slice struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}


//创建切片
func (s *Slice)Create(l int,c int,data ...interface{})  {
	if s == nil|| len(data) == 0 {
		return
	}

	if l < 0 || c < 0 || l > c || len(data) < c {
		return
	}

	//初始化len/cap成员
	s.Len = l
	s.Cap = c

	//申请内存空间,保存在s.data中  相当于保存着vodi * 指针
	s.Data = C.malloc(C.size_t(c) * 8)

	//将指针转换成数据 可计算类型
	p := uintptr(s.Data)

	//循环获取用户传入的数据，依次 写入s.data

	for _, v := range data {
		//将数值类型的p 转换成指针
		*(*int)(unsafe.Pointer(p)) = v //具体化数据类型  强制类型转换
		p += 8                         //移动指针
	}
}

//打印
func (s *Slice) Print() {

	if s == nil || s.Data == nil {
		return
	}

	fmt.Println(s.Len, s.Cap)

	//讲地址转换成数值
	p := uintptr(s.Data)

	//循环读取
	for i := 0; i < s.Len; i++ {
		fmt.Print(*(*int)(unsafe.Pointer(p)))
		p += 8 //指针后移
	}
}


//追加切片元素
func (s *Slice) Append(data ...int) {

	if s == nil {
		return
	}

	for s.Len+len(data)>s.Cap{
		//扩容
		s.Data = C.realloc(s.Data,C.size_t(s.Cap)*2*8)
		s.Cap *=2
	}

	p := uintptr(s.Data)


	p +=uintptr(s.Len) *8

	for  _,v := range data{
		*(*int)(unsafe.Pointer(p)) = v
		p += 8
	}

	s.Len +=len(data)
}


//获取切片元素
func (s *Slice)GetData(index int) int {
	if s ==nil ||s.Data ==nil{
		return -1
	}

	if index<0||index>s.Len {
		return -1
	}

	//地址----> 转换值
	p := uintptr(s.Data)

	p += uintptr(index)*8

	return *(*int)(unsafe.Pointer(p))
}


//查找切片元素 根据值查找对应的切片下标
func (s *Slice)Search(data int) int {
	if s == nil || s.Data == nil {
		return  -1
	}

	p := uintptr(s.Data)
	for  i:= 0 ;i<s.Len;i++{
		if *(*int)(unsafe.Pointer(p)) == data{
			return i
		}
		p +=8
	}
	return -1
}

//删除切片元素  根据下标删除指定元素
func (s *Slice)Delete(index  int)  {
	if s == nil || s.Data == nil {
		return
	}
	if index<0 || index>s.Len {
		return
	}

	p := uintptr(s.Data)

	//后移  最后一个与倒数第二个调换
	p += uintptr(index)*8

	atfer := p

	//后移
	for i:=index;i<s.Len;i++{
		atfer += 8
		*(*int)(unsafe.Pointer(p)) = *(*int)(unsafe.Pointer(atfer))
		p +=8
	}

	s.Len -= 1
}

//插入元素到切片中
func (s *Slice)Insert(data ,index int)  {
	if s == nil || s.Data == nil {
		return
	}
	if index<0 || index>s.Len {
		return
	}

	p:=uintptr(s.Data)

	//保存index 内存地址
	p +=uintptr(index)*8

	//末尾内存地址
	temp := uintptr(s.Data) +uintptr(s.Len)*8

	for i:=s.Len;i>index;i--{
		//前一个元素给后一个元素赋值
		*(*int)(unsafe.Pointer(temp)) = *(*int)(unsafe.Pointer(temp-8))
		temp -=8
	}

	*(*int)(unsafe.Pointer(p)) = data
	s.Len++
}
