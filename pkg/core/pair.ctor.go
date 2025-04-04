package core

// This constructor creates an instance which implements Pair[F, S].
// fist is 1st value of the pair.
// second is 2nd value of the pair.
func PairFrom[F any, S any](first F, second S) Pair[F, S] {
	return &pair[F, S]{f: first, s: second}
}

// This constructor creates a new instance which implements Pair[F2, S2].
// pair is a pair which is used to create a new Pair[F2, S2].
// callback is a function which creates 1st of F2 and 2nd of S2 which are used as values of Pair[F2, S2].
func MapPair[F1 any, S1 any, F2 any, S2 any](pair Pair[F1, S1], callback func(f F1, s S1) (F2, S2)) Pair[F2, S2] {
	return PairFrom(callback(pair.First(), pair.Second()))
}

// This constructor creates a new instance which implements Pair[S2, S2].
// pair is a pair which is used to create a new Pair[F2, S2].
// callback is a function which creates a new Pair[F2, S2].
func MapPairFrom[F1 any, S1 any, F2 any, S2 any](pair Pair[F1, S1], callback func(f F1, s S1) Pair[F2, S2]) Pair[F2, S2] {
	return callback(pair.First(), pair.Second())
}
