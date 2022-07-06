package guid

import (
	"time"

	"github.com/sony/sonyflake"
)

var snowflake *sonyflake.Sonyflake

func init() {
	snowflake = sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime:      time.Date(2022, 6, 6, 0, 0, 0, 0, time.UTC),
		MachineID:      nil,
		CheckMachineID: nil,
	})
}

// NextID 生成唯一ID
func NextID() uint64 {
	id, _ := snowflake.NextID()
	return id
}
