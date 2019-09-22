package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Student struct {
	Name        string `label:"Person Name: " uppercase:"true"`
	Age         int    `label:"Age is: "`
	Sex         string `label:"Sex is: "`
	Description string
}

func main() {

	//初始化
	stu := Student{"sobot",20,"male","智齿科技"}

	//调用
	reflectTag(&stu)
}

func reflectTag(i interface{})  {

	//获取valueof
	reValu:=reflect.ValueOf(i)

	/*
		写法2

		typeValue := reflect.TypeOf(i)

		if typeValue.Kind() != reflect.Ptr||typeValue.Elem().Kind() != reflect.Struct {
			return
		}

		reValue:=reflect.ValueOf(i).Elem()
	*/

	//类型判断
	if reValu.Kind() !=reflect.Ptr ||reValu.Elem().Kind() !=reflect.Struct {
		return
	}

	////遍历
	for i := 0; i < reValu.Elem().NumField(); i++ {  //指针Elem()转换

		//转换成typeof
		reType :=reValu.Elem().Type().Field(i)


		//获取标签
		reTag :=reType.Tag

		tagValu:= reTag.Get("label")

		if tagValu =="" {

			tagValu = reType.Name +":"  //Description：
		}

		value := fmt.Sprintf("%v",reValu.Elem().Field(i))

		//转小写
		if reTag.Get("uppercase") == "true" {

			value = strings.ToUpper(value)
		}else {
			value = strings.ToLower(value)
		}

		//打印
		fmt.Println(tagValu + value)
	}

}
