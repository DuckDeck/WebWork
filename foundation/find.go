package foundation

func BiranySearch(nums []int, target int) (index int) {
	index = -1
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			index = mid
			return
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return
}

//找出第一个值等于等于给定值的元素
func FindFirsetTarget(nums []int, target int) (index int) {
	index = -1
	low := 0
	high := len(nums) - 1
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if mid == 0 || nums[mid-1] != target {
				return mid
			}
		}
	}
	return
}

//找出第一个出现大于该数的位置
func FirstBiggerNumToTarget(nums []int, target int) (index int) {
	index = -1
	low := 0
	high := len(nums) - 1
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] >= target { //如果大于该数
			if mid == 0 || (nums[mid-1] < target) { //如果已经是第0个或者前一个数比它小，就说明找到了
				return target
			} else {
				high = mid - 1 //不然后就向前面找
			}
		} else {
			low = mid + 1 //不然后就后前面找
		}

	}
	return
}

//找出最后一个出现小于该数的位置
func LastSmallerNumToTarget(nums []int, target int) (index int) {
	index = -1
	low := 0
	high := len(nums) - 1
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] > target {
			high = mid - 1
		} else {
			if mid == mid-1 || nums[mid+1] > target {
				return mid
			} else {
				low = mid + 1
			}
		}
	}
	return
}
