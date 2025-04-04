package core

type Pair[F any, S any] interface {

	// Returns 1st value
	First() F

	// Returns 2nd value
	Second() S

	// Returns the values of the pair as a tuple
	AsTuple() (F, S)

	// Swaps 1st value and the given value
	// v is a value of F
	SwapFirst(v F) F

	// Swap 1st value and the value which is provided by the given function.
	// fn is a function which provides the value of F.
	SwapFirstFrom(fn func(v F) F) F

	// Swap 2nd value and the given value
	// v is a value of S
	SwapSecond(v S) S

	// Swap 2st value and the value which is provided by the given function.
	// fn is a function which provides the value of S.
	SwapSecondFrom(fn func(s S) S) S
}
