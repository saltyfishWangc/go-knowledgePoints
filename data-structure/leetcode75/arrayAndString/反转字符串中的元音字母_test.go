package arrayAndString

import "testing"

// 给你一个字符串 s ，仅反转字符串中的所有元音字母，并返回结果字符串。
//
// 元音字母包括 'a'、'e'、'i'、'o'、'u'，且可能以大小写两种形式出现不止一次。
// 示例 1：
//
// 输入：s = "hello"
// 输出："holle"
// 示例 2：
//
// 输入：s = "leetcode"
// 输出："leotcede"
func TestReverseVowels(t *testing.T) {
	s := "aA"
	t.Log(reverseVowels(s))
}

// reverseVowels 反转字符串中的元音字母
func reverseVowels(s string) string {
	// 因为要修改字符串，而字符串是只读的，所以不能直接对字符串进行修改
	arr := []byte(s)
	// 定义双指针，初始化分别是字符串的首尾。首指针大小不能大于尾指针大小
	for i, j := 0, len(arr)-1; i <= j; {
		// 如果两个指针指向的字母都是元音字母，则交换。同时指针移动
		if isYuanYin(arr[i]) && isYuanYin(arr[j]) {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		} else if isYuanYin(arr[i]) {
			// 首指针指向的字母是元音字母，但是尾指针指向的字母不是元音字母，此时无法交换，首指针移动
			j--
		} else {
			// 尾指针指向的字母是元音字母，但是首指针指向的字母不是元音字母，此时无法交换，尾指针移动
			i++
		}
	}
	return string(arr)
}

// isYuanYin 判断是否是元音字母
func isYuanYin(b byte) bool {
	if b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' || b == 'A' || b == 'E' || b == 'I' || b == 'O' || b == 'U' {
		return true
	}
	return false
}
