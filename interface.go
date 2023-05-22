package main

import "fmt"

type Person interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

// WashingMachine 洗衣机
type WashingMachine interface {
	wash()
	dry()
}

// dryer 甩干器
type dryer struct{}

// dry 实现WashingMachine接口的dry()方法
func (d dryer) dry() {
	fmt.Println("甩一甩")
}

// haier 海尔洗衣机
type haier struct {
	dryer
}

// wash 实现实现WashingMachine接口的wash()方法
func (h haier) wash() {
	fmt.Println("洗一洗")
}

func mainInterface() {
	var peo Person = &Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
