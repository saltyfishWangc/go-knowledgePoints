package arrayAndString

import "testing"

func TestMergeAlternately(t *testing.T) {
	//word1 := "abc"
	//word2 := "pqr"

	word1 := "ab"
	word2 := "pqrs"
	t.Log(mergeAlternately(word1, word2))
}

// mergeAlternately 交替合并字符串
// 给你两个字符串 word1 和 word2 。请你从 word1 开始，通过交替添加字母来合并字符串。如果一个字符串比另一个字符串长，就将多出来的字母追加到合并后字符串的末尾。
//
// 返回 合并后的字符串

// 输入：word1 = "ab", word2 = "pqrs"
// 输出："apbqrs"
// 解释：注意，word2 比 word1 长，"rs" 需要追加到合并后字符串的末尾。
// word1：  a   b
// word2：    p   q   r   s
// 合并后：  a p b q   r   s

func mergeAlternately(word1 string, word2 string) string {
	if len(word1) == 0 {
		return word2
	}

	if len(word2) == 0 {
		return word1
	}

	var mergeNum []byte
	p1 := 0
	p2 := 0
	for p1 < len(word1) && p2 < len(word2) {
		mergeNum = append(mergeNum, word1[p1], word2[p2])
		p1++
		p2++
	}

	if p1 < len(word1) {
		mergeNum = append(mergeNum, word1[p1:]...)
	}

	if p2 < len(word2) {
		mergeNum = append(mergeNum, word2[p2:]...)
	}
	return string(mergeNum)

}
