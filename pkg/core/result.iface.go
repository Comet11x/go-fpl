package core

type Result[T any] interface {
	IsOk() bool
	IsError() bool

	UnwrapOr(value T) T

	UnwrapOrFrom(c func() T) T

	UnwrapAsPtrOr(value *T) *T
	UnwrapAsPtrOrFrom(c func() *T) *T

	Unwrap() T
	UnwrapAsPtr() *T

	ToTuple() (T, error)
	ToTuplePtr() (*T, error)

	ToEither() Either[T, error]
	ToEitherPtr() Either[*T, error]

	Ok() Option[T]
	OkPtr() Option[*T]

	Error() Option[error]
}
