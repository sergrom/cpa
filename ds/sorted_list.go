package ds

import "sort"

type SortedList []int

func (list *SortedList) Insert(n int) {
	i := sort.SearchInts(*list, n)
	*list = append(*list, 0)
	copy((*list)[i+1:], (*list)[i:])
	(*list)[i] = n
}

func (list *SortedList) Remove(n int) {
	i := sort.SearchInts(*list, n)
	copy((*list)[i:], (*list)[i+1:])
	*list = (*list)[:len(*list)-1]
}
