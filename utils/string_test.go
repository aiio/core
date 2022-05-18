package utils

import (
	"fmt"
	"testing"
)

func TestRandString(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(RandString(32))
	}
}

func TestRandCodeString(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(RandInt(100))
	}
}
