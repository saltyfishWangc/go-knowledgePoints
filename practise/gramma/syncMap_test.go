package gramma

import (
	"sync"
	"testing"
)

func TestSyncMap(t *testing.T) {
	var m sync.Map
	m.LoadOrStore("a", 1)
	m.Delete("a")
	//fmt.Println(len(m))

	c := make(map[string]interface{})
	_ = len(c)
}
