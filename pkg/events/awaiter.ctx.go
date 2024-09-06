package events

import (
	"errors"
	"reflect"
	"strings"

	"github.com/comet11x/go-fpl/pkg/core"
	"github.com/comet11x/go-fpl/pkg/types"
)

// # TryAwaiterFrom constructor
func TryAwaiterFrom(ee EventEmitter) core.Result[types.Awaiter[types.Void]] {
	if strings.Contains(reflect.TypeOf(ee).String(), "*events.eventEmitter") {
		return core.Ok(ee.(types.Awaiter[types.Void]))
	} else {
		return core.Err[types.Awaiter[types.Void]](errors.New("incompatible types"))
	}
}
