package main

import (
	"golang.org/x/exp/constraints"
)

type PriorityQueue[T constraints.Ordered] struct {
	values  []T
	size    int
	maxSize int
}

func NewPriorityQueue[T constraints.Ordered](maxSize int) PriorityQueue[T] {
	return PriorityQueue[T]{values: make([]T, maxSize), size: 0, maxSize: maxSize}
}

func (pq *PriorityQueue[T]) Push(value T) bool {
	if pq.size == pq.maxSize {
		return false
	}
	pq.values[pq.size] = value
	i, j := pq.size, parent(pq.size)
	for i > 0 && pq.values[i] > pq.values[j] {
		pq.values[i], pq.values[j] = pq.values[j], pq.values[i]
		i, j = j, parent(j)
	}
	pq.size++
	return true
}

func (pq *PriorityQueue[T]) Pop() bool {
	if pq.IsEmpty() {
		return false
	}
	pq.size--
	pq.values[0] = pq.values[pq.size]
	i := 0
	for !pq.isLeaf(i) {
		l, r := children(i)
		if pq.isRightValid(i) {
			pq.values[r], pq.values[i] = pq.values[i], pq.values[r]
			i = r
		} else if pq.isLeftValid(l, i) {
			pq.values[l], pq.values[i] = pq.values[i], pq.values[l]
			i = l
		} else {
			break
		}
	}
	return true
}

func (pq *PriorityQueue[T]) Top() (T, bool) {
	if pq.IsEmpty() {
		var top T
		return top, false
	}
	return pq.values[0], true
}

func (pq *PriorityQueue[T]) Size() int {
	return pq.size
}

func (pq *PriorityQueue[T]) IsEmpty() bool {
	return pq.Size() == 0
}

func parent(index int) int {
	return (index - 1) / 2
}

func children(index int) (int, int) {
	return ((index * 2) + 1), ((index * 2) + 2)
}

func (pq *PriorityQueue[T]) isLeaf(index int) bool {
	return ((2 * index) + 1) >= pq.size
}

func (pq *PriorityQueue[T]) isLeftValid(l int, i int) bool {
	isLeftValid := pq.values[l] > pq.values[i]
	return isLeftValid
}

func (pq *PriorityQueue[T]) isRightValid(i int) bool {
	l, r := children(i)
	isRightIndexable := r < pq.size
	if !isRightIndexable {
		return false
	}
	isRightBiggerThanCurr := pq.values[r] > pq.values[i]
	isRightBiggerThanLeft := pq.values[r] > pq.values[l]
	return isRightBiggerThanLeft && isRightBiggerThanCurr
}
