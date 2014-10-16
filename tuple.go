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
	expires int64
	Tuple
}

func New(expires int64, data ...interface{}) *tuple {
	return &tuple{
		data:    data,
		expires: time.Now().Unix() + expires,
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
	return t.expires <= time.Now().Unix()
}
