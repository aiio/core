package utils

import (
	"strconv"
)

func Per(f float64) string {
	return strconv.FormatFloat((f)*100, 'f', 2, 64) + "%"
}

func FloatToString(f float64, num int) string {
	return strconv.FormatFloat(f, 'f', num, 64)
}
