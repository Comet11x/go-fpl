package core

import (
	"testing"
)

func TestCreateNone(t *testing.T) {
	o := None[string]()
	if o.IsSome() {
		t.Fatal("it must be None[string]")
	}
}

func TestCreateSomeForString(t *testing.T) {
	o := Some("foo")
	if o.IsNone() {
		t.Fatal("it must be Some[string]")
	}
}

func TestCreateSomeOfString(t *testing.T) {
	o := Some("foo")
	if o.Unwrap() != "foo" {
		t.Fatal("a value must be equal")
	}
}

func TestMapNoneForNone(t *testing.T) {
	o1 := None[string]()
	o2 := MapNone[string, int](o1, func() int {
		return 1
	})

	if o2.IsNone() {
		t.Fatal("it must be Some[int]")
	}

	if o2.Unwrap() != 1 {
		t.Fatal("it must be equal 0")
	}
}

func TestMapNoneForSome(t *testing.T) {
	o1 := Some[string]("foo")
	o2 := MapNone[string, int](o1, func() int {
		return 1
	})

	if o2.IsSome() {
		t.Fatal("it must be None[int]")
	}
}

func TestMapSomeForSome(t *testing.T) {
	s := "string"
	o1 := Some[string](s)
	o2 := MapSome[string, int](o1, func(s string) int {
		return len(s)
	})

	if o2.IsNone() {
		t.Fatal("it must be Some[int]")
	}

	if o2.Unwrap() != len(s) {
		t.Fatalf("it mut be equal %d", len(s))
	}
}

func TestMapSomeForNone(t *testing.T) {
	o1 := None[string]()
	o2 := MapSome[string, int](o1, func(s string) int {
		return len(s)
	})

	if o2.IsSome() {
		t.Fatal("it must be None[int]")
	}
}
