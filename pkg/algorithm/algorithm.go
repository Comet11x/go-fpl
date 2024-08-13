package algorithm

import (
	"github.com/comet11x/go-fpl/pkg/types"
)

func Min[T types.Ordered](f T, s T) T {
	if f > s {
		return s
	} else {
		return f
	}
}

func Max[T types.Ordered](f T, s T) T {
	if f < s {
		return s
	} else {
		return f
	}
}

func IsEqual[T types.Equal](f T, s T) bool {
	return f == s
}
