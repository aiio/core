package guid

type IdGen interface {
	NextID() uint64
}
