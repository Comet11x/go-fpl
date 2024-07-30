package core

type either[A any, B any] struct {
	t     int
	left  A
	right B
}

func (e *either[A, B]) IsRight() bool {
	return e.t == _RIGHT
}

func (e *either[A, B]) IsLeft() bool {
	return e.t == _LEFT
}

func (e *either[A, B]) Right() Option[B] {
	if e.IsRight() {
		return Some(e.right)
	} else {
		return None[B]()
	}
}

func (e *either[A, B]) RightAsPtr() Option[*B] {
	if e.IsRight() {
		return Some(&e.right)
	} else {
		return None[*B]()
	}
}

func (e *either[A, B]) Left() Option[A] {
	if e.IsLeft() {
		return Some(e.left)
	} else {
		return None[A]()
	}
}

func (e *either[A, B]) LeftAsPtr() Option[*A] {
	if e.IsLeft() {
		return Some(&e.left)
	} else {
		return None[*A]()
	}
}

func (e *either[A, B]) ToTuple() (A, B) {
	return e.left, e.right
}

func (e *either[A, B]) ToTuplePtr() (*A, *B) {
	return &e.left, &e.right
}

func (e *either[A, B]) UnwrapLeft() A {
	return e.left
}

func (e *either[A, B]) UnwrapLeftPtr() *A {
	return &e.left
}

func (e *either[A, B]) UnwrapRight() B {
	return e.right
}

func (e *either[A, B]) UnwrapRightPtr() *B {
	return &e.right
}
