package utils

import "testing"

type maxHeapEl struct {
	MaxHeapElement

	k int
}

func (e *maxHeapEl) Key() int {
	return e.k
}

func TestMaxHeap(t *testing.T) {
	h := MaxHeap{}
	h.Push(&maxHeapEl{k: 3})
	h.Push(&maxHeapEl{k: 1})
	h.Push(&maxHeapEl{k: 7})

	if expVal := 7; h.Peek().Key() != expVal {
		t.Errorf("Wrong value, want: %v got: %v", expVal, h.Peek().Key())
	}

	h.Pop()

	if expVal := 3; h.Peek().Key() != expVal {
		t.Errorf("Wrong value, want: %v got: %v", expVal, h.Peek().Key())
	}

	h.Pop()

	if expVal := 1; h.Peek().Key() != expVal {
		t.Errorf("Wrong value, want: %v got: %v", expVal, h.Peek().Key())
	}

	h.Pop()
}
