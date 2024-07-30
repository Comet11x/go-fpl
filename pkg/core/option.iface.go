package core

type Option[T any] interface {
	IsSome() bool
	IsNone() bool
	UnwrapOr(value T) T
	UnwrapAsPtrOr(value *T) *T
	Unwrap() T
	UnwrapAsPtr() *T
	Swap(value T) T
	SwapFromPtr(value *T) T
}
