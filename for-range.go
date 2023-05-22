package main

import "fmt"

type forRangeStruct struct {
	id      int
	orderNo string
}

func ForRange() {
	var slice []*forRangeStruct
	slice = append(slice, &forRangeStruct{
		id:      1,
		orderNo: "s1",
	})
	slice = append(slice, &forRangeStruct{
		id:      2,
		orderNo: "s2",
	})
	slice = append(slice, &forRangeStruct{
		id:      3,
		orderNo: "s3",
	})

	for k, v := range slice {
		fmt.Printf("k的值：%d，v=的值：%v，k的地址：%p，v的地址：%p\n", k, v, &k, &v)
	}
}
