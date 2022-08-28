package main

import (
	"testing"
)

func insertValues[T any](pq *PriorityQueue[T], in []T) {
	for _, v := range in {
		pq.Push(v)
	}
}

func TestInt(t *testing.T) {
	in := []int{1, 3, 2, 100, 50, 80}
	ex := []int{100, 80, 50, 3, 2, 1}
	pq := NewPriorityQueue(func(lhs, rhs int) bool { return lhs > rhs }, len(in))
	insertValues(&pq, in)
	compare(pq, in, ex, t, func(lhs, rhs int) bool { return lhs == rhs })
}

func TestFloat(t *testing.T) {
	in := []float32{12.23, 4343.123, .213, 213e3}
	ex := []float32{213e3, 4343.123, 12.23, .213}
	pq := NewPriorityQueue(func(lhs, rhs float32) bool { return lhs > rhs }, len(in))
	insertValues(&pq, in)
	compare(pq, in, ex, t, func(lhs, rhs float32) bool { return lhs == rhs })
}

func TestString(t *testing.T) {
	in := []string{"123123", "basd", "sada", "aa"}
	ex := []string{"sada", "basd", "aa", "123123"}
	pq := NewPriorityQueue(func(lhs, rhs string) bool { return lhs > rhs }, len(in))
	insertValues(&pq, in)
	compare(pq, in, ex, t, func(lhs, rhs string) bool { return lhs == rhs })
}

func compare[T any](pq PriorityQueue[T], in []T, ex []T, t *testing.T, isEqual func(T, T) bool) {
	if pq.Size() != len(ex) {
		t.Fatal("Expected ", len(ex), "values. Got", pq.Size())
	}
	i := 0
	for !pq.IsEmpty() && i < len(ex) {
		top, ok := pq.Top()
		if !ok || !isEqual(top, ex[i]) {
			t.Fatal("Expected ", ex[i], ", got ", top)
		}
		pq.Pop()
		i++
	}
}
