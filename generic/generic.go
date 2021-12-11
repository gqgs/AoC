package generic

import (
	"constraints"
	"math/rand"
)

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
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

type MinHeap[T constraints.Ordered] struct{ Heap Stack[T] }

func (h MinHeap[T]) Len() int            { return len(h.Heap) }
func (h MinHeap[T]) Less(i, j int) bool  { return h.Heap[i] < h.Heap[j] }
func (h MinHeap[T]) Swap(i, j int)       { h.Heap[i], h.Heap[j] = h.Heap[j], h.Heap[i] }
func (h MinHeap[T]) Min() T              { return h.Heap[0] }
func (h *MinHeap[T]) Push(x interface{}) { h.Heap.Push(x.(T)) }
func (h *MinHeap[T]) Pop() interface{}   { return h.Heap.Pop() }

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
