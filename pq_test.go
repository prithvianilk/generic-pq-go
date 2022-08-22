package main

import (
	"testing"

	"golang.org/x/exp/constraints"
)

func TestInt(t *testing.T) {
	in := []int{1, 3, 2, 100, 50, 80}
	x := []int{100, 80, 50, 3, 2, 1}
	compare(in, x, t)
}

func TestFloat(t *testing.T) {
	in := []float32{12.23, 4343.123, .213, 213e3}
	x := []float32{213e3, 4343.123, 12.23, .213}
	compare(in, x, t)
}

func TestString(t *testing.T) {
	in := []string{"123123", "basd", "sada", "aa"}
	x := []string{"sada", "basd", "aa", "123123"}
	compare(in, x, t)
}

func compare[T constraints.Ordered](in []T, x []T, t *testing.T) {
	maxSize := len(in)
	pq := NewPriorityQueue[T](maxSize)
	for _, v := range in {
		pq.Push(v)
	}
	i := 0
	for !pq.IsEmpty() {
		top, ok := pq.Top()
		if !ok || top != x[i] {
			t.Fail()
		}
		i++
		pq.Pop()
	}
}
