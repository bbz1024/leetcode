package codetop

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		num := nums[mid]
		if num == target {
			return mid
		}
		if nums[0] <= target {
			if nums[0] <= num {
				if nums[0] <= target && target < num {
					right = mid - 1
				} else {
					left = mid + 1
				}
			}
		} else {
			if num < target && target <= nums[len(nums)-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

	}
	return -1
}
