package async

import (
	"github.com/comet11x/go-fpl/pkg/core"
)

type Try[T any] interface {
	IsSuccess() Promise[bool]
	IsFailure() Promise[bool]
	Success() Promise[T]
	Failure() Promise[any]
	ToResult(errorFactory ...func(any) error) Promise[core.Result[T]]
}
