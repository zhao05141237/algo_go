package searchRange

/**
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
你的算法时间复杂度必须是 O(log n) 级别。
如果数组中不存在目标值，返回 [-1, -1]。
示例 1:
输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
示例 2:
输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]
*/

func SearchRange(nums []int, target int) []int {
	c := []int{-1, -1}
	if len(nums) < 1 {
		return c
	} else if len(nums) == 1 {
		if nums[0] == target {
			c[0] = 0
			c[1] = 0
			return c
		} else {
			return c
		}
	} else if len(nums) == 2 {
		if nums[0] == target && nums[1] == target {
			c[0] = 0
			c[1] = 1
			return c
		} else if nums[0] == target && nums[1] != target {
			c[0] = 0
			c[1] = 0
			return c
		} else if nums[0] != target && nums[1] == target {
			c[1] = 1
			c[0] = 1
			return c
		} else {
			return c
		}
	}

	low := 0
	high := len(nums) - 1
	mid := low + (high-low)>>1

	for low <= high {
		if nums[low] == target && c[0] == -1 {
			c[0] = low
		} else if nums[low] == target && low < c[0] {
			c[0] = low
		}
		if nums[high] == target && c[1] == -1 {
			c[1] = high
		} else if nums[high] == target && high > c[1] {
			c[1] = high
		}

		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			if c[0] == -1 {
				c[0] = mid
			} else if c[0] > mid {
				c[0] = mid
			} else if c[1] == -1 {
				c[1] = mid
			} else if c[1] < mid {
				c[1] = mid
			}
			low++
		}
		mid = low + (high-low)>>1
	}

	if c[0] == -1 && c[1] != -1 {
		c[0] = c[1]
	} else if c[0] > c[1] && c[1] != -1 {
		c[0], c[1] = c[1], c[0]
	} else if c[1] == -1 && c[0] != -1 {
		c[1] = c[0]
	}

	return c

}
