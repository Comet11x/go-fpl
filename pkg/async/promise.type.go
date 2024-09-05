package async

import "github.com/comet11x/go-fpl/pkg/core"

type Resolve[T any] func(T)
type Reject func(any)
type args[T any] core.Pair[Resolve[T], Reject]