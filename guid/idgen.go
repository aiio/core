package guid

import (
	"time"

	"github.com/yitter/idgenerator-go/idgen"
)

type id struct {
	client *idgen.DefaultIdGenerator
}

// NewID 初始化一个 id gen
func NewID(startTime time.Time) *id {
	return &id{client: idgen.NewDefaultIdGenerator(&idgen.IdGeneratorOptions{
		Method:            1,
		WorkerId:          workerID,
		BaseTime:          startTime.UnixMilli(),
		WorkerIdBitLength: 6,
		SeqBitLength:      6,
		MaxSeqNumber:      0,
		MinSeqNumber:      5,
		TopOverCostCount:  2000,
	})}
}

func (s *id) NextID() uint64 {
	return uint64(s.client.NewLong())
}
