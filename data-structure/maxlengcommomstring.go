package data_structure

func longestCommonPrefix1(strs []string) string {
	// 先拿到长度最短的字符串
	// 循环遍历长度最短的字符串从最后一个字符开始比较，相等就往前取字符继续比较，不相等就返回""
	targetStr := strs[0]
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < len(targetStr) {
			targetStr = strs[i]
		}
	}

	// targetEndIndex := len(targetStr) - 1

	// finalEndIndex := targetEndIndex

	// for targetEndIndex >= 0 {
	//     for _, str := range strs {
	//         if targetStr[finalEndIndex] != str[finalEndIndex] {
	//             targetEndIndex--
	//             finalEndIndex--
	//             break
	//         }
	//     }
	//     targetEndIndex--
	// }
	// if finalEndIndex < 0 {
	//     return ""
	// } else {
	//     finalEndIndex += 1
	//     return string(targetStr[:finalEndIndex])
	// }

	inde := 0
	finalEndIndex := -1
	isEnd := false
	for ; inde < len(targetStr); inde++ {
		for _, str := range strs {
			if targetStr[inde] != str[inde] {
				isEnd = true
				break
			}
		}
		if !isEnd {
			finalEndIndex = inde
		}
	}
	if finalEndIndex < 0 {
		return ""
	} else {
		finalEndIndex += 1
		return string(targetStr[:finalEndIndex])
	}
}
