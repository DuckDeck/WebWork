package foundation

import (
	"fmt"
	"math"
)

func Fibonacci(index int) int {
	if index <= 0 {
		return index
	} else if index <= 1 {
		return index
	} else {
		one := 0
		two := 1
		res := 0
		for i := 2; i <= index; i++ {
			res = one + two
			one = two
			two = res
		}
		return res
	}
}

//一只青蛙一次可以跳上1级台阶，也可以跳上2级。求该青蛙跳上一个n级的台阶总共有多少种跳法。
//每次只能跳一级或者跳两级，那么调N个台阶就可以分解为
//先跳1级，然后剩余的是跳一个n - 1级的台阶，
//先跳2级，然后剩余的是跳一个n - 2级的台阶，
//所以就是一个F(n-1) + F(n-2)个跳法
func FrogJump(step int) int {
	if step <= 2 {
		return step
	}
	one := 1
	two := 2
	res := 0
	for i := 3; i <= step; i++ {
		res = one + two
		one = two
		two = res
	}
	return res
}

//一只青蛙一次可以跳上1级台阶，也可以跳上2级……它也可以跳上n级。求该青蛙跳上一个n级的台阶总共有多少种跳法。
//n = 1 时 有一种跳法， F(1) = 1
//n = 2 时 有两种跳法 F(2-1) 和 F(2-2)
//n = 3 时 有三种跳法 一阶，二阶，三阶 那么就是第一次跳出1阶后面剩下：f(3-1);第一次跳出2阶，剩下f(3-2)；第一次3阶，那么剩下f(3-3) 因此结论是f(3) = f(3-1)+f(3-2)+f(3-3)
//n = n时 有k种跳法 一阶，二阶，三阶.....n阶 结果就是f(n) = f(n-1) + f(n-2) + ... + f(n-n) = f(0) + f(1) + f(2) + ... + f(n-1)
//因为f(n-1) = f(0) + f(1) + f(2) + ... + f(n-2) 所以 f(n) = f(n-1) + f(n-1) = 2 * f(n-1) = 4 * f(n-2) = 2 ^ 2* f(n-2) = 2^(n-1)

func NFrogJump(step int) int {
	if step <= 0 {
		return -1
	} else if step == 1 {
		return 1
	} else {
		return 2 * NFrogJump(step-1)
	}
}

func NFrogJump2(step int) int {
	if step <= 0 {
		return -1
	} else if step == 1 {
		return 1
	} else {
		return int(math.Pow(2, float64(step-1)))
	}
}

func TestFrog() {
	res := NFrogJump2(5)
	fmt.Printf("%d\n", res)
	res2 := NFrogJump2(5)
	fmt.Printf("%d", res2)
}

//我们可以用2*1的小矩形横着或者竖着去覆盖更大的矩形。请问用n个2*1的小矩形无重叠地覆盖一个2*n的大矩形，总共有多少种方法？
