package main

import (
	"fmt"
	"reflect"
)

type User2 struct {
	Id   int
	Name string
	Age  int
}

func (u User2) Hello() {
	fmt.Println("Hello")
}

func (u User2) SeeYou(name string) {
	fmt.Println("SeeYou, ", name)
}

func (u User2) ReturnSeeYou(name string) string {
	return fmt.Sprintf("See You %s", name)
}

// func mainReflect2() {
func mainReflect2() {
	u := User2{1, "sz", 20}
	//Poni(&u)
	callMethod(u)
}

func Poni(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("类型：", t)
	fmt.Println("字符串类型：", t.Name())

	v := reflect.ValueOf(o)
	fmt.Println(v)

	for i := 0; i < t.NumField(); i++ {
		// 取每个字段
		f := t.Field(i)
		fmt.Printf("字段 %s : %v\n", f.Name, f.Type)

		// 获取字段的值信息
		// Interface()：获取字段对应的值
		val := v.Field(i).Interface()
		fmt.Println("val : ", val)
	}

	fmt.Println("=================方法=================")
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println("方法：", m.Name, "的类型是", m.Type)
	}
}

func callMethod(o interface{}) {
	v := reflect.ValueOf(o)
	// 获取方法
	m := v.MethodByName("SeeYou")
	// 构建一些参数
	args := []reflect.Value{reflect.ValueOf("wangc")}
	// 没参数的情况下：var args []reflect.Value
	// 调用方法
	m.Call(args)

	//t := reflect.TypeOf(o)
	t := reflect.ValueOf(o)
	//m2, ok := t.MethodByName("ReturnSeeYou")
	//if !ok {
	//	panic("不存在的方法")
	//}
	m2 := t.MethodByName("ReturnSeeYou")
	args2 := []reflect.Value{reflect.ValueOf("wangc")}
	callRes := m2.Call(args2)
	for _, val := range callRes {
		fmt.Println("返回结果：", val)
	}
}
