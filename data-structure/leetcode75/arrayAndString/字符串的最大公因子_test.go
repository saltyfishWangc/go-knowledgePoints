package arrayAndString

import (
	"strings"
	"testing"
)

func TestGcdOfStrings(t *testing.T) {
	str1 := "ABCABC"
	str2 := "ABC"
	t.Log(gcdOfStrings(str1, str2))
}

// gcdOfStrings 字符串的最大公因子
// 对于字符串s 和t，只有在s = t + ... + t（t 自身连接 1 次或多次）时，我们才认定“t 能除尽 s”。
//
// 给定两个字符串str1和str2。返回 最长字符串x，要求满足x 能除尽 str1 且 x 能除尽 str2 。
//
// 示例 1：
//
// 输入：str1 = "ABCABC", str2 = "ABC"
// 输出："ABC"
// 示例 2：
//
// 输入：str1 = "ABABAB", str2 = "ABAB"
// 输出："AB"
// 示例 3：
//
// 输入：str1 = "LEET", str2 = "CODE"
// 输出：""
func gcdOfStrings(str1, str2 string) string {

	// 得到两个字符串长度的最大公约数
	n := gcd(len(str1), len(str2))
	// 得到在字符串中的最大公约数对应的字符串
	subStr := str1[:n]
	// 判断最大公约数对应的字符串是否是两个字符串的最大公因子
	if check(str1, subStr) && check(str2, subStr) {
		return subStr
	}
	return ""
}

// check 校验s2是否能除尽s1，也就是说s1是由n个s2组成的
func check(s1, s2 string) bool {
	lenx := len(s1) / len(s2)
	var sb strings.Builder
	for i := 0; i < lenx; i++ {
		sb.WriteString(s2)
	}
	return s1 == sb.String()
}

// 辗转相除法得到两个数的最大公约数
func gcd(a, b int) int {
	remainder := a % b
	for remainder != 0 {
		a = b
		b = remainder
		remainder = a % b
	}
	return b
}
