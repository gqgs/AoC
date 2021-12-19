package generic

import (
	"constraints"
	heaplib "container/heap"
	"math/rand"
)

func Min[T constraints.Ordered](l ...T) (min T) {
	min = l[0]
	for _, e := range l[1:] {
		if e < min {
			min = e
		}
	}
	return
}

func Max[T constraints.Ordered](l ...T) (max T) {
	max = l[0]
	for _, e := range l[1:] {
		if e > max {
			max = e
		}
	}
	return
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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

type heap[T any] struct {
	stack    Stack[T]
	lessFunc func(x, y T) bool
}

func (h heap[T]) Len() int            { return len(h.stack) }
func (h heap[T]) Less(i, j int) bool  { return h.lessFunc(h.stack[i], h.stack[j]) }
func (h heap[T]) Swap(i, j int)       { h.stack[i], h.stack[j] = h.stack[j], h.stack[i] }
func (h heap[T]) Min() T              { return h.stack[0] }
func (h *heap[T]) Push(x interface{}) { h.stack.Push(x.(T)) }
func (h *heap[T]) Pop() interface{}   { return h.stack.Pop() }

func NewMinHeap[T any](lessFunc func(e1, e2 T) bool) minHeap[T] {
	return minHeap[T]{
		heap[T]{
			make(Stack[T], 0),
			lessFunc,
		},
	}
}

type minHeap[T any] struct{ heap heap[T] }

func (h minHeap[T]) Len() int { return h.heap.Len() }
func (h minHeap[T]) Min() T   { return h.heap.Min() }
func (h *minHeap[T]) Pop() T  { return heaplib.Pop(&h.heap).(T) }
func (h *minHeap[T]) Push(l ...T) {
	for _, e := range l {
		heaplib.Push(&h.heap, e)
	}
}

func QuickSelect[T constraints.Ordered](list []T, k int) T {
	return selectKth(list, 0, len(list)-1, k)
}

func selectKth[T constraints.Ordered](list []T, left, right, k int) T {
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

func partition[T constraints.Ordered](list []T, left, right, pivotIndex int) int {
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
