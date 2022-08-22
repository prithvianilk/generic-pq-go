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
	index, par := pq.size, parent(pq.size)
	for index > 0 && pq.values[index] > pq.values[par] {
		pq.swapValues(index, par)
		index, par = par, parent(par)
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
	index := 0
	for !pq.isLeaf(index) {
		leftChild, rightChild := children(index)
		if pq.isRightValid(index) {
			pq.swapValues(index, rightChild)
			index = rightChild
		} else if pq.isLeftValid(index) {
			pq.swapValues(index, leftChild)
			index = leftChild
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

func (pq *PriorityQueue[T]) swapValues(i int, j int) {
	pq.values[j], pq.values[i] = pq.values[i], pq.values[j]
}

func children(index int) (int, int) {
	leftChild := (index * 2) + 1
	rightChild := (index * 2) + 2
	return leftChild, rightChild
}

func (pq *PriorityQueue[T]) isLeaf(index int) bool {
	leftChild := (2 * index) + 1
	return leftChild >= pq.size
}

func (pq *PriorityQueue[T]) isRightValid(index int) bool {
	leftChild, rightChild := children(index)
	isRightInRange := rightChild < pq.size
	if !isRightInRange {
		return false
	}
	isRightBiggerThanCurr := pq.values[rightChild] > pq.values[index]
	isRightBiggerThanLeft := pq.values[rightChild] > pq.values[leftChild]
	return isRightBiggerThanLeft && isRightBiggerThanCurr
}

func (pq *PriorityQueue[T]) isLeftValid(index int) bool {
	leftChild, _ := children(index)
	isLeftValid := pq.values[leftChild] > pq.values[index]
	return isLeftValid
}
