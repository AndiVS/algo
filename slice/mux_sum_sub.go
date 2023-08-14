package slice

func maxSubArray(nums []int) int {
	maxV := nums[0]
	curSum := 0
	for _, v := range nums {
		if curSum < 0 {
			curSum = 0
		}
		curSum += v
		if curSum > maxV {
			maxV = curSum
		}

	}

	return maxV
}
