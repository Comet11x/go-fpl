package slice

import (
	"testing"
)

func TestHead(t *testing.T) {
	s := []int{1, 2, 3}
	if Head(s).Unwrap() != 1 {
		t.Fatal("It must be equal 1")
	}
}

func TestLast(t *testing.T) {
	s := []int{1, 2, 3}
	if Last(s).Unwrap() != 3 {
		t.Fatal("It must be equal 1")
	}
}

func TestTail(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := Tail(s1)

	if len(s2) != 2 {
		t.Fatal("the length of the slice must be equal 2")
	}

	if s2[0] != 2 && s2[1] != 3 {
		t.Fatal("the slice has incorrect values")
	}
}

func TestMap(t *testing.T) {
	s1 := []string{"a", "aa", "aaa"}
	s2 := Map(s1, func(v string) int { return len(v) })

	if len(s2) != 3 {
		t.Fatal("the length of the slice must be equal 3")
	}

	if s2[0] != 1 && s2[1] != 2 && s2[2] == 3 {
		t.Fatal("the slice has incorrect values")
	}
}
