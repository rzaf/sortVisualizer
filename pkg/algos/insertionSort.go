package algos

import "sort"

func InsertionSort(s sort.Interface) {
	l := s.Len()

	for i := 0; i < l; i++ {
		for j := 0; j < i; j++ {
			if s.Less(i, j) {
				s.Swap(i, j)
			}
		}
	}
}
