package main

import "testing"

type A interface {
	hello() string
	say()
}

type B interface {
	hello() string
	say()
}

func TestInterface(t *testing.T) {
	var a A
	var b B
	t.Log("a是否等于b:", a == b)
}
