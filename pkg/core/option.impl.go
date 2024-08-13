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
