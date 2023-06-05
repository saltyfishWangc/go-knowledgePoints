package array

import "testing"

func TestContainsDuplicate(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	t.Log(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {
	mp := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := mp[num]; ok {
			return true
		} else {
			mp[num] = struct{}{}
		}
	}
	return false
}
