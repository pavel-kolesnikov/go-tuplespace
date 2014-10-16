package tuplespace

import (
	"reflect"
	"testing"
)

func TestTupleSpaceReadAndTake(t *testing.T) {
	space := NewSpace()

	space.Write(New(600, `foo`, `bar`))

	recv1 := space.Read(New(0, `foo`))

	if t1 := <-recv1; !reflect.DeepEqual(t1.Values(), []interface{}{`foo`, `bar`}) {
		t.Fatal(`failed to Read from TupleSpace.`)
	}

	recv2 := space.Take(New(0, `foo`))

	if t2 := <-recv2; !reflect.DeepEqual(t2.Values(), []interface{}{`foo`, `bar`}) {
		t.Fatal(`failed to Take from TupleSpace.`)
	}

	if space.Len() > 0 {
		t.Fatal(`remove tuple from Take method is failed.`)
	}
}

func TestTupleSpaceWatchAndCancel(t *testing.T) {
	space := NewSpace()
	recv := make(chan Tuple)

	id := space.Watch(New(0, `foo`), recv)

	go func() {
		space.Write(New(600, `foo`, `bar`))
	}()

	if t1 := <-recv; !reflect.DeepEqual(t1.Values(), []interface{}{`foo`, `bar`}) {
		t.Fatal(`failed to Watch tuple.`)
	}

	ok := space.Cancel(id)
	if !ok {
		t.Fatal(`failed to cancel watcher.`)
	}
}
