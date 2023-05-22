package main

import (
	"reflect"
	"testing"
)

// TestDeepEqual reflect.DeepEqual是比较两个变量的值的，通过反射，如果字段和对应的值都相同，则相等
func TestDeepEqual(t *testing.T) {
	mapA := make(map[int]int, 3)
	mapA[1] = 1
	mapA[2] = 2
	mapA[3] = 3

	mapB := make(map[int]int, 3)
	mapB[1] = 1
	mapB[2] = 2
	mapB[3] = 3

	t.Logf("mapA的地址：%p mapB的地址：%p", mapA, mapB)
	t.Log("mapA是否与mapB相等：", reflect.DeepEqual(mapA, mapB))
	t.Log("mapA是否与mapB相等：", reflect.DeepEqual(&mapA, &mapB))
}
