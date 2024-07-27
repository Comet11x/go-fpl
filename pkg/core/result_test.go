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
		t.Fatalf("it must be equal %s", s)
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

func TestUnwrapBar(t *testing.T) {
	s := "foo"
	r := Ok[string](s)

	value := r.UnwrapOr("bar")

	if value != s {
		t.Fatalf("it must be equal %s", s)
	}
	
}

func TestOkUnwrap(t *testing.T) {
	s := "bar"
	r := Ok[string](s)
	if r.Ok().Unwrap() != s {
		t.Fatalf("it must be equal %s", s)
	}

	if r.Error().Unwrap() != nil {
		t.Fatal("it must be equal nil")
	}
}

func TestErrorUnwrap(t *testing.T) {
	r := Err[string](errors.New("error"))
	if len(r.Ok().Unwrap()) != 0 {
		t.Fatal("it must be equal 0")
	}

	if r.Error().Unwrap() == nil {
		t.Fatal("it must be not equal nil")
	}
}