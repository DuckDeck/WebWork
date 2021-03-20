package foundation

import "fmt"

var result []int
var n int

func NQueen() {

	fmt.Print("Please input the Number N for the queen\n")
	if n == 0 {
		n = 2
	}
	result = make([]int, n)
	callQueen(0)
	printQueen()
}

func callQueen(row int) {
	if row == n {
		return
	}

	for i := 0; i < n; i++ {
		if isOK(row, i) {
			result[row] = i
			callQueen(row + 1)
		}
	}
}

func isOK(row int, column int) (ok bool) {
	var leftup = column - 1
	var rightup = column + 1
	for i := row - 1; i >= 0; i-- {
		if result[i] == column {
			return false
		}
		if leftup >= 0 {
			if result[i] == leftup {
				return false
			}
		}
		if rightup < n {
			if result[i] == rightup {
				return false
			}
		}
		leftup--
		rightup++
	}
	return true
}

func printQueen() {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if result[i] == j {
				fmt.Print("Q ")
			} else {
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
