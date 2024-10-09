package core

func PairFrom[F any, S any](first F, second S) Pair[F, S] {
	return &pair[F, S]{f: first, s: second}
}

func MapPair[F1 any, S1 any, F2 any, S2 any](pair Pair[F1, S1], callback func(f F1, s S1) (F2, S2)) Pair[F2, S2] {
	return PairFrom(callback(pair.First(), pair.Second()))
}

func MapPairFrom[F1 any, S1 any, F2 any, S2 any](pair Pair[F1, S1], callback func(f F1, s S1) Pair[F2, S2]) Pair[F2, S2] {
	return callback(pair.First(), pair.Second())
}
