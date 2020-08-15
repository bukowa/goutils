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
	FlagSet *flag.FlagSet
	Name    string
	Default interface{}
	Help    string
}

func NewOpts(name string, dfl interface{}, help string, sets ...*flag.FlagSet) *Opts {
	var set *flag.FlagSet

	if len(sets) == 0 {
		set = flag.CommandLine
	} else if len(sets) == 1 {
		set = sets[0]
	} else {
		panic("Only one FlagSet can be passed!")
	}
	return &Opts{
		FlagSet: set,
		Name:    name,
		Default: dfl,
		Help:    help,
	}
}

type String struct {
	*Opts
	Value string
}

type Bool struct {
	*Opts
	Value bool
}

func (s *String) Set() {
	switch t := s.Opts.Default.(type) {
	case string:
		s.FlagSet.StringVar(&s.Value, s.Name, t, s.Help)
	default:
		panic(ErrDefaultValueType("string", t))
	}
}

func (b *Bool) Set() {
	switch t := b.Opts.Default.(type) {
	case bool:
		b.FlagSet.BoolVar(&b.Value, b.Name, t, b.Help)
	default:
		panic(ErrDefaultValueType("bool", t))
	}
}
