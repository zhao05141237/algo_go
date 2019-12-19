package searchRange

import "testing"

func TestSearchRange(t *testing.T) {
	nums := []int{7, 3, 4, 8, 4, 5}
	t.Log(SearchRange(nums, 8))
}
