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

func TestUnwrapForOk(t *testing.T) {
	s := "foo"
	r := Ok[string](s)
	if r.Ok().Unwrap() != s {
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


func TestUnwrapOrPtrForOk(t *testing.T) {
	s := "foo"
	s2 := "bar"
	r := Ok[string](s)

	value := r.UnwrapOrPtr(&s2)

	if *value == s2 {
		t.Fatalf("it must be not equal %x", s2)
	}

	if *value != s {
		t.Fatalf("it must be equal %x", s)
	}

}


func TestUnwrapOrPtrForError(t *testing.T) {
	s := "foo"
	r := Err[string](errors.New("error"))
	v := r.UnwrapOrPtr(&s)
	if v != &s {
		t.Fatalf("it must be equal %x", s)
	}

	if r.Error().Unwrap() == nil {
		t.Fatal("it must be not equal nil")
	}
}