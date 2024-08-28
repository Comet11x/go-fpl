package core

type Pair[F any, S any] interface {
	First() F
	Second() S
	ToTuple() (F, S)
	SwapFirst(v F) F
	SwapFirstFrom(fn func(v F) F) F
	SwapSecond(v S) S
	SwapSecondFrom(fn func(s S) S) S
}
