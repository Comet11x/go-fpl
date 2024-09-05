package async

import (
	"github.com/comet11x/go-fpl/pkg/core"
)

type Awaiter[T any] interface {
	Await() core.Either[T, any]
}

