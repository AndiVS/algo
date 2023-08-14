package main

import "fmt"

func main() {
	fmt.Print(bubbleSort([]int{5, 3, 6, 1}))
}

func bubbleSort(sl []int) []int {
	for i := 0; i < len(sl); i++ {
		for j := 0; j < len(sl); j++ {
			if sl[i] < sl[j] {
				sl[i], sl[j] = sl[j], sl[i]
			}
		}

	}

	return sl
}
