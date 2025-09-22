package ds

import "sort"

type SortedInts []int

func (list *SortedInts) Insert(n int) {
	i := sort.SearchInts(*list, n)
	*list = append(*list, 0)
	copy((*list)[i+1:], (*list)[i:])
	(*list)[i] = n
}

func (list *SortedInts) Remove(n int) {
	i := sort.SearchInts(*list, n)
	copy((*list)[i:], (*list)[i+1:])
	*list = (*list)[:len(*list)-1]
}

// ---------------

type item struct {
	val int
}

func lessFn(el1, el2 *item) bool {
	return el1.val < el2.val
}

// SortedArr ...
type SortedArr struct {
	arr  []*item
	less func(el1, el2 *item) bool
}

// NewSortedArr ...
func NewSortedArr(lessFn func(el1, el2 *item) bool) *SortedArr {
	return &SortedArr{
		arr:  make([]*item, 0),
		less: lessFn,
	}
}

// Insert ...
func (s *SortedArr) Insert(el *item) {
	i := sort.Search(len(s.arr), func(x int) bool { return s.less(el, s.arr[x]) })
	s.arr = append(s.arr, nil)
	copy(s.arr[i+1:], s.arr[i:])
	s.arr[i] = el
}

// Remove ...
func (s *SortedArr) Remove(el *item) {
	if i, ok := s.Find(el); ok {
		copy(s.arr[i:], s.arr[i+1:])
		s.arr = s.arr[:len(s.arr)-1]
	}
}

// Find ...
func (s *SortedArr) Find(el *item) (int, bool) {
	return sort.Find(len(s.arr), func(x int) int {
		if s.less(el, s.arr[x]) {
			return -1
		} else if s.less(s.arr[x], el) {
			return 1
		}
		return 0
	})
}

// Len ...
func (s *SortedArr) Len() int {
	return len(s.arr)
}

func (s *SortedArr) GetArr() []*item {
	return s.arr
}
