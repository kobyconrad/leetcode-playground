package main

import "fmt"

func shuffle(nums []int, n int) []int {

	startArr := nums[:n]
	endArr := nums[n:]
	answerArr := []int{}

	// fmt.Println(startArr)
	// fmt.Println(endArr)

	for i := 0; i < n; i++ {
		answerArr = append(answerArr, startArr[i])
		answerArr = append(answerArr, endArr[i])
	}

	fmt.Println(answerArr)

	myIntReturn := []int{2}
	return myIntReturn
}

func main() {

	myArr := []int{2, 5, 1, 3, 4, 7}
	myInt := 3
	shuffle(myArr, myInt)

	// startArr := myArr[:2]
	// endArr := myArr[2:]

	// fmt.Println(startArr)
	// fmt.Println(endArr)

	// for _, item := range myArr {
	// 	fmt.Println(item)
	// }
}
