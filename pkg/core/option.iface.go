package core

type Option[T any] interface {
	IsSome() bool
	IsNone() bool
	IfSome(func(value T)) Option[T]
	IfNone(func()) Option[T]
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
	MapSome(fn func(v T) T) Option[T]
	MapSomeFrom(fn func(v T) Option[T]) Option[T]
	MapNone(fn func() T) Option[T]
	MapNoneFrom(fn func() Option[T]) Option[T]
}
