package main

import (
	"fmt"
	"reflect"
)

func main() {

	//声明变量
	var a int = 10
	var b bool = true
	var c string = "sobot"
	var d float64 = 90.8

	//调用方法
	kindReflect(a)
	kindReflect(b)
	kindReflect(c)
	kindReflect(d)
}

/*
 结构体类型只能类型断言
*/
func kindReflect(i interface{}) {

	//转换成valueof
	reValu := reflect.ValueOf(i)

	//获取kind
	kind := reValu.Kind()

	//类型判断
	switch kind {
	case reflect.Int:
		fmt.Println("int:", reValu.Int())
	case reflect.Bool:
		fmt.Println("bool:", reValu.Bool())
	case reflect.String:
		fmt.Println("string:", reValu.String())
	case reflect.Float64:
		fmt.Println("float", reValu.Float())
	}
}
