package practise

import (
	"strings"
	"testing"
)

// 问题描述
// 给定两个字符串，请编写程序，确定其中⼀个字符串的字符重新排列后，能否变成另⼀
// 个字符串。 这⾥规定【⼤⼩写为不同字符】，且考虑字符串重点空格。给定⼀个string
// s1和⼀个string s2，请返回⼀个bool，代表两串是否重新排列后可相同。 保证两串的
// ⻓度都⼩于等于5000。
func TestIsTwoStringSameAfterSorted(t *testing.T) {
	s1 := "你好n"
	s2 := "n好你"
	t.Log(isTwoStringSameAfterSorted(s1, s2))
}

// 只需要遍历一个字符串，判断其每个字符在两个字符串中出现的次数是不是一样得，如果不一样，那排序后不同。如果都是一样的，那排序后肯定是相同的两个字符串
func isTwoStringSameAfterSorted(s1, s2 string) bool {
	rune1 := []rune(s1)
	rune2 := []rune(s2)

	if len(rune1) > 5000 || len(rune2) > 5000 || len(rune1) != len(rune2) {
		return false
	}

	for _, v := range rune1 {
		if strings.Count(s1, string(v)) != strings.Count(s2, string(v)) {
			return false
		}
	}
	return true
}
