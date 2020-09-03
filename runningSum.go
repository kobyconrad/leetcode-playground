package main

func runningSum(nums []int) []int {
	answerArr := []int{}
	arrLength := len(nums)

	for i := 0; i < arrLength; i++ {

		if i == 0 {
			answerArr = append(answerArr, nums[i])
		} else {
			answerArr = append(answerArr, answerArr[i-1]+nums[i])
		}
	}

	return answerArr
}
