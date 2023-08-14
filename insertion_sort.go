package main

func insertionSort(sl []int) []int {
	n := len(sl)
	for i := 1; i < n; i++ {
		cur := sl[i]
		j := i - 1

		for j >= 0 && sl[j] > cur {
			sl[j+1] = sl[j]
			j--
		}
		sl[j+1] = cur
	}

	return sl
}
