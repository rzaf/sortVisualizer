package algos

import "sort"

func BubbleSort(s sort.Interface) {
	l := s.Len()

	for i := 0; i < l; i++ {
		for j := 1; j < l-i; j++ {
			if s.Less(j, j-1) {
				s.Swap(j-1, j)
			}
		}
	}
}
