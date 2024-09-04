package async

import (
	"github.com/comet11x/go-fpl/pkg/core"
)

type Try[T any] interface {
	IsSuccess() Promise[bool]
	IsFailure() Promise[bool]
	IfSuccess(func(value T))
	IfFailure(func(value any))
	Success() Promise[T]
	Failure() Promise[any]
	AsResult(errorFactory ...func(any) error) Promise[core.Result[T]]
}
