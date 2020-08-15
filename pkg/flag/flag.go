package flag

import (
	"errors"
	"flag"
	"fmt"
)

func ErrDefaultValueType(s string, t interface{}) error {
	return errors.New(fmt.Sprintf("invalid type %v", t))
}

type Opts struct {
	Name    string
	Default interface{}
	Help    string
}

type String struct {
	Opts
	Value string
}

type Bool struct {
	Opts
	Value bool
}

func (s *String) Set() {
	switch t := s.Opts.Default.(type) {
	case string:
		flag.StringVar(&s.Value, s.Name, t, s.Help)
	default:
		panic(ErrDefaultValueType("string", t))
	}
}

func (b *Bool) Set() {
	switch t := b.Opts.Default.(type) {
	case bool:
		flag.BoolVar(&b.Value, b.Name, t, b.Help)
	default:
		panic(ErrDefaultValueType("bool", t))
	}
}
