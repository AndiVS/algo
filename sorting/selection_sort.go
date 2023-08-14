package sorting

func SelectionSort(sl []int) []int {
	n := len(sl)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if sl[j] < sl[minIndex] {
				minIndex = j
			}
		}
		sl[i], sl[minIndex] = sl[minIndex], sl[i]
	}

	return sl
}
