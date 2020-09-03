package main

func shuffle(nums []int, n int) []int {

	startArr := nums[:n]
	endArr := nums[n:]
	answerArr := []int{}

	for i := 0; i < n; i++ {
		answerArr = append(answerArr, startArr[i])
		answerArr = append(answerArr, endArr[i])
	}

	return answerArr
}

func main() {

	// myArr := []int{2, 5, 1, 3, 4, 7}
	// myInt := 3
	// shuffle(myArr, myInt)

	runningSum([]int{1, 2, 3})

}
