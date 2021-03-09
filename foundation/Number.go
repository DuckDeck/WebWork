package foundation

import (
	"fmt"
	"math"
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

//把只包含因子2、3和5的数称作丑数
//例如6、8都是丑数，但14不是，因为它包含因子7。 习惯上我们把1当做是第一个丑数。求按从小到大的顺序的第N个丑数。
func FindUglyNum(index int) int {
	if index <= 0 {
		return 0
	}
	uglyList := []int{1}
	two := 0
	three := 0
	five := 0
	count := 1
	for count <= index {
		minValue := min(2*uglyList[two], 3*uglyList[three], 5*uglyList[five])
		uglyList = append(uglyList, minValue)
		count++
		if minValue == 2*uglyList[two] {
			two++
		}
		if minValue == 3*uglyList[three] {
			three++
		}
		if minValue == 5*uglyList[five] {
			five++
		}
	}
	return uglyList[count-1]
}

func min(a int, b int, c int) int {
	return int(math.Min(math.Min(float64(a), float64(b)), float64(c)))
}
