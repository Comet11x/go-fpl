package hashmap

import (
	"github.com/comet11x/go-fpl/pkg/algorithm"
	"github.com/comet11x/go-fpl/pkg/algorithm/slice"
	"github.com/comet11x/go-fpl/pkg/core"
	"github.com/comet11x/go-fpl/pkg/sync"
)

func Get[K comparable, T any](m map[K]T, k K, mtx ...sync.RWLocker) core.Option[T] {
	l := slice.Head(mtx).UnwrapOrFrom(sync.FakeRWLocker)
	l.Lock()
	v, ok := m[k]
	l.Unlock()
	return core.OptionFrom[T](v, ok)
}

func Set[K comparable, T any](m map[K]T, k K, v T, mtx ...sync.RWLocker) {
	l := slice.Head(mtx).UnwrapOrFrom(sync.FakeRWLocker)
	l.Lock()
	m[k] = v
	l.Unlock()
}

func Keys[K comparable, T any](m map[K]T, mtx ...sync.RWLocker) []K {
	var out []K
	l := slice.Head(mtx).UnwrapOrFrom(sync.FakeRWLocker)
	l.RLock()
	out = make([]K, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	l.RUnlock()
	return out
}

func Values[K comparable, T any](m map[K]T, mtx ...sync.RWLocker) []T {
	var out []T
	l := slice.Head(mtx).UnwrapOrFrom(sync.FakeRWLocker)

	l.RLock()
	out = make([]T, 0, len(m))
	for _, v := range m {
		out = append(out, v)
	}
	l.RUnlock()

	return out
}

func Zip[K comparable, T any](keys []K, values []T, mtx ...sync.RWLocker) map[K]T {
	var out map[K]T
	l := slice.Head(mtx).UnwrapOrFrom(sync.FakeRWLocker)

	l.RLock()
	m := algorithm.Min(len(keys), len(values))
	out = make(map[K]T, m)
	for i := 0; i < m; i++ {
		out[keys[i]] = values[i]
	}
	l.RUnlock()

	return out
}
