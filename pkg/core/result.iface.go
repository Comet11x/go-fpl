package core

type Result[T any] interface {
	IsOk() bool
	IsError() bool
	UnwrapOr(value T) T

	Unwrap() T

	UnwrapPtr() *T

	ToTuple() (T, error)

	ToEither() Either[T, error]

	Ok() Option[T]

	Error() Option[error]
}
