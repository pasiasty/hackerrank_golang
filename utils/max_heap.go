package utils

// MaxHeapElement is interface of elements stored in MaxHeap.
type MaxHeapElement interface {
	Key() int
}

// MaxHeap keeps max value always on top.
type MaxHeap struct {
	elements []MaxHeapElement
	count    int
}

func (h *MaxHeap) ensureLength() {
	for len(h.elements) < h.count {
		h.elements = append(h.elements, nil)
	}
}

func (h *MaxHeap) Peek() MaxHeapElement {
	return h.elements[0]
}

func (h *MaxHeap) Pop() {
	h.count--
	h.elements[0] = h.elements[h.count]
	h.pushDown(0)
}

func (h *MaxHeap) Push(e MaxHeapElement) {
	h.count++
	h.ensureLength()
	h.elements[h.count-1] = e
	h.pushUp(h.count - 1)
}

func (h *MaxHeap) pushDown(idx int) {
	if idx == (h.count - 1) {
		return
	}

	if (idx*2 + 1) < h.count {
		if h.elements[idx*2+1].Key() > h.elements[idx*2].Key() && h.elements[idx*2+1].Key() > h.elements[idx].Key() {
			temp := h.elements[idx*2+1]
			h.elements[idx*2+1] = h.elements[idx]
			h.elements[idx] = temp
			h.pushDown(idx*2 + 1)
			return
		}
	}

	if h.elements[idx*2].Key() > h.elements[idx].Key() {
		temp := h.elements[idx*2]
		h.elements[idx*2] = h.elements[idx]
		h.elements[idx] = temp
		h.pushDown(idx * 2)
	}
}

func (h *MaxHeap) pushUp(idx int) {
	if idx == 0 {
		return
	}

	if h.elements[idx].Key() > h.elements[idx>>1].Key() {
		temp := h.elements[idx>>1]
		h.elements[idx>>1] = h.elements[idx]
		h.elements[idx] = temp
		h.pushUp(idx >> 1)
	}
}
