package fuse

import "unsafe"

// buffer provides a mechanism for constructing a message from
// multiple segments.
type buffer []byte

// alloc allocates size bytes and returns a pointer to the new
// segment.
func (w *buffer) alloc(size uintptr) unsafe.Pointer {
	s := int(size)
	if len(*w)+s > cap(*w) {
		old := *w
		*w = make([]byte, len(*w), 2*cap(*w)+s)
		copy(*w, old)
	}
	l := len(*w)
	*w = (*w)[:l+s]
	return unsafe.Pointer(&(*w)[l])
}

// reset clears out the contents of the buffer.
func (w *buffer) reset() {
	for i := range (*w)[:cap(*w)] {
		(*w)[i] = 0
	}
	*w = (*w)[:0]
}

func newBuffer(id RequestID, extra uintptr) (buffer, *outHeader) {
	buf := make(buffer, 0, unsafe.Sizeof(outHeader{})+extra)
	h := (*outHeader)(buf.alloc(unsafe.Sizeof(outHeader{})))
	h.Unique = uint64(id)
	return buf, h
}
