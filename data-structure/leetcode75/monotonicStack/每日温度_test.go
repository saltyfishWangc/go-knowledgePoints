package monotonicStack

import "testing"

// 给定一个整数数组temperatures，表示每天的温度，返回一个数组answer，其中answer[i]是指对于第 i 天，下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用0 来代替。
//
//示例 1:
//
//输入: temperatures = [73,74,75,71,69,72,76,73]
//输出:[1,1,4,2,1,1,0,0]
//示例 2:
//
//输入: temperatures = [30,40,50,60]
//输出:[1,1,1,0]
//示例 3:
//
//输入: temperatures = [30,60,90]
//输出: [1,1,0]

func TestDailyTemperatures(t *testing.T) {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	t.Log(dailyTemperatures(temperatures))
}

func dailyTemperatures(temperatures []int) []int {
	answer := make([]int, len(temperatures))

	// 定义栈
	var stack []int

	// 遍历每日温度数组
	for i, v := range temperatures {
		// 栈不为空，且栈顶元素是大于当前温度
		for len(stack) > 0 && v > temperatures[stack[len(stack)-1]] {
			// 存在，就要依次出栈
			preIndex := stack[len(stack)-1]
			answer[preIndex] = i - preIndex
			// 栈顶出栈后剩余的栈内元素
			stack = stack[:len(stack)-1]
		}
		// 元素在数组中的下标入栈
		stack = append(stack, i)
	}
	return answer
}
