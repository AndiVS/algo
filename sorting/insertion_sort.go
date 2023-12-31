package sorting

func InsertionSort(sl []int) []int {
	n := len(sl)
	for i := 1; i < n; i++ {
		cur := sl[i]
		j := i - 1
		for ; j >= 0 && sl[j] > cur; j-- {
			sl[j+1] = sl[j]
		}
		sl[j+1] = cur
	}

	return sl
}
