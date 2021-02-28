package foundation

//所有排序默认从小到大
func InsertSort() {

}

func BubbleSort(nums []int) {
	length := len(nums)
	for i := length - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if nums[j] > nums[i] {
				Swap(nums, i, j)
			}
		}
	}
}
