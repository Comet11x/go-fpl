package async

import (
	"github.com/comet11x/go-fpl/pkg/core"
)

type DeferredCall[T any] interface {
	Await() core.Option[core.Try[T]]
	Cancel() bool
}
