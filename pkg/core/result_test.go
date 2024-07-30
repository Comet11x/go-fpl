package core

import (
	"errors"
	"testing"
)

func TestCreateOkResult(t *testing.T) {
	s := "foo"
	r := Ok[string](s)
	if !r.IsOk() || r.IsError() {
		t.Fatal("it must be Ok[string]")
	}
}

func TestCreateErrResult(t *testing.T) {
	r := Err[string](errors.New("error"))

	if r.IsOk() || !r.IsError() {
		t.Fatal("it must be Err[string]")
	}
}

func TestUnwrapAsPtrForOk(t *testing.T) {
	i := 1
	r := Ok[int](i)
	o := r.UnwrapAsPtr()
	if i != *o {
		t.Fatalf("it must be equal %d", i)
	}
}

func TestUnwrapForOk(t *testing.T) {
	s := "foo"
	r := Ok[string](s)
	if r.Unwrap() != s {
		t.Fatalf("it must be equal %s ", s)
	}

	if r.Error().Unwrap() != nil {
		t.Fatal("it must be equal nil")
	}
}

func TestUnwrapForError(t *testing.T) {
	r := Err[string](errors.New("error"))
	if len(r.Ok().Unwrap()) != 0 {
		t.Fatal("it must be equal 0")
	}

	if r.Error().Unwrap() == nil {
		t.Fatal("it must be not equal nil")
	}
}

func TestUnwrapOrForOk(t *testing.T) {
	s := "foo"
	r := Ok[string](s)

	value := r.UnwrapOr("bar")

	if value != s {
		t.Fatalf("it must be equal %s", s)
	}

}

func TestUnwrapForOK(t *testing.T) {
	s := "bar"
	r := Ok[string](s)
	if r.Ok().Unwrap() != s {
		t.Fatalf("it must be equal %s", s)
	}

	if r.Error().Unwrap() != nil {
		t.Fatal("it must be equal nil")
	}
}

func TestUnwrapOrForError(t *testing.T) {
	s := "foo"
	r := Err[string](errors.New("error"))
	v := r.UnwrapOr(s)
	if v != s {
		t.Fatalf("it must be equal %s", s)
	}

	if r.Error().Unwrap() == nil {
		t.Fatal("it must be not equal nil")
	}
}

func TestUnwrapAsPtrOrForOk(t *testing.T) {
	s := "foo"
	s2 := "bar"
	r := Ok[string](s)

	value := r.UnwrapAsPtrOr(&s2)

	if *value == s2 {
		t.Fatalf("it must be not equal %x", s2)
	}

	if *value != s {
		t.Fatalf("it must be equal %x", s)
	}

}

func TestUnwrapAsPtrOrForError(t *testing.T) {
	s := "foo"
	r := Err[string](errors.New("error"))
	v := r.UnwrapAsPtrOr(&s)
	if v != &s {
		t.Fatalf("it must be equal %x", s)
	}

	if r.Error().Unwrap() == nil {
		t.Fatal("it must be not equal nil")
	}
}

func TestOkPtrForOk(t *testing.T) {
	v := 1
	r := Ok[int](v)

	o := r.OkPtr()

	if o.IsNone() || !o.IsSome() {
		t.Fatal("it mut be Some[int]")
	}
}

func TestOkPtrForErr(t *testing.T) {
	r := Err[int](errors.New("error"))

	o := r.OkPtr()

	if !o.IsNone() || o.IsSome() {
		t.Fatal("it mut be None[int]")
	}
}

func TestToTupleForOk(t *testing.T) {
	v1 := 1
	r := Ok[int](v1)

	v2, err := r.ToTuple()

	if v1 != v2 {
		t.Fatal("it must be equal")
	}

	if err != nil {
		t.Fatal("it must be not equal")
	}
}

func TestToTupleForErr(t *testing.T) {
	r := Err[int](errors.New("error"))

	_, err := r.ToTuple()

	if err == nil {
		t.Fatal("it must be not equal")
	}
}
