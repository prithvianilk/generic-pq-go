package main

import (
	"testing"

	"golang.org/x/exp/constraints"
)

func TestInt(t *testing.T) {
	in := []int{1, 3, 2, 100, 50, 80}
	ex := []int{100, 80, 50, 3, 2, 1}
	compare(in, ex, t)
}

func TestFloat(t *testing.T) {
	in := []float32{12.23, 4343.123, .213, 213e3}
	ex := []float32{213e3, 4343.123, 12.23, .213}
	compare(in, ex, t)
}

func TestString(t *testing.T) {
	in := []string{"123123", "basd", "sada", "aa"}
	ex := []string{"sada", "basd", "aa", "123123"}
	compare(in, ex, t)
}

func compare[T constraints.Ordered](in []T, ex []T, t *testing.T) {
	maxSize := len(in)
	pq := NewPriorityQueue[T](maxSize)
	for _, v := range in {
		pq.Push(v)
	}
	i := 0
	for !pq.IsEmpty() && i < len(ex) {
		top, ok := pq.Top()
		if !ok || top != ex[i] {
			t.Fatal("Expected ", ex[i], ", got ", top)
		}
		i++
		pq.Pop()
	}
	if !pq.IsEmpty() {
		t.Fatal("Expected ", len(ex), "values. Got", i+pq.Size())
	}
}
