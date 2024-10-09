package slice

import (
	"github.com/comet11x/go-fpl/pkg/algorithm"
	"github.com/comet11x/go-fpl/pkg/core"
	"github.com/comet11x/go-fpl/pkg/types"
)

func Head[S ~[]T, T any](iterable S) core.Option[T] {
	if len(iterable) > 0 {
		return core.Some(iterable[0])
	} else {
		return core.None[T]()
	}
}

func Last[S ~[]T, T any](iterable S) core.Option[T] {
	if len(iterable) > 0 {
		return core.Some(iterable[len(iterable)-1])
	} else {
		return core.None[T]()
	}
}

func Tail[S ~[]T, T any](iterable S) S {
	if len(iterable) > 0 {
		return iterable[1:]
	} else {
		return make([]T, 0)
	}
}

func ForEach[T any](iterable []T, callback func(T)) {
	for _, v := range iterable {
		callback(v)
	}
}

func Map[T any, U any](iterable []T, callback func(T) U) []U {
	out := make([]U, len(iterable))
	for i, v := range iterable {
		out[i] = callback(v)
	}
	return out
}

func Reduce[T any, R any](iterable []T, callback func(current T, previous ...R) R, initial ...R) R {
	var ret R
	for i, v := range iterable {
		if i == 0 {
			ret = callback(v, initial...)
		} else {
			ret = callback(v, ret)
		}
	}
	return ret
}

func PartialReduce[T any, R any](iterable []T, callback func(current T, previous ...R) (R, bool), initial ...R) R {
	var ret R
	var done bool
	for i, v := range iterable {
		if i == 0 {
			ret, done = callback(v, initial...)
		} else {
			ret, done = callback(v, ret)
		}
		if done {
			break
		}
	}
	return ret
}

func Min[T types.Ordered](iterable []T) core.Option[T] {
	if len(iterable) == 0 {
		return core.None[T]()
	}
	min := iterable[0]
	for i := 1; i < len(iterable); i++ {
		if min > iterable[i] {
			min = iterable[i]
		}
	}
	return core.Some(min)
}

func Max[T types.Ordered](iterable []T) core.Option[T] {
	if len(iterable) == 0 {
		return core.None[T]()
	}
	max := iterable[0]
	for i := 1; i < len(iterable); i++ {
		if max < iterable[i] {
			max = iterable[i]
		}
	}
	return core.Some(max)
}

func Sum[T types.Number](iterable []T) T {
	var sum T
	for _, v := range iterable {
		sum += v
	}
	return sum
}

func Filter[T any](iterable []T, callback func(T) bool) []T {
	out := make([]T, 0)
	for _, v := range iterable {
		if callback(v) {
			out = append(out, v)
		}
	}
	return out
}

func Some[T any](iterable []T, callback func(T) bool) bool {
	out := false
	for _, v := range iterable {
		if callback(v) {
			out = true
			break
		}
	}
	return out
}

func Every[T any](iterable []T, callback func(T) bool) bool {
	out := true
	for _, v := range iterable {
		if !callback(v) {
			out = false
			break
		}
	}
	return out
}

func FindIndex[T any](iterable []T, callback func(T) bool) int {
	idx := -1
	for i, v := range iterable {
		if callback(v) {
			idx = i
			break
		}
	}
	return idx
}

func Count[T any](iterable []T, callback func(T) bool) uint {
	var n uint = 0
	for _, v := range iterable {
		if callback(v) {
			n += 1
		}
	}
	return n
}

func Zip[T any](f []T, s []T) [][]T {
	var out [][]T
	m := algorithm.Min(len(f), len(s))
	out = make([][]T, 0, m)
	for i := 0; i < m; i++ {
		out = append(out, []T{f[i], s[i]})
	}

	return out
}
