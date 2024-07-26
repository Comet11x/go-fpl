package core

type Either[A any, B any] struct {
	left  *A
	right *B
}

func EitherFrom[A any, B any](l *A, r *B) Option[Either[A, B]] {
	if l != nil {
		return Some(Either[A, B]{left: l})
	} else if r != nil {
		return Some(Either[A, B]{right: r})
	} else {
		return None[Either[A, B]]()
	}
}

func EitherFromResult[T any](r Result[T]) Either[T, error] {
	if r.IsOk() {
		return Either[T, error]{left: r.ok}
	} else {
		return Either[T, error]{right: &r.err}
	}
}

func Left[A any, B any](data A) Either[A, B] {
	return Either[A, B]{left: &data}
}

func Right[A any, B any](data B) Either[A, B] {
	return Either[A, B]{right: &data}
}

func (e *Either[A, B]) IsRight() bool {
	return e.left == nil
}

func (e *Either[A, B]) IsLeft() bool {
	return e.right == nil
}

func (e *Either[A, B]) Right() Option[B] {
	return Option[B]{value: e.right}
}

func (e *Either[A, B]) Left() Option[A] {
	return Option[A]{value: e.left}
}

func (e *Either[A, B]) ToTuple() (*A, *B) {
	return e.left, e.right
}

func (e *Either[A, B]) UnwrapLeft() *A {
	return e.left
}

func (e *Either[A, B]) UnwrapLeftClone() A {
	return *e.left
}

func (e *Either[A, B]) UnwrapRight() *B {
	return e.right
}

func (e *Either[A, B]) UnwrapRightClone() B {
	return *e.right
}

func MapRight[A any, B any, R any](e Either[A, B], fn func(ptr *B) R) Either[A, R] {
	if e.IsRight() {
		return Right[A, R](fn(e.right))
	} else {
		return Either[A, R]{left: e.left}
	}
}

func MapLeft[A any, B any, L any](e Either[A, B], fn func(ptr *A) L) Either[L, B] {
	if e.IsLeft() {
		return Left[L, B](fn(e.left))
	} else {
		return Either[L, B]{right: e.right}
	}
}
