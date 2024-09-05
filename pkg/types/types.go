package types

type Void = struct{}

type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type UInt interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Float interface {
	~float32 | ~float64
}

type Complex interface {
	~complex64 | ~complex128
}

type Number interface {
	Int | UInt | Float
}

type Ordered interface {
	Int | UInt | Float | ~rune | ~string
}

type Equal interface {
	Ordered | Complex
}

type Iterable interface {
	~[]any | ~map[any]any | ~string
}
