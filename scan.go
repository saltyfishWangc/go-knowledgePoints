package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func mainScan() {
	//var (
	//	name    string
	//	age     int
	//	married bool
	//)

	//fmt.Scan(&name, &age, &married)

	//fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
	//fmt.Scanln(&name, &age, &married)
	//fmt.Printf("输入结果 name:%s age:%d married:%t\n", name, age, married)

	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	fmt.Print("请输入内容：")
	text, _ := reader.ReadString('\n') // 读到换行
	text = strings.TrimSpace(text)
	fmt.Printf("%#v\n", text)
}
