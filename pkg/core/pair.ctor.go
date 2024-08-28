package core

func PairFrom[F any, S any](f F, s S) Pair[F, S] {
	return &pair[F, S]{f: f, s: s}
}
