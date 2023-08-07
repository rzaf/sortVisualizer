package algos

import "sort"

func SelectionSort(s sort.Interface) {
	for i := 0; i < s.Len(); i++ {
		minIndex := i
		for j := i + 1; j < s.Len(); j++ {
			if s.Less(j, minIndex) {
				minIndex = j
			}
		}
		s.Swap(minIndex, i)
	}
}
