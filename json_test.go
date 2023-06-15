package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJson(t *testing.T) {
	//st := struct {
	//	name string
	//	age  int
	//}{
	//	name: "wangc",
	//	age:  2,
	//}
	//
	//marshal, _ := json.Marshal(st)
	//t.Log("json.Marshal输出：", marshal)
	//marshalIndent, _ := json.MarshalIndent(st, "hello ", "bye")
	//t.Log("json.MarshalIndent输出：", marshalIndent)

	f := func() {
		a := 2
		fmt.Println("hello: ", a)
	}
	fJson, _ := json.Marshal(f)
	t.Log(fJson)
}
