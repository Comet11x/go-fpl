package core

type either[L any, R any] struct {
	t     int
	left  L
	right R
}

func (e *either[L, R]) IsRight() bool {
	return e.t == _RIGHT
}

func (e *either[L, R]) IsLeft() bool {
	return e.t == _LEFT
}

func (e *either[L, R]) IfLeft(fn func(L)) Either[L, R] {
	if e.IsLeft() {
		fn(e.left)
	}
	return e
}

func (e *either[L, R]) IfLeftAsPtr(fn func(*L)) Either[L, R] {
	if e.IsLeft() {
		fn(&e.left)
	}
	return e
}

func (e *either[L, R]) IfRight(fn func(R)) Either[L, R] {
	if e.IsRight() {
		fn(e.right)
	}
	return e
}

func (e *either[L, R]) IfRightAsPtr(fn func(*R)) Either[L, R] {
	if e.IsRight() {
		fn(&e.right)
	}
	return e
}

func (e *either[L, R]) Right() Option[R] {
	if e.IsRight() {
		return Some(e.right)
	} else {
		return None[R]()
	}
}

func (e *either[L, R]) RightAsPtr() Option[*R] {
	if e.IsRight() {
		return Some(&e.right)
	} else {
		return None[*R]()
	}
}

func (e *either[L, R]) Left() Option[L] {
	if e.IsLeft() {
		return Some(e.left)
	} else {
		return None[L]()
	}
}

func (e *either[L, R]) LeftAsPtr() Option[*L] {
	if e.IsLeft() {
		return Some(&e.left)
	} else {
		return None[*L]()
	}
}

func (e *either[L, R]) AsTuple() (L, R) {
	return e.left, e.right
}

func (e *either[L, R]) AsTuplePtr() (*L, *R) {
	return &e.left, &e.right
}

func (e *either[L, R]) UnwrapLeft() L {
	if e.IsRight() {
		panic("called `Either.UnwrapLeft()` on an `Right` value")
	}
	return e.left
}

func (e *either[L, R]) UnwrapLeftOr(v L) L {
	if e.IsLeft() {
		return e.left
	} else {
		return v
	}
}

func (e *either[L, R]) UnwrapLeftOrFrom(c func() L) L {
	if e.IsLeft() {
		return e.left
	} else {
		return c()
	}
}

func (e *either[L, R]) UnwrapLeftAsPtr() *L {
	if e.IsRight() {
		panic("called `Either.UnwrapLeftAsPtr()` on an `Right` value")
	}
	return &e.left
}

func (e *either[L, R]) UnwrapLeftAsPtrOr(value *L) *L {
	if e.IsLeft() {
		return &e.left
	} else {
		return value
	}
}

func (e *either[L, R]) UnwrapLeftAsPtrOrFrom(c func() *L) *L {
	if e.IsLeft() {
		return &e.left
	} else {
		return c()
	}
}

func (e *either[L, R]) UnwrapRight() R {
	if e.IsLeft() {
		panic("called `Either.UnwrapRight()` on an `Left` value")
	}

	return e.right
}

func (e *either[L, R]) UnwrapRightOr(v R) R {
	if e.IsRight() {
		return e.right
	} else {
		return v
	}
}

func (e *either[L, R]) UnwrapRightOrFrom(c func() R) R {
	if e.IsRight() {
		return e.right
	} else {
		return c()
	}
}

func (e *either[L, R]) UnwrapRightAsPtr() *R {
	if e.IsLeft() {
		panic("called `Either.UnwrapRightAsPtr()` on an `Left` value")
	}
	return &e.right
}

func (e *either[L, R]) UnwrapRightAsPtrOr(value *R) *R {
	if e.IsRight() {
		return &e.right
	} else {
		return value
	}
}

func (e *either[L, R]) UnwrapRightAsPtrOrFrom(c func() *R) *R {
	if e.IsRight() {
		return &e.right
	} else {
		return c()
	}
}
