package core

type Either[L any, R any] interface {
	IsLeft() bool
	IsRight() bool
	IfLeft(func(value L)) Either[L, R]
	IfLeftAsPtr(func(value *L)) Either[L, R]
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
