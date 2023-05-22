package main

import "fmt"

func mainByte() {
	str := "wangc你好"
	fmt.Printf("字节长度:%d\n", len(str))

	// str[1] 对应的是a，打印出来的是对应的ascii码值
	fmt.Println(str[1])

	strBytes := []byte(str)
	fmt.Printf("字符串转化为字节数组后的长度:%d\n", len(strBytes))
	strBytes[5] = 'M'
	// 下标为5的位置是你这个字符的第一个字节下标，所以你就被改变了
	fmt.Printf("修改字节数组某个字节后新的字符串:%s\n", string(strBytes))

	strRune := []rune(str)
	fmt.Printf("字符串转化为字符数组后的长度:%d\n", len(strRune))
	strRune[5] = 'M'
	fmt.Printf("修改字符数组某个字符后新的字符串:%s\n", string(strRune))
}
