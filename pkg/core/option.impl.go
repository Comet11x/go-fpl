package core

type option[T any] struct {
	t     int
	value T
}

func (o *option[T]) IsSome() bool {
	return o.t == _SOME
}

func (o *option[T]) IsNone() bool {
	return o.t == _NONE
}

func (o *option[T]) IfSome(fn func(v T)) bool {
	cond := o.IsSome()
	if cond {
		fn(o.value)
	}
	return cond
}

func (o *option[T]) IfNone(fn func()) bool {
	cond := o.IsNone()
	if cond {
		fn()
	}
	return cond
}

func (o *option[T]) MapSome(fn func(v T) T) Option[T] {
	if o.IsSome() {
		return Some(fn(o.value))
	} else {
		return o
	}
}

func (o *option[T]) MapSomeFrom(fn func(v T) Option[T]) Option[T] {
	if o.IsSome() {
		return fn(o.value)
	} else {
		return o
	}
}

func (o *option[T]) MapNone(fn func() T) Option[T] {
	if o.IsNone() {
		return Some(fn())
	} else {
		return o
	}
}

func (o *option[T]) MapNoneFrom(fn func() Option[T]) Option[T] {
	if o.IsNone() {
		return fn()
	} else {
		return o
	}
}

func (o *option[T]) UnwrapOr(value T) T {
	if o.IsNone() {
		return value
	} else {
		return o.value
	}
}

func (o *option[T]) UnwrapOrValueFrom(c func() T) T {
	if o.IsNone() {
		return c()
	} else {
		return o.value
	}
}

func (o *option[T]) UnwrapAsPtrOr(value *T) *T {
	if o.IsNone() {
		return value
	} else {
		return &o.value
	}
}

func (o *option[T]) UnwrapAsPtrOrPtrFrom(c func() *T) *T {
	if o.IsNone() {
		return c()
	} else {
		return &o.value
	}
}

func (o *option[T]) Unwrap() T {
	return o.value
}

func (o *option[T]) UnwrapAsPtr() *T {
	return &o.value
}

func (o *option[T]) Swap(value T) T {
	prev := o.value
	o.value = value
	return prev
}

func (o *option[T]) SwapFrom(c func() T) T {
	prev := o.value
	o.value = c()
	return prev
}

func (o *option[T]) SwapAsPtr(value *T) T {
	prev := o.value
	o.value = *value
	return prev
}

func (o *option[T]) SwapAsPtrFrom(c func() *T) T {
	prev := o.value
	cur := c()
	o.value = *cur
	return prev
}

func (o *option[T]) ToTuple() (T, bool) {
	return o.value, o.IsSome()
}

func (o *option[T]) ToTupleAsPtr() (*T, bool) {
	return &o.value, o.IsSome()
}
