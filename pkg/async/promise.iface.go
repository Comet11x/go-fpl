package async

type Promise[T any] interface {
	// extends Awaiter[T]
	Awaiter[T]

	//
	Then(func(T)) Promise[T]
	Catch(func(any)) Promise[T]
	Finally(func()) Promise[T]
	IsPending() bool
	IsResolve() bool
	IsRejected() bool
	Status() string
}
