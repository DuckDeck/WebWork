package foundation

import (
	"fmt"
)

//在一个二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
//请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

func TwoDimensionArrayFind(array [][]int, target int) bool {
	rows := len(array)
	column := len(array[0])
	//通常我用先用最简单的方式来找,这样就没有利用到题目给的特性，一个某方便的序的数组
	// for i := 0; i < rows; i++ {
	// 	for j := 0; j < column; j++ {
	// 		if array[i][j] == target {
	// 			return true
	// 		}
	// 	}
	// }

	//可以利用这个题的特性，从左下角的数据开始出发,比他小的数必定在它的上侧，就往上找.比他大的数必定在它的右侧，就往右找

	for i := rows - 1; i >= 0 && i < rows; {
		for j := 0; j < column && j >= 0; {
			if target == array[i][i] {
				return true
			} else if target < array[i][j] {
				i--
			} else {
				j++
			}

		}
	}

	return false
}

func TestTwoDimensionArrayFind() {
	var arr = [][]int{
		{1, 3, 4, 8, 11, 17},
		{5, 7, 10, 13, 14, 19},
		{6, 9, 10, 15, 18, 20},
		{9, 12, 16, 22, 26, 28},
		{12, 15, 27, 31, 35, 44}}

	for i := 0; i < 5; i++ {
		for j := 0; j < 6; j++ {
			if arr[i][j] < 10 {
				fmt.Printf("  %d ", arr[i][j])
			} else {
				fmt.Printf(" %d ", arr[i][j])
			}

		}
		fmt.Printf("\n")
	}

	TwoDimensionArrayFind(arr, 31)

}

//一个整型数组里除了两个数字之外，其他的数字都出现了两次。请写程序找出这两个只出现一次的数字
//比如这样的 2,4,5,5,4,6,7,1,19,2,6,7
//我的思路是用一个hash表，出现同一个数就删除，循环完后里面的两个数就是只出现一次的数
//网上的思路是如果两个数相同，那么两个数字异或就为0,所以可以一直向后面异或过去，最后出现的值就是两个数的异或值
func findOneRepeatNum(array []int) []int {
	if len(array) <= 0 {
		return nil
	}

	tmp := array[0]
	for _, item := range array {
		if item == tmp {
			continue
		}
		tmp = tmp ^ item
	}
	count := 0
	for tmp%2 == 0 {
		tmp = tmp >> 1
		count += 1
	}
	mask := 1 << count
	var firstNum = -1
	var secondNum = -1
	for _, item := range array {
		if mask&item == 0 {
			if firstNum == -1 {
				firstNum = item
			} else {
				firstNum = firstNum ^ item
			}

		} else {
			if secondNum == -1 {
				secondNum = item
			} else {
				secondNum = secondNum ^ item
			}
		}
	}
	return []int{firstNum, secondNum}
}

//第二种方式用hash表
func findOneRepeatNum1(array []int) []int {
	if len(array) <= 0 {
		return nil
	}
	res := []int{}
	tmp := make(map[int]bool)
	for _, item := range array {
		if _, ok := tmp[item]; ok {
			delete(tmp, item)
		} else {
			tmp[item] = true
		}
	}
	for item := range tmp {
		res = append(res, item)
	}
	return res
}

//两个方法都行
func TestFindRepeatNum() {
	nums := []int{2, 4, 5, 5, 4, 6, 7, 19, 3, 2, 6, 7}
	res := findOneRepeatNum(nums)
	for _, item := range res {
		fmt.Printf("%d\n", item)
	}
}
