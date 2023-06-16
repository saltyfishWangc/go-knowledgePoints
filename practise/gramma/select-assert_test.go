package gramma

import (
	"fmt"
	"testing"
)

type Student struct {
	Name string
}

func TestSelectAssert(t *testing.T) {
	stu := &Student{Name: "wangc"}
	typeAssertInSelect(stu)
}

// 关于类型断言在select中的使用：
// 如果case后面只有一个类型，那么msg的类型就转成了对应的类型
// 如果case后面有多个类型，那么msg的类型还是interface{}，那它想转成对应的类型，就必须得用v.(*Student)或者v.(Student)了
// 注意：只有一个变量是interface{}类型时，才能用v.(type)，所以会报语法错误说v是一个non-interface类型
func typeAssertInSelect(v interface{}) {
	switch msg := v.(type) {
	//case *Student, Student:
	//	fmt.Println(msg.Name)
	case *Student:
		fmt.Println(msg.Name)
	case Student:
		fmt.Println(msg.Name)
	}
}

type People interface {
	Show()
}
type Student1 struct{}

func (stu *Student1) Show() {}

func live() People {
	var stu *Student1
	return stu
}

func TestInterface(t *testing.T) {
	a := live()
	if a == nil {
		t.Log("AAAAAAA")
	} else {
		t.Log("BBBBBBB")
	}
}
