package foundation

import "fmt"

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
