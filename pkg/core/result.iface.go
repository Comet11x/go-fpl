package core

type Result[T any] interface {
	IsOk() bool
	IsError() bool

	UnwrapOr(value T) T
	UnwrapOrPtr(value *T) *T

	Unwrap() T
	UnwrapPtr() *T

	ToTuple() (T, error)
	ToTuplePtr() (*T, error)

	ToEither() Either[T, error]
	ToEitherPtr() Either[*T, error]

	Ok() Option[T]
	OkPtr() Option[*T]

	Error() Option[error]
}
