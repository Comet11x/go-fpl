package core

type Option[T any] interface {
	IsSome() bool
	IsNone() bool
	UnwrapOr(value T) T
	UnwrapOrValueFrom(func() T) T
	UnwrapAsPtrOr(value *T) *T
	UnwrapAsPtrOrPtrFrom(func() *T) *T
	Unwrap() T
	UnwrapAsPtr() *T
	ToTuple() (T, bool)
	ToTupleAsPtr() (*T, bool)
	Swap(value T) T
	SwapFrom(func() T) T
	SwapAsPtr(value *T) T
	SwapAsPtrFrom(func() *T) T
	MapIfSome(fn func(v T) T) Option[T]
	MapIfNone(fn func() T) Option[T]
}
