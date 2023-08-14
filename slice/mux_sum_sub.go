package slice

func MaxSum(sl []int) int {
	maxV := 0
	curSum := 0
	for _, v := range sl {
		curSum += v
		if curSum < 0 {
			curSum = 0
		} else if curSum > maxV {
			maxV = curSum
		}

	}

	return maxV
}
