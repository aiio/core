package utils

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
)

// Nanoid 生成21位Nanoid
func Nanoid() string {
	id, _ := gonanoid.New()
	return id
}
