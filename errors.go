package godeque

import "fmt"
type Status int

const (
    ZeroCount Status = iota + 1
    NotFound
)

type DequeErr struct {
    Status    Status
    Message   string
}

func (de DequeErr) Error() string {
    return fmt.Sprintf("Status: %d, Message: %s", de.Status, de.Message)
}