package core

type Result[T any, E any] interface {

	// Returns true if the item of Result is Ok[T, E].
	IsOk() bool

	// Returns false if the item of Result is Err[T, E].
	IsErr() bool

	// Calls a callback if the result is Ok[T, E].
	// fn is a function which gets the ok value of T.
	// It returns itself.
	IfOk(fn func(value T)) Result[T, E]

	// Calls a callback if the result is Ok[T, E]
	// fn is a function which gets the ok value as a pointer.
	// It returns itself.
	IfOkAsPtr(func(value *T)) Result[T, E]

	// Calls a callback if the result is Err[T, E].
	// fn is a function which gets the err value.
	// It returns itself.
	IfErr(fn func(err E)) Result[T, E]

	// Calls a callback if the result is Error.
	// fn is a function which gets the err value as a pointer.
	// It returns itself.
	IfErrAsPtr(fn func(err *E)) Result[T, E]

	// Returns the contained ok value if the result is Ok[T, E], otherwise it raises panic with the given message.
	Expect(msg string) T

	// Returns the contained ok value if the result is Ok[T, E], otherwise it raises panic.
	Unwrap() T

	// Returns the contained ok if the result is Ok[T, E], otherwise it returns the given value.
	// value is an alternative value which may be returned if the result is Err[T, E].
	UnwrapOr(value T) T

	// Returns the contained ok if the result is Ok[T, E], otherwise it returns a default value of T.
	UnwrapOrDefault() T

	// Returns the contained ok if the result is Ok[T, E], otherwise it returns a default value of T which is returned by the callback function
	// fn is a callback function
	UnwrapOrElse(fn func() T) T

	// Returns a pointer of the stored value or raise
	ExpectAsPtr(msg string) *T

	// Returns a pointer of a value of the result or an alternative pointer
	UnwrapAsPtrOr(value *T) *T

	// Returns a pointer of a value of the result or a pointer from the function
	UnwrapAsPtrOrElse(fn func() *T) *T

	// Returns an contained error value if the result is the Err[T, E], otherwise it raises a panic.
	UnwrapErr() E

	// Returns an error value if the result is the Err[T, E], otherwise it returns an alternative error which is given as the parameter of this method.
	// err is an alternative error which is returned if the result is Ok[T, E]
	UnwrapErrOr(err E) E

	// Returns the contained error value or an error value by default
	UnwrapErrOrDefault() E

	// Returns the contained Ok[T] value as a pointer
	UnwrapAsPtr() *T

	// Returns a tuple which has the contained Ok[T] value
	AsTuple() (T, E)

	// Returns a tuple of pointers
	AsTupleOfPtr() (*T, *E)

	// Returns a tuple of the ok value and the error value
	AsEither() Either[T, E]

	// Returns the ok value and the error value as pointers which are wrapped in Either.
	AsEitherPtr() Either[*T, *E]

	// Returns the contained ok value which is wrapped in Option.
	Ok() Option[T]

	// Returns the contained ok value as a pointer of T which is wrapped in Option
	OkAsPtr() Option[*T]

	// Returns the contained error value which is wrapped in Option
	Err() Option[E]

	// Returns the contained error value as a pointer of E which is wrapped in Option
	ErrAsPtr() Option[*E]
}
