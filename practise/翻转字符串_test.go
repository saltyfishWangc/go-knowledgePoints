package practise

import "testing"

// 问题描述:
// 请实现⼀个算法，在不使⽤【额外数据结构和储存空间】的情况下，翻转⼀个给定的字符串(可以使⽤单个过程变量)。
// 给定⼀个string，请返回⼀个string，为翻转后的字符串。保证字符串的⻓度⼩于等于5000。
func TestReverseString(t *testing.T) {
	s := "ni你好吗ni"
	t.Log(reverseString(s))
}

func reverseString(s string) (string, bool) {
	runeArr := []rune(s)
	if len(runeArr) > 5000 {
		return "", false
	}
	for i := 0; i < len(runeArr)/2; i++ {
		runeArr[i], runeArr[len(runeArr)-i-1] = runeArr[len(runeArr)-i-1], runeArr[i]
	}
	return string(runeArr), true
}
