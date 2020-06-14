package deep

import (
	"reflect"
	"testing"
)

type T struct {
	x  int
	ss []string
	m  map[string]int
}

func TestStruct(t *testing.T) {
	m1 := map[string]int{
		"a": 1,
		"b": 2,
	}
	t1 := T{
		x:  1,
		ss: []string{"a", "b"},
		m:  m1,
	}
	t2 := T{
		x:  1,
		ss: []string{"a", "b"},
		m:  m1,
	}
	t3 := T{
		x:  1,
		ss: []string{"a", "b", "c"},
		m:  m1,
	}

	if !reflect.DeepEqual(t1, t2) {
		t.Errorf("want %#v got %#v", t1, t2)
	}

	if !reflect.DeepEqual(t1, t3) {
		t.Errorf("want %#v got %#v", t1, t3)
	}
}
