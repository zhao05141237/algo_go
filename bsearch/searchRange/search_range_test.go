package searchRange

import "testing"

func TestSearchRange(t *testing.T) {
	nums := []int{1, 2, 3, 3, 3, 3, 4, 5, 9}
	t.Log(SearchRange(nums, 3))
}
