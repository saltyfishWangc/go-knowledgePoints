package data_structure

import "testing"

func TestNewSkipList(t *testing.T) {
	sl := NewSkipList()
	sl.Add(1, "wangc")
	sl.Add(2, "aba")
	//sl.Add(3, "twelve")
	//sl.Add(4, "chao")
	//sl.Add(5, "juan")
	//sl.Add(6, "kuang")
	//sl.Add(7, "twelve")
	//sl.Add(8, "eleven")
	//sl.Add(9, "twelve")
	//sl.Add(10, "eleven")
	t.Logf("now skip list: %d", sl.Len())
	t.Logf("now skip list level : %d", sl.level)
	sl.Print()
	//t.Logf("get %d from skip : %+v", 11, sl.Search(11))
	//sl.Remove(11)
	//t.Logf("now skip list: %d", sl.Len())
	//t.Logf("now skip list level : %d", sl.level)
	//t.Logf("get %d from skip : %+v", 11, sl.Search(11))
	//
	//t.Logf("now skip list: %+v", sl)
}
