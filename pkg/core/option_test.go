package core

import (
	"testing"
)

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
