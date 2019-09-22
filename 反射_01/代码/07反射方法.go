package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	Name  string `json:"name"`
	Age int `json:"age"`
	Score float32
}

func (d Doctor)SetName()  {
	fmt.Println("SetName")
}

func (d Doctor)Add(i ,j int) int{
	return i+j
}

func main() {

	doct := Doctor{"sobot",20,90}

	reflectMethod(doct)
}

func reflectMethod(i interface{})  {

	refValu := reflect.ValueOf(i)

	if refValu.Kind() != reflect.Struct {  //不是指针类型
		return
	}

	numMeth:=refValu.NumField()

	fmt.Println("方法:",numMeth)

	//调用add 方法 ASII码顺序
	refValu.Method(1).Call(nil)

	//调用第1个
	var s []reflect.Value
	s = append(s,reflect.ValueOf(100))
	s = append(s,reflect.ValueOf(200))

	//返回是 Value
	sum := refValu.Method(0).Call(s)

	//获取切片第一个元素，断言
	fmt.Println(sum[0].Int())
}
