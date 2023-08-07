package algos

import "sort"

func heapify(arr sort.Interface, n int, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr.Less(largest, left) {
		largest = left
	}

	if right < n && arr.Less(largest, right) {
		largest = right
	}

	if largest != i {
		arr.Swap(i, largest)
		heapify(arr, n, largest)
	}
}

func HeapSort(arr sort.Interface) {
	n := arr.Len()
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	for i := n - 1; i >= 0; i-- {
		arr.Swap(0, i)
		heapify(arr, i, 0)
	}
}
