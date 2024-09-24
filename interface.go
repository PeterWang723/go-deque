package godeque

type Deque[T any] interface{
	Cap() int
	Len() int
	PushBack(item T)
}