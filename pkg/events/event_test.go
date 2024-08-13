package events

import (
	"testing"
	"time"
)

func TestCreateEventWithPayload(t *testing.T) {
	t.Fatal("{} ", time.Now().UnixNano(), time.Now().UnixNano(), time.Now().UnixNano())
	e := CreateEventWithPayload("foo", []int{1, 3, 5}, "TEST")

	if e.Name() != "foo" {
		t.Fatal("event name must be 'foo'")
	}

	if e.From().Unwrap().(string) != "TEST" {
		t.Fatal("it must be equal TEST")
	}

	prob := []int{1, 3, 5}

	p := e.Payload().Unwrap().([]int)
	for i := 0; i < len(p); i++ {
		if p[i] != prob[i] {
			t.Fatal("it must be equal")
		}
	}
}
