package main

import (
	"fmt"
	"reflect"
)

func main() {

	//定义变量
	var a int = 10

	//调用
	baseReflect(a)
}

//类型转换
func aseReflect(i interface{}) {

	//转换成valueof
	refValue := reflect.ValueOf(i)

	//valueof 转换成interface
	inter :=refValue.Interface()

	//断言转换成int
	num := inter.(int)

	//打印
	fmt.Println("int:",num)

}
