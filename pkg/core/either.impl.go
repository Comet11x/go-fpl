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

func (e *either[A, B]) IfLeft(fn func(A)) bool {
	cond := e.IsLeft()
	if cond {
		fn(e.left)
	}
	return cond
}

func (e *either[A, B]) IfRight(fn func(B)) bool {
	cond := e.IsRight()
	if cond {
		fn(e.right)
	}
	return cond
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

func (e *either[A, B]) UnwrapLeftOr(v A) A {
	if e.IsLeft() {
		return e.left
	} else {
		return v
	}
}

func (e *either[A, B]) UnwrapLeftOrFrom(c func() A) A {
	if e.IsLeft() {
		return e.left
	} else {
		return c()
	}
}

func (e *either[A, B]) UnwrapLeftAsPtr() *A {
	return &e.left
}

func (e *either[A, B]) UnwrapLeftAsPtrOrFrom(c func() *A) *A {
	if e.IsLeft() {
		return &e.left
	} else {
		return c()
	}
}

func (e *either[A, B]) UnwrapRight() B {
	return e.right
}

func (e *either[A, B]) UnwrapRightOr(v B) B {
	if e.IsRight() {
		return e.right
	} else {
		return v
	}
}

func (e *either[A, B]) UnwrapRightOrFrom(c func() B) B {
	if e.IsRight() {
		return e.right
	} else {
		return c()
	}
}

func (e *either[A, B]) UnwrapRightAsPtr() *B {
	return &e.right
}

func (e *either[A, B]) UnwrapRightAsPtrOrFrom(c func() *B) *B {
	if e.IsRight() {
		return &e.right
	} else {
		return c()
	}
}
