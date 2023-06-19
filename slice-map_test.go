package main

import "testing"

func TestSliceAndMap(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	changeSlice(s1)
	t.Log(s1)
	t.Logf("切片的长度：%d 容量：%d", len(s1), cap(s1))
	m1 := map[int]int{1: 1, 2: 2, 3: 3, 4: 4}
	changeMap(m1)
	t.Log(m1)
	t.Logf("map的长度：%d", len(m1))
}

// 切片只是一个结构体，在接口中是值传递，如果是修改元素值，会影响外面的切片，因为他们底层是指向的同一个数组地址。但是如果是改变了切片容量，那就会发生拷贝，不会影响外面的切片
func changeSlice(s []int) {
	s[0] = 2
	s = append(s, 7, 8, 9, 10)
}

func changeMap(m map[int]int) {
	m[1] = 2
	m[5] = 5
	m[6] = 6
}

func TestSliceArr(t *testing.T) {
	str1 := []string{"a", "b", "c"}
	str2 := str1[1:]
	str2[1] = "new"
	t.Log(str1)
	str2 = append(str2, "z", "x", "y")
	t.Log(str1)
	t.Log(str2)
}

func TestSliceAndArrEqual(t *testing.T) {
	t.Log([...]string{"1"} == [...]string{"1"}) // 数组只能与长度、类型相同的其他数组比较
	//t.Log([]string{"1"} == []string{"1"})       // 切片之间不能直接比较
}
