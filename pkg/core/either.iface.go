package core

type Either[L any, R any] interface {
	// Returns true if the either is Left, otherwise it returns false.
	IsLeft() bool

	// Returns true if the either is Right, otherwise it returns false.
	IsRight() bool

	// Calls the function if the either is Left.
	// fn is a function which is called by the either.
	// Returns itself.
	IfLeft(fn func(value L)) Either[L, R]

	// Calls the function if the the either if Left.
	// fn is a function that is called by the either
	IfLeftAsPtr(fn func(value *L)) Either[L, R]

	IfRight(func(value R)) Either[L, R]
	IfRightAsPtr(func(value *R)) Either[L, R]

	Left() Option[L]
	LeftAsPtr() Option[*L]
	Right() Option[R]
	RightAsPtr() Option[*R]
	ToTuple() (L, R)
	ToTuplePtr() (*L, *R)
	UnwrapLeft() L
	UnwrapLeftOr(value L) L
	UnwrapLeftOrFrom(func() L) L
	UnwrapLeftAsPtr() *L
	UnwrapLeftAsPtrOrFrom(func() *L) *L

	UnwrapRight() R
	UnwrapRightOr(value R) R
	UnwrapRightOrFrom(func() R) R
	UnwrapRightAsPtr() *R
	UnwrapRightAsPtrOrFrom(func() *R) *R
}
