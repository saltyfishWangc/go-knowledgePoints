package main

import (
	"fmt"
	"reflect"
)

type MyInt int

func mainReflect() {
	//var a float64 = 6.4
	//
	//typeOf := reflect.TypeOf(&a)
	//fmt.Println("a的类型：", typeOf)
	//fmt.Println("a的类型：", typeOf.Kind())
	//
	//valueOf := reflect.ValueOf(&a)
	//fmt.Println("a的值：", valueOf)
	//fmt.Println("a的值：", valueOf.Kind())
	//
	//var my MyInt = 7
	//typeOfMy := reflect.TypeOf(my)
	//fmt.Println("my的类型：", typeOfMy)
	//fmt.Println("my的类型：", typeOfMy.Kind())
	//
	//valueOfMy := reflect.ValueOf(my)
	//fmt.Println("my的值：", valueOfMy)
	//fmt.Println("my的值：", valueOfMy.Kind())

	//var x float64 = 3.4
	////reflect_type(x)
	//reflect_set_value(x)
	//fmt.Println("main:", x)
	//reflect_set_value(&x)
	//fmt.Println("main:", x)

	//var data float64 = 5.0
	//of := reflect.ValueOf(data)
	//of := reflect.ValueOf(&data)
	//of.SetFloat(9.0)
	//of.Elem().SetFloat(9.0)
	//fmt.Println("data:", data)

	// 反射操作结构体
	user := ReflectUser{20, 21, "wangc"}
	//fmt.Printf("反射修改前：%v", user)
	//valueOfUser := reflect.ValueOf(&user)
	//valueOfUser.Elem().FieldByName("Id").SetInt(21)
	//fmt.Printf("反射修改后：%v", user)

	// 获取到每个字段的值
	//valueOfUser := reflect.ValueOf(user)
	//typeOfUser := reflect.TypeOf(user)
	//fmt.Println("user的类型是：%v", valueOfUser.Type())
	//fmt.Println("user的类型是：%v", typeOfUser)
	//for i := 0; i < valueOfUser.NumField(); i++ {
	//	structField := valueOfUser.Field(i)
	//	switch structField.Type().Kind() {
	//	case reflect.Int:
	//		fmt.Printf("第:%d个字段是：%v，类型是:%s, 值是:%d\n", i, typeOfUser.Field(i).Name, structField.Type().Name(), structField.Int())
	//	case reflect.Float64:
	//		fmt.Printf("第:%d个字段是：%v，类型是:%s, 值是:%f\n", i, typeOfUser.Field(i).Name, structField.Type().Name(), structField.Float())
	//	case reflect.String:
	//		fmt.Printf("第:%d个字段是：%v，类型是:%s, 值是:%s\n", i, typeOfUser.Field(i).Name, structField.Type().Name(), structField.String())
	//	}
	//}

	//valueOfUserTag := reflect.ValueOf(user)
	valueOfUserTag := reflect.ValueOf(&user).Elem()
	for i := 0; i < valueOfUserTag.NumField(); i++ {
		// 值信息中获取到的字段获取到的是字段的值，关于字段的结构(字段定义的信息：字段名、tag标签)还得看类型信息
		fmt.Printf("第:%d个字段：%v\n", i, valueOfUserTag.Field(i))
		fmt.Printf("第:%d个字段：%v\n", i, valueOfUserTag.Field(i).Type())
		fmt.Printf("第:%d个字段：%v\n", i, valueOfUserTag.Type().Field(i))
	}

}

type ReflectUser struct {
	Id   int    `json:"json-id" gorm:"gorm-id"`
	Age  int    `json:"json-age" gorm:"gorm-age"`
	Name string `json:"json-name" gorm:"gorm-name"`
}

func reflect_type(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Println("类型是：", t)
	k := t.Kind() //kind()可以获取具体类型
	fmt.Println("具体类型是：", k)
	switch k {
	case reflect.Float64:
		fmt.Println("a is float64")
	case reflect.String:
		fmt.Println("string")
	}
}

// reflect_set_value 反射修改值
func reflect_set_value(a interface{}) {
	v := reflect.ValueOf(a)
	k := v.Kind()
	switch k {
	case reflect.Float64:
		// 反射修改值
		//v.Set(9.0)
		v.SetFloat(6.9)
		fmt.Println("a is ", v)
		fmt.Println("a is ", v.Float())
	case reflect.Ptr:
		v.SetInt(9)
		fmt.Println(v.Pointer())
	}
}
