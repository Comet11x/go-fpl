package async

type Promise[T any] interface {
	// Extends Awaiter[T]
	Awaiter[T]

	// Adds a callback function which will be call if the promise has the RESOLVE state
	Then(func(T)) Promise[T]

	// Adds a callback function which will be call if the promise has the REJECTED state
	Catch(func(any)) Promise[T]

	// Adds a callback function which will be call if the promise has the RESOLVE state
	Finally(func()) Promise[T]

	// Returns true if the promise has the pending state
	IsPending() bool

	// Returns true if the promise has the resolved state
	IsResolved() bool

	// Returns true if the promise has the rejected state
	IsRejected() bool

	// Returns a status of the promise
	Status() string

	// Returns an item of Future[T]
	Future() Future[T]
}
