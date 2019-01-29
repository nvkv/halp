package output

import "github.com/nvkv/halp/pkg/types/data/v1"

type OutputError struct {
	str string
}

func NewOutputError(s string) *OutputError {
	return &OutputError{s}
}

func (e *OutputError) Error() string {
	return e.str
}

type Output interface {
	Send(schedule []data.Day) *OutputError
}
