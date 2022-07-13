package guid

import (
	"time"
)

var SNOYFLAKE = "snoyflake"
var IDGEN = "idgen"

func Select(engine string, startTime time.Time) IdGen {
	switch engine {
	case IDGEN:
		return NewID(startTime)
	case SNOYFLAKE:
		return NewSf(startTime)
	}
	panic("engine does not support")
}
