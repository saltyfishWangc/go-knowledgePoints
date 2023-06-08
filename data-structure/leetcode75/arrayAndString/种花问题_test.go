package arrayAndString

import (
	"testing"
)

// 假设有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花不能种植在相邻的地块上，它们会争夺水源，两者都会死去。
// 给你一个整数数组flowerbed 表示花坛，由若干 0 和 1 组成，其中 0 表示没种植花，1 表示种植了花。另有一个数n ，能否在不打破种植规则的情况下种入n朵花？能则返回 true ，不能则返回 false。
// 示例 1：
//
// 输入：flowerbed = [1,0,0,0,1], n = 1
// 输出：true
// 示例 2：
//
// 输入：flowerbed = [1,0,0,0,1], n = 2
// 输出：false
func TestCanPlaceFlowers(t *testing.T) {
	flowered := []int{1, 0, 0, 0, 1}
	n := 1

	//flowered := []int{1, 0, 0, 0, 1}
	//n := 2
	t.Log(canPlaceFlowers(flowered, n))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	// 遍历数组，如果花已经种完了就没必要继续判断了
	for i := 0; i < len(flowerbed) && n > 0; {
		if flowerbed[i] == 1 {
			// 要跳一格
			i += 2
		} else {
			// 因为如果前一格等于1的话，i就会加2格，所以当flowerbed[i]=0的话，说明前面一格一定是0
			// 判断下一格是否等于1，如果不等于1的话，那这一格就可以种花
			if i+1 < len(flowerbed) && flowerbed[i+1] == 0 {
				i += 2
				n--
			} else {
				// 因为下一格是1，说明当前格不能种花，所以要跳3格
				i += 3
			}
		}
	}
	return n == 0
}
