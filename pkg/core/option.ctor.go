package core

func OptionFrom[T any](value any, ok bool) Option[T] {
	if ok {
		return Some(value.(T))
	} else {
		return None[T]()
	}
}

func Some[T any](value T) Option[T] {
	o := option[T]{t: _SOME, value: value}
	return &o
}

func None[T any]() Option[T] {
	o := option[T]{}
	return &o
}

func MapSome[T any, U any](o Option[T], fn func(v T) U) Option[U] {
	if o.IsNone() {
		return None[U]()
	} else {
		return Some(fn(o.Unwrap()))
	}
}

func MapSomeFrom[T any, U any](o Option[T], fn func(v T) Option[U]) Option[U] {
	if o.IsNone() {
		return None[U]()
	} else {
		return fn(o.Unwrap())
	}
}

func MapNone[T any, U any](o Option[T], fn func() U) Option[U] {
	if o.IsNone() {
		return Some(fn())
	} else {
		return None[U]()
	}
}

func MapNoneFrom[T any, U any](o Option[T], fn func() Option[U]) Option[U] {
	if o.IsNone() {
		return fn()
	} else {
		return None[U]()
	}
}

func MapOkAsOption[T any, E any, U any](r Result[T, E], fn func(T) U) Option[U] {
	return MapSome(r.Ok(), fn)
}

func MapOkAsOptionFrom[T any, E any, U any](r Result[T, E], fn func(T) Option[U]) Option[U] {
	return MapSomeFrom(r.Ok(), fn)
}

func MapErrAsOption[T any, E any, U any](r Result[T, E], fn func() U) Option[U] {
	if r.IsErr() {
		return Some(fn())
	} else {
		return None[U]()
	}
}

func MapErrAsOptionFrom[T any, E any, U any](r Result[T, E], fn func(E) Option[U]) Option[U] {
	if r.IsErr() {
		return fn(r.UnwrapErr())
	} else {
		return None[U]()
	}
}
