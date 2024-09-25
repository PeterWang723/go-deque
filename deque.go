package godeque

const minCapacity = 32

type deque[T any] struct {
	buf []T
	head int
	tail int
	count int
	minCap int
}

func Make[T any](capacity, minimum int) Deque[T] {
	
	minCap := minCapacity

	for minCap < minimum {
		minCap <<= 1
	}

	var buffer []T

	if capacity != 0 {
		bufSize := minCap
		for bufSize < capacity {
			bufSize <<= 1
		}
		buffer = make([]T, bufSize)
	}

	return &deque[T]{buf: buffer, minCap: minCap}
}

func (q *deque[T]) Cap() int {
	if q == nil {
		return 0
	}

	return len(q.buf)
}

func (q *deque[T]) Len() int {
	if q == nil {
		return 0
	}

	return q.count
}

func (q *deque[T]) PushBack(item T) {
	q.grow()

	q.buf[q.tail] = item
	q.tail = q.next(q.tail)
	q.count++
}

func (q *deque[T]) PushFront(elem T) {
	q.grow()
	q.head = q.prev(q.head)
	q.buf[q.head] = elem
	q.count++
}

func (q *deque[T]) PopFront() error{
	
}

func (q *deque[T]) grow() {
	if q.count != len(q.buf) {
		return 
	}

	if len(q.buf) == 0 {
		if q.minCap == 0 {
			q.minCap = minCapacity
		}
		q.buf = make([]T, q.minCap)
		return
	}

	q.resize()


}

func (q *deque[T]) next(i int) int{
	return (i + 1) & (len(q.buf) - 1) 
}

func (q *deque[T]) prev(i int) int {
	return (i - 1) & (len(q.buf) - 1)
}

func (q *deque[T]) resize() {
	newBuf := make([]T, q.count << 1)
	if q.tail > q.head {
		copy(newBuf, q.buf[q.head:q.tail])
	} else {
		n := copy(newBuf, q.buf[q.head:])
		copy(newBuf[n:], q.buf[:q.tail])
	}

	q.head = 0
	q.tail = q.count
	q.buf = newBuf
}



