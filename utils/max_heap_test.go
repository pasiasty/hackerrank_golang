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
		t.Errorf("Peek() Wrong value, want: %v got: %v", expVal, h.Peek().Key())
	}
	if expVal := 5; h.Len() != expVal {
		t.Errorf("Len() Wrong value, want: %v got: %v", expVal, h.Len())
	}

	toUpdateDown.k = 1
	h.UpdatePosition(toUpdateDown)

	if expVal := 7; h.Peek().Key() != expVal {
		t.Errorf("Peek() Wrong value, want: %v got: %v", expVal, h.Peek().Key())
	}
	if expVal := 5; h.Len() != expVal {
		t.Errorf("Len() Wrong value, want: %v got: %v", expVal, h.Len())
	}

	h.Pop()

	if expVal := 3; h.Peek().Key() != expVal {
		t.Errorf("Peek() Wrong value, want: %v got: %v", expVal, h.Peek().Key())
	}
	if expVal := 4; h.Len() != expVal {
		t.Errorf("Len() Wrong value, want: %v got: %v", expVal, h.Len())
	}

	h.Pop()

	if expVal := 2; h.Peek().Key() != expVal {
		t.Errorf("Peek() Wrong value, want: %v got: %v", expVal, h.Peek().Key())
	}
	if expVal := 3; h.Len() != expVal {
		t.Errorf("Len() Wrong value, want: %v got: %v", expVal, h.Len())
	}

	h.Pop()

	if expVal := 1; h.Peek().Key() != expVal {
		t.Errorf("Peek() Wrong value, want: %v got: %v", expVal, h.Peek().Key())
	}
	if expVal := 2; h.Len() != expVal {
		t.Errorf("Len() Wrong value, want: %v got: %v", expVal, h.Len())
	}

	toUpdateUp.k = 12
	h.UpdatePosition(toUpdateUp)

	if expVal := 12; h.Peek().Key() != expVal {
		t.Errorf("Peek() Wrong value, want: %v got: %v", expVal, h.Peek().Key())
	}
	if expVal := 2; h.Len() != expVal {
		t.Errorf("Len() Wrong value, want: %v got: %v", expVal, h.Len())
	}

	h.Pop()

	if expVal := 1; h.Peek().Key() != expVal {
		t.Errorf("Peek() Wrong value, want: %v got: %v", expVal, h.Peek().Key())
	}
	if expVal := 1; h.Len() != expVal {
		t.Errorf("Len() Wrong value, want: %v got: %v", expVal, h.Len())
	}

	h.Pop()

	if expVal := 0; h.Len() != expVal {
		t.Errorf("Len() Wrong value, want: %v got: %v", expVal, h.Len())
	}
}
