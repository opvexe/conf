package main

import (
	"fmt"
	"reflect"
)

//定义对象
type Man struct {
	Name    string
	Age     int
	Company string
}

func main() {

	//初始化对象
	man := Man{"sobot", 25, "北京智齿科技"}

	//调用
	structReflect(man)
}

func structReflect(i interface{}) {

	//转换成valueof
	refValu := reflect.ValueOf(i)

	//valueof 转换成interface
	inter := refValu.Interface()

	//断言转换
	man := inter.(Man)

	//打印
	fmt.Println(man.Name)
	fmt.Println(man.Age)
	fmt.Println(man.Company)
}
