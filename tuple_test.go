package tuplespace

import (
	"reflect"
	"testing"
)

func TestTuple(t *testing.T) {
	tuple := New(600, `foo`, `bar`, `baz`)

	if tuple.Len() != 3 {
		t.Fatal(`invalid tuple length.`)
	}

	if !reflect.DeepEqual(tuple.Values(), []interface{}{`foo`, `bar`, `baz`}) {
		t.Fatal(`invalid tuple values.`)
	}

	if !tuple.Match(New(60, `foo`, `bar`)) {
		t.Fatal(`Tuple#Match is wrong implementation.`)
	}

	if tuple.Match(New(60, `baz`)) {
		t.Fatal(`Tuple#Match is wrong implementation.`)
	}

	expiredTuple := New(0, `foo`)

	if !expiredTuple.IsExpired() {
		t.Fatal(`Tuple#IsExpired is wrong implementation.`)
	}
}
