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
	UnwrapLeftPtr() *A
	UnwrapRight() B
	UnwrapRightPtr() *B
}
