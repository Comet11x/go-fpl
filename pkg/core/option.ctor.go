package core

func Some[T any](value T) Option[T] {
	o := option[T]{t: _SOME, value: value}
	return &o
}

func None[T any]() Option[T] {
	o := option[T]{t: _NONE}
	return &o
}

func MapSome[T any, U any](o Option[T], fn func(v T) U) Option[U] {
	if o.IsNone() {
		return None[U]()
	} else {
		return Some[U](fn(o.Unwrap()))
	}
}

func MapNone[T any, U any](o Option[T], fn func() U) Option[U] {
	if o.IsNone() {
		return Some[U](fn())
	} else {
		return None[U]()
	}
}
