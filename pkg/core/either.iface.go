package core

type Either[A any, B any] interface {
	IsLeft() bool
	IsRight() bool
	Left() Option[A]
	LeftAsPtr() Option[*A]
	Right() Option[B]
	RightAsPtr() Option[*B]
	ToTuple() (A, B)
	ToTuplePtr() (*A, *B)
	UnwrapLeft() A
	UnwrapLeftOr(A) A
	UnwrapLeftOrFrom(func() A) A
	UnwrapLeftAsPtr() *A
	UnwrapLeftAsPtrOrFrom(func() *A) *A

	UnwrapRight() B
	UnwrapRightOr(B) B
	UnwrapRightOrFrom(func() B) B
	UnwrapRightAsPtr() *B
	UnwrapRightAsPtrOrFrom(func() *B) *B
}
