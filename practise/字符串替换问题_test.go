package practise

import (
	"strings"
	"testing"
	"unicode"
)

// 问题描述:
// 请编写⼀个⽅法，将字符串中的空格全部替换为“%20”。 假定该字符串有⾜够的空间存
// 放新增的字符，并且知道字符串的真实⻓度(⼩于等于1000)，同时保证字符串由【⼤⼩写的英⽂字⺟组成】。 给定⼀个string为原始的串，返回替换后的string。
func TestReplaceString(t *testing.T) {

}

func replaceString(s string) (string, bool) {
	if len([]byte(s)) > 1000 {
		return s, false
	}

	for _, v := range s {
		if string(v) != " " && !unicode.IsLetter(v) {
			return s, false
		}
	}
	return strings.Replace(s, " ", "%20", -1), true
}

type Param map[string]interface{}

type Show struct {
	Param
}

func TestGrammar(t *testing.T) {
	s := new(Show)
	s.Param["RMB"] = 100
}
