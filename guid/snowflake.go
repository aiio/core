package guid

import (
	"time"

	"github.com/sony/sonyflake"
)

type sf struct {
	client *sonyflake.Sonyflake
}

// NewSf 初始化一个 sonyflake
func NewSf(startTime time.Time) *sf {
	return &sf{client: sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: startTime,
		MachineID: func() (uint16, error) {
			return workerID, nil
		},
		CheckMachineID: nil,
	})}
}

func (s *sf) NextID() uint64 {
	id, _ := s.client.NextID()
	return id
}
