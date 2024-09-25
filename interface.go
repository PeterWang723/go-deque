package godeque

type Deque[T any] interface{
	Cap() int
	Len() int
	PushBack(T)
	PushFront(T)
	PopFront() (T, error)
}