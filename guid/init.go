package guid

import (
	"time"

	"github.com/aiio/core/config"
)

var gen IdGen
var workerID uint16

func init() {
	workerID = uint16(config.GetDefaultEnvToInt("IG_WORKER_ID", 1))
	gen = Select(config.GetDefaultEnv("IG_DEFAULT_ENGINE", "idgen"),
		time.Date(2022, 6, 6, 0, 0, 0, 0, time.UTC))
}

// NextID 生成唯一ID
func NextID() uint64 {
	id := gen.NextID()
	return id
}
