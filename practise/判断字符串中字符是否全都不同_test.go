package practise

import (
	"strings"
	"testing"
)

// TestIsCharAllDifferentInString 判断字符串中字符是否全都不同
// 题目：
// 请实现⼀个算法，确定⼀个字符串的所有字符【是否全都不同】。这⾥我们要求【不允
// 许使⽤额外的存储结构】。 给定⼀个string，请返回⼀个bool值,true代表所有字符全都
// 不同，false代表存在相同的字符。 保证字符串中的字符为【ASCII字符】。字符串的⻓
// 度⼩于等于【3000】
func TestIsCharAllDifferentInString(t *testing.T) {

}

// isCharAllDifferentInString
// 遍历字符串，用strings.Count判断每个字符在整个字符串s中的个数，如果超过1，说明该字符重复
func isCharAllDifferentInString(s string) bool {
	// 判断字符串的长度是否小于3000。strings.Count会根据子串的长度判断，如果子串长度为0，则返回字符串的字符长度
	if strings.Count(s, "") > 3000 {
		return false
	}
	for _, v := range s {
		// ASCII字符一共有256个，其中128个是常用字符，可以在键盘上输入，128之后的字符是键盘上无法找到的
		if v > 127 {
			return false
		}
		if strings.Count(s, string(v)) > 1 {
			return false
		}
	}

	// 还有一种办法是通过strings.LastIndex(s, subS)
	//for k, v := range s {
	//	// ASCII字符一共有256个，其中128个是常用字符，可以在键盘上输入，128之后的字符是键盘上无法找到的
	//	if v > 127 {
	//		return false
	//	}
	//	// strings.LastIndex返回字符在字符串中最后的位置
	//	if strings.LastIndex(s, string(v)) != k {
	//		return false
	//	}
	//}
	return true
}
