package main

import (
	"fmt"
	"reflect"
)

const DefineTagKey = "ReflectTag"

type Person struct {
	Name string `ReflectTag:"-"`
	Age  int    `ReflectTag:"age,omitempy"`
	Id   int    `ReflectTag:"idx,string"`
	Sex  string
}

func main() {

	//初始化
	person := Person{"sobot", 20, 10086, "M"}

	//获取valueof
	reVlu := reflect.ValueOf(person)

	//获取typeof
	reType := reflect.TypeOf(person)

	for i := 0; i < reVlu.NumField(); i++ {

		//获取成员对象
		strField := reType.Field(i)

		//获取成员Name
		name := strField.Name

		//获取成员类型
		typ := strField.Type

		//获取tag
		tagVau := strField.Tag.Get(DefineTagKey)

		fmt.Printf("%v --%v -- %v \n", name, typ, tagVau)
	}
}
