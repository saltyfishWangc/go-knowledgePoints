package main

import "testing"

const (
	a = iota
	b = iota
)

const (
	name = "menglu"
	age  = 12
	c    = iota // 在这个常量声明块中，iota被用于第3行，所以c的值是2
	d    = iota
)

// iota叫常量生成器，用于生成一组相似规则初始化的常量，但是不用每行都写一遍初始化表达式。在一个const声明语句中，在【第一个声明的常量所在的行，iota将会被置为0】，然后在每一个有常量声明的行加一
func TestIota(t *testing.T) {
	t.Log(a)
	t.Log(b)
	t.Log(c)
	t.Log(d)
}
