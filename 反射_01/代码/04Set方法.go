package main

import (
	"fmt"
	"reflect"
)

func main() {

	//声明变量
	var a int = 10

	//调用
	setReflect(&a)

	//打印
	fmt.Println(a)
}

/*
  Elem() 操作 类似:
	var  a int  = 10
	var  p  *int
	p = &a
  SetInt:
   *p = 1000
*/

func setReflect(i interface{}) {

	//获取valueof
	reValu := reflect.ValueOf(i)

	//查看类型
	fmt.Printf("kind:%v\n", reValu.Kind())

	//类型判断
	if reValu.Kind() == reflect.Ptr && reValu.Elem().Kind() == reflect.Int {

		reValu.Elem().SetInt(1000)

	}
}
