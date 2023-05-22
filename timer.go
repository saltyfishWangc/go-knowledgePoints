package main

import (
	"fmt"
	"time"
)

func mainTimer() {
	timer := time.NewTimer(3 * time.Second)
	//timer.Reset(1 * time.Second)
	fmt.Println(time.Now())
	fmt.Println(<-timer.C)
}
