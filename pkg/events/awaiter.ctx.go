package events

import (
	"reflect"
	"strings"

	"github.com/comet11x/go-fpl/pkg/core"
	"github.com/comet11x/go-fpl/pkg/types"
)

func Awaiter(ee EventEmitter) core.Option[types.Awaiter[types.Void]] {
	if strings.Contains(reflect.TypeOf(ee).String(), "*events.eventEmitter") {
		return core.Some(ee.(types.Awaiter[types.Void]))
	} else {
		return core.None[types.Awaiter[types.Void]]()
	}
}
