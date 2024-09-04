package async

import "github.com/comet11x/go-fpl/pkg/core"

type ResolveHandler[T any] func(T)
type RejectedHandler func(any)
type FinallyHandler func()
type Args[T any] core.Pair[ResolveHandler[T], RejectedHandler]
