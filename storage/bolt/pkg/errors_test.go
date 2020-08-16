package pkg

import (
	"reflect"
	"testing"
)

type TestModel struct {
	key string
}

func (m *TestModel) Key() []byte {
	return []byte(m.key)
}

func TestError_ForModel(t *testing.T) {
	var m = TestModel{""}
	var e = NewError("%s")
	var e2 = e.ForModel(&m)

	e.model = &m

	if reflect.DeepEqual(&e, &e2) {
		t.Error()
	}
	if reflect.DeepEqual(e, e2) {
		t.Error()
	}
	if e.msg != e2.msg {
		t.Error()
	}
}

func TestError_String(t *testing.T) {
	var e = NewError("%v")
	e = *e.ForModel(&TestModel{key: ""})
	if e.String() != "&{}" {
		t.Error(e.String())
	}
}

func TestError_Error(t *testing.T) {
	var e = NewError("%s")
	e = *e.ForModel(&TestModel{key: "a"})
	if e.String() != "&{a}" {
		t.Error(e.String())
	}
}
