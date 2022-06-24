package utils

import (
	"fmt"
	"testing"
)

func TestNanoid(t *testing.T) {
	for i := 0; i < 10; i++ {
		id := Nanoid()
		fmt.Printf("Generated id: %s \n", id)
	}
}
