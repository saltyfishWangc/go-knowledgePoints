package main

import (
	"testing"
)

type mapUser struct {
	name string
	age  int
}

func TestMapCanBeAddress(t *testing.T) {
	mp := make(map[int]*mapUser)
	mu1 := &mapUser{name: "wangc", age: 25}
	mp[1] = mu1
	t.Log("修改前：", mp[1])
	mu1.name = "kk"
	t.Log("修改前：", mp[1])
	t.Log("查一个不存在的值：", mp[2])

	mp2 := make(map[int]string)
	t.Log("查一个不存在的值：", mp2[1])

	_, ok := mp2[2]
	if !ok {
		t.Log("不存在的值")
	}

}
