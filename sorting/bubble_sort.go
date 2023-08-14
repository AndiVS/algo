package sorting

func BubbleSort(sl []int) []int {
	n := len(sl)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if sl[j] > sl[j+1] {
				sl[j], sl[j+1] = sl[j+1], sl[j]
			}
		}
	}
	return sl
}
