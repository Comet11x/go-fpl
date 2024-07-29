package slice

import (
	"github.com/comet11x/go-fpl/pkg/core"
)

func Head[S ~[]T, T any](s S) core.Option[T] {
	if len(s) > 0 {
		return core.Some(s[0])
	} else {
		return core.None[T]()
	}
}

func Last[S ~[]T, T any](s S) core.Option[T] {
	if len(s) > 0 {
		return core.Some(s[len(s)-1])
	} else {
		return core.None[T]()
	}
}

func Tail[S ~[]T, T any](s S) S {
	if len(s) > 0 {
		return s[1:]
	} else {
		return make([]T, 0)
	}
}

func Map[S ~[]T, T any, R ~[]U, U any](iterable S, clb func(T) U) {
	out := make(R, len(iterable))
	for i, v := range iterable {
		out[i] = clb(v)
	}

}

func Filter[T any](iterable []T, clb func(T) bool) []T {
	out := make([]T, 0)
	for _, v := range iterable {
		if clb(v) {
			out = append(out, v)
		}
	}
	return out
}

func Some[T any](iterable []T, clb func(T) bool) bool {
	out := false
	for _, v := range iterable {
		if clb(v) {
			out = true
			break
		}
	}
	return out
}

func Every[T any](iterable []T, clb func(T) bool) bool {
	out := true
	for _, v := range iterable {
		if !clb(v) {
			out = false
			break
		}
	}
	return out
}

func FindIndex[T any](iterable []T, clb func(T) bool) int {
	idx := -1
	for i, v := range iterable {
		if clb(v) {
			idx = i
			break
		}
	}
	return idx
}

func Count[T any](iterable []T, clb func(T) bool) uint {
	var c uint = 0

	for _, v := range iterable {
		if clb(v) {
			c += 1
		}
	}

	return c
}
