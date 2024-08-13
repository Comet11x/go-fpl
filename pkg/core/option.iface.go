package core

type Option[T any] interface {
	IsSome() bool
	IsNone() bool
	UnwrapOr(value T) T
	UnwrapOrFrom(func() T) T
	UnwrapAsPtrOr(value *T) *T
	UnwrapAsPtrOrFrom(func() *T) *T
	Unwrap() T
	UnwrapAsPtr() *T
	ToTuple() (T, bool)
	ToTupleAsPtr() (*T, bool)
	Swap(value T) T
	SwapFrom(func() T) T
	SwapAsPtr(value *T) T
	SwapAsPtrFrom(func() *T) T
}
