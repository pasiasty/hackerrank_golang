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
	h := NewMaxHeap()

	toUpdateUp := &maxHeapEl{k: 0}
	toUpdateDown := &maxHeapEl{k: 10}

	h.Push(&maxHeapEl{k: 3})
	h.Push(&maxHeapEl{k: 2})
	h.Push(&maxHeapEl{k: 7})
	h.Push(toUpdateUp)
	h.Push(toUpdateDown)

	if expVal := 10; h.Peek().Key() != expVal {
		t.Errorf("Wrong value, want: %v got: %v", expVal, h.Peek().Key())
	}

	toUpdateDown.k = 1
	h.UpdatePosition(toUpdateDown)

	if expVal := 7; h.Peek().Key() != expVal {
		t.Errorf("Wrong value, want: %v got: %v", expVal, h.Peek().Key())
	}

	h.Pop()

	if expVal := 3; h.Peek().Key() != expVal {
		t.Errorf("Wrong value, want: %v got: %v", expVal, h.Peek().Key())
	}

	h.Pop()

	if expVal := 2; h.Peek().Key() != expVal {
		t.Errorf("Wrong value, want: %v got: %v", expVal, h.Peek().Key())
	}

	h.Pop()

	if expVal := 1; h.Peek().Key() != expVal {
		t.Errorf("Wrong value, want: %v got: %v", expVal, h.Peek().Key())
	}

	toUpdateUp.k = 12
	h.UpdatePosition(toUpdateUp)

	if expVal := 12; h.Peek().Key() != expVal {
		t.Errorf("Wrong value, want: %v got: %v", expVal, h.Peek().Key())
	}

	h.Pop()

	if expVal := 1; h.Peek().Key() != expVal {
		t.Errorf("Wrong value, want: %v got: %v", expVal, h.Peek().Key())
	}
}
