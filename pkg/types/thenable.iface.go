package types

type Thenable[T any] interface {
	Awaiter[T]
	Unwrap[T]

	IsFulfilled() bool
	IsPending() bool
	IsRejected() bool

	Then(func(T)) Thenable[T]
	Catch(func(error)) Thenable[T]
}
