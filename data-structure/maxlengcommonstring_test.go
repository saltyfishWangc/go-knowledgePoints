package data_structure

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	str := []string{"flower", "flow", "flight"}
	t.Log(longestCommonPrefix(str))
}
