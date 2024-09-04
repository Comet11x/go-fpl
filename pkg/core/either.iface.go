package core

type Either[L any, R any] interface {
	IsLeft() bool
	IsRight() bool
	IfLeft(func(L)) bool
	IfRight(func(R)) bool
	Left() Option[L]
	LeftAsPtr() Option[*L]
	Right() Option[R]
	RightAsPtr() Option[*R]
	ToTuple() (L, R)
	ToTuplePtr() (*L, *R)
	UnwrapLeft() L
	UnwrapLeftOr(L) L
	UnwrapLeftOrFrom(func() L) L
	UnwrapLeftAsPtr() *L
	UnwrapLeftAsPtrOrFrom(func() *L) *L

	UnwrapRight() R
	UnwrapRightOr(R) R
	UnwrapRightOrFrom(func() R) R
	UnwrapRightAsPtr() *R
	UnwrapRightAsPtrOrFrom(func() *R) *R
}
