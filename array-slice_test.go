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

func TestEmptySlice(t *testing.T) {
	// 这里只是声明了一个切片，但是没有实例化，所以会报下标越界
	//var s1 []int
	//s1[0] = 1

	// // 这里只是声明了一个切片，但是没有实例化，所以会报下标越界
	//s1 := []int{}
	//s1[0] = 1

	//s1 := make([]int, 0, 0)
	//s1[0] = 1

	//s1 := make([]int, 0, 1)
	//s1[0] = 1

	s1 := make([]int, 1, 1)
	s1[0] = 1
	t.Log("切片s1：", s1)

	var s2 []int
	s2 = append(s2, 1)
	t.Log("切片s2：", s2)

	s3 := []int{}
	s3 = append(s3, 1)
	t.Log("切片s3：", s3)
}

// TestSliceForParam
// 切片其实只是一个结构体，只是里面有个指针指向数组。切片作为函数参数传参还是值传递
func TestSliceForParam(t *testing.T) {
	s1 := make([]int, 2)
	t.Log("初始化的值：", s1)
	sliceTest1(s1)
	t.Log("函数处理后的值：", s1)

	s := []int{1, 1, 1}
	newS := myAppend(s)
	t.Log("初始化s的值：", s)
	t.Log("切片作为函数参数传递后newS的值：", newS)

	s = newS

	myAppendPtr(&s)
	t.Log("切片指针作为函数参数传递后s的值：", s)

	s2 := make([]int, 1, 4)
	newS2 := myAppend(s2)
	t.Log("初始化s的值：", s2)
	t.Log("切片作为函数参数传递后newS2的值：", newS2)
}

func sliceTest1(s1 []int) {
	s1[0] = 1
}

func myAppend(s []int) []int {
	s = append(s, 100)
	return s
}

func myAppendPtr(s *[]int) {
	*s = append(*s, 100)
}

// TestSliceResize
// 切片扩容：从案例可以看出，切片扩容并不是像网上说的那样直接就是2倍，而是分情况的。可看源码：/runtime/slice.go中的growslice()
func TestSliceResize(t *testing.T) {
	s1 := []int{1, 2}
	t.Logf("初始化切片的容量：%d", cap(s1))
	for i := 0; i < 3; i++ {
		s1 = append(s1, i)
		t.Logf("添加元素后切片的容量：%d", cap(s1))
	}

	s2 := []int{1, 2}
	t.Logf("初始化切片的容量：%d", cap(s2))
	s2 = append(s2, 4, 5, 6)
	t.Logf("批量添加元素后切片的容量：%d", cap(s2))
}
