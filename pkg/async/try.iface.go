package async

import (
	"github.com/comet11x/go-fpl/pkg/core"
)

type Try[T any] interface {
	IsSuccess() Promise[bool]
	IsFailure() Promise[bool]
	IfSuccess(func(value T))
	IfFailure(func(value any))
	Success() Promise[core.Option[T]]
	Failure() Promise[core.Option[any]]
	AsResult(errorFactory ...func(any) error) Promise[core.Result[T]]
}
