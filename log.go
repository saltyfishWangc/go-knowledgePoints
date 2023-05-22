package main

import (
	"log"
	"os"
)

func mainLog() {
	logger1 := log.New(os.Stdout, "[wangc]", log.Llongfile|log.Ldate|log.Ltime)
	logger1.Println("你好, logger1")

	logger2 := log.New(os.Stdout, "[test]", log.Lshortfile|log.Ltime)
	logger2.Println("hello, logger2")
	logger1.Println("你好, logger1")
}
