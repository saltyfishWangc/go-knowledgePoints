package data_structure

import (
	"testing"
	string2 "wangc.com/go-knowledgePoints/data-structure/string"
)

func TestLongestCommonPrefix1(t *testing.T) {
	str := []string{"flower", "flow", "flight"}
	t.Log(string2.longestCommonPrefix(str))
}
