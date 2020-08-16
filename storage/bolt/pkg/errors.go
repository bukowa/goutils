package pkg

import "fmt"

type Error struct {
	model Model
	msg   string
}

func NewError(msg string) Error {
	return Error{msg: msg}
}

func (e Error) ForModel(m Model) *Error {
	ne := e
	ne.model = m
	return &ne
}

func (e Error) ForString(s string) *Error {
	ne := e
	ne.msg = fmt.Sprintf(ne.msg, s)
	return &ne
}

func (e *Error) Error() string {
	return e.String()
}

func (e *Error) String() string {
	return fmt.Sprintf(e.msg, e.model)
}
