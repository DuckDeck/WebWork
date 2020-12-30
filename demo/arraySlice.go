package demo

import "fmt"

func Slice1() {
	souce := [...]int{1, 2, 3, 4, 5, 6}
	sli := souce[0:3]
	fmt.Printf("sli value is %v\n", sli)
	fmt.Printf("sli len is %v\n", len(sli))
	fmt.Printf("sli cap is %v\n", cap(sli))

}
