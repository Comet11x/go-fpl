package core

type Either[A any, B any] interface {
	IsRight() bool
	IsLeft() bool
	Right() Option[B]
	Left() Option[A]
	ToTuple() (A, B)
	ToTuplePtr() (*A, *B)
	UnwrapLeft() A
	UnwrapLeftPtr() *A
	UnwrapRight() B
	UnwrapRightPtr() *B
}
