package foundation

import (
	"fmt"
	"math/rand"
	"time"
)

func CreateNums(low int, high int, count int) (nums []int) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < count; i++ {
		n := rand.Intn(high - low)
		nums = append(nums, n+low)
	}
	return
}

func Swap(nums []int, i int, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}

func PrintNum(sep string, nums []int) {
	for _, i := range nums {
		fmt.Printf("%d%s", i, sep)
	}
	fmt.Printf("\n")
}
