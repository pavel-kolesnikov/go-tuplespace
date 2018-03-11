package tuplespace

import (
	"reflect"
	"time"
)

type Tuple interface {
	Len() int
	Values() []interface{}
	Match(tuple Tuple) bool
	IsExpired() bool
}

type tuple struct {
	data    []interface{}
	expires time.Time
	Tuple
}

func New(expires time.Duration, data ...interface{}) Tuple {
	return &tuple{
		data:    data,
		expires: time.Now().Add(expires),
	}
}

func (t *tuple) Len() int {
	return len(t.data)
}

func (t *tuple) Values() []interface{} {
	return t.data
}

func (t1 *tuple) Match(t2 Tuple) bool {
	if t1.Len() < t2.Len() {
		return false
	}

	for idx, t2val := range t2.Values() {
		t1val := t1.data[idx]

		if !reflect.DeepEqual(t1val, t2val) {
			return false
		}
	}

	return true
}

func (t *tuple) IsExpired() bool {
	return time.Until(t.expires) <= 0
}
