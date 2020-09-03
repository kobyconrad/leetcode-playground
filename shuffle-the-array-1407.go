package main

import "fmt"

func shuffle(nums []int, n int) []int {
	fmt.Println(nums)
	fmt.Println(n)

	myIntReturn := []int{2}

	return myIntReturn
}

func main() {

	myArr := []int{1, 2, 3}
	// myInt := 3
	// shuffle(myArr, myInt)

	startArr := myArr[:2]
	endArr := myArr[2:]

	fmt.Println(startArr)
	fmt.Println(endArr)

	// for _, item := range myArr {
	// 	fmt.Println(item)
	// }
}
