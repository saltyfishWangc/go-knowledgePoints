package doublePointer

import "testing"

// 给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
//
// 字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）
// 示例 1：
//
// 输入：s = "abc", t = "ahbgdc"
// 输出：true
// 示例 2：
//
// 输入：s = "axc", t = "ahbgdc"
// 输出：false
func TestIsSubsequence(t *testing.T) {
	s, t1 := "abc", "ahbgdc"
	t.Log(isSubsequence(s, t1))
}

func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}
	if len(s) == 0 {
		return true
	}

	i, j := 0, 0
	for ; i < len(s); i++ {
		for ; j < len(t); j++ {
			if s[i] == t[j] {
				if i == len(s)-1 {
					// 说明s遍历完了
					return true
				} else {
					// 继续遍历s
					j++
					break
				}
			}
		}
	}
	return false
}
