package main

func bubbleSort(sl []int) []int {
	//for i := 0; i < len(sl); i++ {
	//	for j := 0; j < len(sl); j++ {
	//		if sl[i] < sl[j] {
	//			sl[i], sl[j] = sl[j], sl[i]
	//		}
	//	}
	//
	//}

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
