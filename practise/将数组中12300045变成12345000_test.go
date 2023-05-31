package practise

import (
	"fmt"
	"testing"
)

func TestChangeArrayNumber(t *testing.T) {
	s1 := []int{1, 2, 3, 0, 0, 0, 4, 5}
	t.Log(changeArrayNumber(s1))
}

func changeArrayNumber(arr []int) []int {
	var newSlice []int
	// 统计0的个数
	var zeroCount int
	for _, v := range arr {
		if v != 0 {
			newSlice = append(newSlice, v)
		} else {
			zeroCount += 1
		}
	}
	fmt.Println("去除0之后的切片：", newSlice)
	// 定义一个全是0的切片
	zeroSlice := make([]int, zeroCount)
	return append(newSlice, zeroSlice...)
}
