package core

type Try[T any] interface {
	IsSuccess() bool

	IsFailure() bool

	Success() Option[T]

	Failure() Option[any]

	AsResult(errorFactory ...func(any) error) Result[T]
}
