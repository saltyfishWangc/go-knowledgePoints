package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Username string
	Age      int
}

func main() {
	// json不能序列化切片吗
	slice := make([]User, 2)

	user1 := User{
		"wangc",
		1,
	}

	slice = append(slice, user1)

	jsonStr, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("json序列化出错")
		os.Exit(0)
	}
	fmt.Println("json串:", string(jsonStr))

	var arr [2]int
	arr[0] = 5
	arr[1] = 6

	sliceInt := make([]int, 2)
	sliceInt[0] = 1
	sliceInt[1] = 2
	sliceInt[2] = 3

	arrayInt := new([5]int)
	arrayInt[0] = 1
	arrayInt[1] = 2
	arrayInt[2] = 3
}
