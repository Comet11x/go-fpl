package types

type Unwrap[T any] interface {
	Unwrap() T
}
