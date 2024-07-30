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

func (o *option[T]) UnwrapAsPtrOr(value *T) *T {
	if o.IsNone() {
		return value
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

func (o *option[T]) SwapFromPtr(value *T) T {
	prev := o.value
	o.value = *value
	return prev
}
