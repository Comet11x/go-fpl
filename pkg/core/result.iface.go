package core

type Result[T any] interface {

	// Returns true if the item of Result is Ok[T]
	IsOk() bool

	// Returns false if the item of Result is Err[T]
	IsError() bool

	// Returns a value of the result or a default value
	UnwrapOr(value T) T

	UnwrapOrDefault() T

	UnwrapOrFrom(c func() T) T

	UnwrapAsPtrOr(value *T) *T
	UnwrapAsPtrOrFrom(c func() *T) *T

	// Returns the contained Ok[T] value
	// It Panics
	Unwrap() T

	// Returns the contained Ok[T] value as a pointer
	UnwrapAsPtr() *T

	// Returns a tuple which has the contained Ok[T] value
	ToTuple() (T, error)
	ToTupleAsPtr() (*T, error)

	ToEither() Either[T, error]
	ToEitherPtr() Either[*T, error]

	Ok() Option[T]
	OkPtr() Option[*T]

	Error() Option[error]
}
