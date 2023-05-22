package main

import "testing"

// TestNilIsEqualNil 两个nil比较，不仅比较值，还要比较类型
// 如果两个显式类型的变量做比较，必须类型相同，否则代码直接报错说类型不匹配，无法编译
func TestNilIsEqualNil(t *testing.T) {
	var int32Nil *int32 = nil

	var int64Nil interface{} = nil

	t.Log("int32Nil 是否等于 nil：", int32Nil == nil)
	t.Log("int64Nil 是否等于 nil：", int64Nil == nil)
	t.Log("int32Nil 是否等于 int64Nil：", int32Nil == int64Nil)

}
