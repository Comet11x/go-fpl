package core

type Try[T any] interface {
	// Returns true if the execution finishes successfully.
	IsSuccess() bool

	// Returns true it the execution finishes with a failure.
	IsFailure() bool

	// Calls a callback function if the execution is successful.
	IfSuccess(func(value T)) Try[T]

	// Calls a callback function if the execution is successful.
	IfSuccessAsPtr(func(value *T)) Try[T]

	// Calls a callback function if the execution is failed.
	IfFailure(func(value any)) Try[T]

	// Calls a callback function if the execution is failed.
	IfFailureAsPtr(func(value *any)) Try[T]

	// Returns a value of the successful execution of a function.
	// It can be empty if the execution is failed.
	Success() Option[T]

	// Returns a value of the failed execution of a function.
	// It can be empty if the execution finishes successfully.
	Failure() Option[any]

	// Translates Try[T] to Result[T]
	AsResult(errorFactory ...func(any) error) Result[T]
}
