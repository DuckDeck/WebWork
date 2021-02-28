package foundation

//所有排序默认从小到大
func InsertSort(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}
	for i := 1; i < length; i++ {
		value := nums[i]
		j := i - 1
		for ; j >= 0; j-- {
			if nums[j] > value {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		nums[j+1] = value
	}
}

func ChooseSort(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}
	min := 0
	for i := 0; i < length; i++ {
		min = i
		for j := i + 1; j < length; j++ {
			if nums[min] > nums[j] {
				min = j
			}
		}
		if min != i {
			Swap(nums, i, min)
		}

	}
}

func BubbleSort(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}
	for i := length - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if nums[j] > nums[i] {
				Swap(nums, i, j)
			}
		}
	}
}

func BubbleSort1(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}
	for i := length - 1; i >= 0; i-- {
		flag := false
		for j := i - 1; j >= 0; j-- {
			if nums[j] > nums[i] {
				Swap(nums, i, j)
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}
