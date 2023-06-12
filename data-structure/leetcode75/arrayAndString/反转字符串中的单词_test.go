package arrayAndString

import (
	"strings"
	"testing"
)

func TestReverseWords(t *testing.T) {
	s := "the sky is blue"
	t.Log(reverseWords(s))
}

// 遍历字符串
func reverseWords1(s string) string {
	var strArr []string

	var sb strings.Builder
	for _, v := range []byte(s) {
		if v == ' ' {
			// 遇到空格，判断sb是否为空，为空不做任何处理，不为空则表示遍历完一个单词
			if sb.Len() != 0 {
				strArr = append(strArr, sb.String())
				sb.Reset()
			}
		} else {
			sb.WriteByte(v)
		}
	}
	// 考虑到最后一个字符不是空格
	if sb.Len() != 0 {
		strArr = append(strArr, sb.String())
		sb.Reset()
	}

	// 从切片strArr尾部开始遍历
	for i := len(strArr) - 1; i >= 0; i-- {
		sb.WriteString(strArr[i])
		if i != 0 {
			sb.WriteString(" ")
		}
	}
	return sb.String()
}

// 不引入额外数组空间
func reverseWords(s string) string {

	var sb strings.Builder

	var res strings.Builder

	for _, v := range []byte(s) {
		if v == ' ' {
			// 遇到空格，判断sb是否为空，为空不做任何处理，不为空则表示遍历完一个单词
			if sb.Len() != 0 {
				res.WriteString(sb.String())
				res.WriteString(" ")
				sb.Reset()
			}
		} else {
			sb.WriteByte(v)
		}
	}
	// 考虑到最后一个字符不是空格
	if sb.Len() != 0 {
		res.WriteString(sb.String())
		sb.Reset()
	}
	return res.String()
}
