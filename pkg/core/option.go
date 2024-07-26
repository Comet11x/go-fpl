package core

type Option[T any] struct {
	value *T
}

func CreateOptionFromPtr[T any](ptr *T) Option[T] {
	if ptr == nil {
		return Option[T]{}
	} else {
		return Option[T]{value: ptr}
	}
}

func Some[T any](value T) Option[T] {
	return Option[T]{value: &value}
}

func None[T any]() Option[T] {
	return Option[T]{}
}

func (o *Option[T]) IsSome() bool {
	return o.value != nil
}

func (o *Option[T]) IsNone() bool {
	return o.value == nil
}

func (o *Option[T]) UnwrapOr(value T) *T {
	if o.IsNone() {
		return &value
	} else {
		return o.value
	}
}

func (o *Option[T]) Unwrap() *T {
	return o.value
}

func (o *Option[T]) UnwrapClone() T {
	return *o.value
}

func (o *Option[T]) Swap(value *T) *T {
	ptr := o.value
	o.value = value
	return ptr
}

func MapSome[T any, U any](o Option[T], fn func(v *T) U) Option[U] {
	if o.IsNone() {
		return None[U]()
	} else {
		return Some[U](fn(o.value))
	}
}

func MapNone[T any, U any](o Option[T], fn func() U) Option[U] {
	if o.IsNone() {
		return Some[U](fn())
	} else {
		return None[U]()
	}
}
