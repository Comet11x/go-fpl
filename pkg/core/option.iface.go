package core

// Option is the interface which
type Option[T any] interface {
	// Returns true if the option is Some
	IsSome() bool

	// Returns true if the option is None, otherwise is returns false
	IsNone() bool

	// Calls the function if the options is Some[T]
	IfSome(func(value T)) Option[T]

	// Calls the function if the option is Some[T]
	IfSomeAsPtr(func(value *T)) Option[T]

	// Calls the function if the option is None[T]
	IfNone(func()) Option[T]

	// Returns the value if the option is Some[T], otherwise it returns the given value
	UnwrapOr(value T) T

	// Returns the value if the option is Some[T], otherwise it returns the value from the given function
	UnwrapOrValueFrom(func() T) T

	// Returns the value if the option is Some[T], otherwise it returns the given value
	UnwrapAsPtrOr(value *T) *T

	// Returns a pointer of the value if the option is Some[T], otherwise it returns the pointer from the given function
	UnwrapAsPtrOrPtrFrom(func() *T) *T

	// Returns the stored value if it exists, otherwise it raises panic
	Unwrap() T

	// Returns a pointer to the stored value if it exists, otherwise it raises panic
	UnwrapAsPtr() *T

	AsTuple() (T, bool)

	AsTupleOfPtr() (*T, bool)

	Swap(value T) T
	SwapFrom(func() T) T
	SwapAsPtr(value *T) T
	SwapAsPtrFrom(func() *T) T
	MapSome(fn func(v T) T) Option[T]
	MapSomeFrom(fn func(v T) Option[T]) Option[T]
	MapNone(fn func() T) Option[T]
	MapNoneFrom(fn func() Option[T]) Option[T]
}
