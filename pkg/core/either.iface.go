package core

type Either[L any, R any] interface {
	// Returns true if the `Either` has a `Left` value, otherwise it returns false.
	IsLeft() bool

	// Returns true if the `Either` is `Right`, otherwise it returns false.
	IsRight() bool

	// Calls the `fn` function ad passes a `Left` value to `fn` if the `Either` has the one.
	//
	// Returns itself.
	IfLeft(fn func(value L)) Either[L, R]

	// Calls the `fn` function ad passes a `Left` value to `fn` if a either if Left.
	//
	// Returns itself.
	IfLeftAsPtr(fn func(value *L)) Either[L, R]

	// Calls the function `fn` and passes a `Right` value to `fn` if a instance which implements `Either` has a `Right` value.
	//
	// Returns itself.
	IfRight(fn func(value R)) Either[L, R]

	// `IfRightAsPtr` calls the function if the instance which implements `Either` has a `Right` value.
	//
	// Returns itself.
	IfRightAsPtr(fn func(value *R)) Either[L, R]

	// Returns `Option[L]` which will be `Some` if `Either` has a `Left` value, otherwise it will be `None`.
	Left() Option[L]

	// Returns `Option[*L]` which will be `Some` if `Either` has a `Left` value, otherwise it will be `None`.
	LeftAsPtr() Option[*L]

	// Returns `Option[R]` which will be `Some` if `Either` has a `Right` value, otherwise it will be `None`.
	Right() Option[R]

	// Returns `Option[*R]` which will be `Some` if `Either` has a `Right` value, otherwise it will be `None`.
	RightAsPtr() Option[*R]

	// Returns a tuple which has `Left` and `Right values.
	AsTuple() (L, R)

	// Returns a tuple which has pointers to `Left` and `Right values.
	AsTuplePtr() (*L, *R)

	// Returns the stored `Left` value of `Either` if `Ether` has the one,
	// otherwise it rases a panic with this message "called `Either.UnwrapLeft()` on a `Right` value".
	//
	// Warning! It can raise a panic.
	UnwrapLeft() L

	// Returns the stored `Left` value or the given `Left` value.
	UnwrapLeftOr(value L) L

	// Returns the stored `Left` value or a `Left` value which is taken from a `fn` function.
	UnwrapLeftOrFrom(fn func() L) L

	// Returns a pinter to the stored `Left` value of `Either` if `Ether` has the one,
	// otherwise it rases a panic with this message "called `Either.UnwrapLeftAsPtr()` on a `Right` value".
	//
	// Warning! It can raise a panic.
	UnwrapLeftAsPtr() *L

	// Returns a pinter to the stored `Left` value or the given pointer to a `Left` value.
	UnwrapLeftAsPtrOr(value *L) *L

	// Returns a pinter to the stored `Left` value or a pointer to a `Left` value which is taken from a `fn` function.
	UnwrapLeftAsPtrOrFrom(func() *L) *L

	// Returns the stored `Right` value of `Either` if `Ether` has the one,
	// otherwise it rases a panic with this message "called `Either.UnwrapRight()` on a `Left` value".
	//
	// Warning! It can raise a panic.
	UnwrapRight() R

	// Returns the stored `Right` value or the given `Right` value.
	UnwrapRightOr(value R) R

	// Returns the stored `Right` value or a `Right` value which is taken from a `fn` function.
	UnwrapRightOrFrom(fn func() R) R

	// Returns a pinter to the stored `Right` value of `Either` if `Ether` has the one,
	// otherwise it rases a panic with this message "called `Either.UnwrapRightAsPtr()` on a `Left` value".
	//
	// Warning! It can raise a panic.
	UnwrapRightAsPtr() *R

	// Returns a pinter to the stored `Right` value or the given pointer to a `Right` value.
	UnwrapRightAsPtrOr(value *R) *R

	// Returns a pinter to the stored `Right` value or a pointer to a `Right` value which is taken from a `fn` function.
	UnwrapRightAsPtrOrFrom(func() *R) *R
}
