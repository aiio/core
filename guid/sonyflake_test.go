package guid

import (
	"log"
	"testing"
)

func TestNextID(t *testing.T) {
	for i := 0; i < 100; i++ {
		log.Println(NextID())
	}
}
