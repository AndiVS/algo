package slice

func MaxSum(sl []int) int {
	temp := []int{}
	curSum := 0
	for _, v := range sl {
		curSum += v
		if curSum < 0 {
			curSum = 0
		} else {
			temp = append(temp, curSum)
		}

	}

	maxV := 0
	for _, v := range temp {
		if v > maxV {
			maxV = v
		}
	}

	return maxV
}
