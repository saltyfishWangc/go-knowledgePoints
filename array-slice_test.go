package main

import (
	"fmt"
	"testing"
)

func TestArraySlice(t *testing.T) {
	ArraySlice()
}

func TestSliceReference(t *testing.T) {
	SliceReference()
}

func TestSliceInitiation(t *testing.T) {
	SliceInitiation()
}

func TestSliceInitiation2(t *testing.T) {
	var sli1 []int
	fmt.Printf("切片：%v,长度：%d, 容量：%d\n", sli1, len(sli1), cap(sli1))
	sli1 = append(sli1, 1)
	fmt.Printf("切片：%v,长度：%d, 容量：%d\n", sli1, len(sli1), cap(sli1))

	sli2 := make([]int, 0, 4)
	fmt.Printf("切片：%v,长度：%d, 容量：%d\n", sli2, len(sli2), cap(sli2))
	sli2 = append(sli2, 1)
	sli2 = append(sli2, 1)
	sli2 = append(sli2, 1)
	sli2 = append(sli2, 1)
	sli2 = append(sli2, 1)
	fmt.Printf("切片：%v,长度：%d, 容量：%d\n", sli2, len(sli2), cap(sli2))

	sli3 := make([]int, 4, 4)
	fmt.Printf("切片：%v,长度：%d, 容量：%d\n", sli3, len(sli3), cap(sli3))
	sli3 = append(sli3, 1)
	sli3 = append(sli3, 1)
	sli3 = append(sli3, 1)
	sli3 = append(sli3, 1)
	sli3 = append(sli3, 1)
	fmt.Printf("切片：%v,长度：%d, 容量：%d\n", sli3, len(sli3), cap(sli3))

}
