package string

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	t.Log(longestCommonPrefix(strs))
}

// longestCommonPrefix 查找字符串数组中的最长公共前缀，如果不存在公共前缀，返回空字符串""
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 去第一个字符串作为比较对象
	length := len(strs[0])
	for i := 0; i < length; i++ {
		// 比较的字符
		ch := strs[0][i]
		// 因为第一个已经拿出去的，这里要从第二个开始
		for j := 1; j < len(strs); j++ {
			//
			if len(strs[j]) < i+1 || ch != strs[j][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
