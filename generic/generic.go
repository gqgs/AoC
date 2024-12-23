package generic

import (
	"cmp"
	heaplib "container/heap"
	"math/rand"
)

type Set[T comparable] map[T]struct{}

func (s Set[T]) Contains(e T) bool {
	_, exists := s[e]
	return exists
}

func (s *Set[T]) Add(e T) {
	(*s)[e] = struct{}{}
}

func (s *Set[T]) Len() int {
	return len(*s)
}

func (s Set[T]) Intersect(s2 Set[T]) Set[T] {
	intersection := make(Set[T])
	for ss := range s {
		if s2.Contains(ss) {
			intersection.Add(ss)
		}
	}
	return intersection
}

func (s Set[T]) Union(s2 Set[T]) Set[T] {
	union := make(Set[T])
	for ss := range s {
		union.Add(ss)
	}
	for ss := range s2 {
		union.Add(ss)
	}

	return union
}

func NewSet[T comparable](l ...T) Set[T] {
	s := make(Set[T])
	for _, e := range l {
		s[e] = struct{}{}
	}
	return s
}

func Abs[T int](a T) int {
	if a < 0 {
		return -int(a)
	}
	return int(a)
}

type Queue[T any] []T

func (s *Queue[T]) Push(r T) {
	*s = append(*s, r)
}

func (s *Queue[T]) Pop() T {
	r := (*s)[0]
	(*s) = (*s)[1:]
	return r
}

func (s Queue[T]) Empty() bool {
	return len(s) == 0
}

type Stack[T any] []T

func (s *Stack[T]) Push(r T) {
	*s = append(*s, r)
}

func (s *Stack[T]) Pop() T {
	n := len(*s)
	r := (*s)[n-1]
	*s = (*s)[0 : n-1]
	return r
}

func (s *Stack[T]) Peak() T {
	next := s.Pop()
	s.Push(next)
	return next
}

func (s Stack[T]) Empty() bool {
	return len(s) == 0
}

func (s Stack[T]) Len() int {
	return len(s)
}

type heap[T any] struct {
	stack    Stack[T]
	lessFunc func(x, y T) bool
}

func (h heap[T]) Len() int           { return len(h.stack) }
func (h heap[T]) Less(i, j int) bool { return h.lessFunc(h.stack[i], h.stack[j]) }
func (h heap[T]) Swap(i, j int)      { h.stack[i], h.stack[j] = h.stack[j], h.stack[i] }
func (h heap[T]) Min() T             { return h.stack[0] }
func (h *heap[T]) Push(x any)        { h.stack.Push(x.(T)) }
func (h *heap[T]) Pop() any          { return h.stack.Pop() }

func NewMinHeap[T any](lessFunc func(e1, e2 T) bool) MinHeap[T] {
	return MinHeap[T]{
		heap[T]{
			make(Stack[T], 0),
			lessFunc,
		},
	}
}

type MinHeap[T any] struct{ heap heap[T] }

func (h MinHeap[T]) Len() int { return h.heap.Len() }
func (h MinHeap[T]) Min() T   { return h.heap.Min() }
func (h *MinHeap[T]) Pop() T  { return heaplib.Pop(&h.heap).(T) }
func (h *MinHeap[T]) Push(l ...T) {
	for _, e := range l {
		heaplib.Push(&h.heap, e)
	}
}

func QuickSelect[T cmp.Ordered](list []T, k int) T {
	return selectKth(list, 0, len(list)-1, k)
}

func selectKth[T cmp.Ordered](list []T, left, right, k int) T {
	for {
		if left == right {
			return list[left]
		}
		pivotIndex := partition(list, left, right, rand.Intn(right-left+1)+left)
		if k == pivotIndex {
			return list[k]
		}
		if k < pivotIndex {
			right = pivotIndex - 1
		} else {
			left = pivotIndex + 1
		}
	}
}

func partition[T cmp.Ordered](list []T, left, right, pivotIndex int) int {
	pivot := list[pivotIndex]
	list[pivotIndex], list[right] = list[right], list[pivotIndex]
	storeIndex := left
	for i := left; i < right; i++ {
		if list[i] < pivot {
			list[storeIndex], list[i] = list[i], list[storeIndex]
			storeIndex++
		}
	}
	list[storeIndex], list[right] = list[right], list[storeIndex]
	return storeIndex
}
