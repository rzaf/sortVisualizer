package sortv

import (
	"sort"
)

var steps []*Step

/// Sortable

type Sortable []int

func (s Sortable) Len() int {
	return len(s)
}

func (s Sortable) Less(i, j int) bool {
	steps = append(steps, NewStep(StepLess, i, j))
	return s[i] < s[j]
}

func (s Sortable) Swap(i, j int) {
	steps = append(steps, NewStep(StepSwap, i, j))
	s[i], s[j] = s[j], s[i]
}

/// sortAlgo

type sortAlgo struct {
	algo func(sort.Interface)
	name string
}

func NewSortAlgo(algo func(sort.Interface), name string) *sortAlgo {
	return &sortAlgo{algo, name}
}

func (s *sortAlgo) Sort(s2 Sortable) {
	s.algo(s2)
}
