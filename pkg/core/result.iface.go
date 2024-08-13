package core

type Result[T any] interface {

	// Returns true if the item of Result is Ok[T]
	IsOk() bool

	// Returns false if the item of Result is Err[T]
	IsError() bool

	// Returns a value of the result or an alternative value
	UnwrapOr(value T) T

	// Returns a default value
	UnwrapOrDefault() T

	// Returns a value of the result or a value from the function
	UnwrapOrValueFrom(c func() T) T

	// Returns a pointer of a value of the result or an alternative pointer
	UnwrapAsPtrOr(value *T) *T

	// Returns a pointer of a value of the result or a pointer from the function
	UnwrapAsPtrOrPtrFrom(c func() *T) *T

	// Returns the contained Ok[T] value.
	Unwrap() T

	// Returns an error of the result
	UnwrapErr() error

	// Returns an error of the result or an alternative error
	UnwrapErrOr(error) error

	//
	UnwrapErrOrDefault() error

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
