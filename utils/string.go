package utils

import (
	"bytes"
	"crypto/rand"
	"math/big"
)

// RandString 生成随机字符串
func RandString(size int) string {
	// 过滤相似字符（类似：i, l, 1, L, o, 0, O）
	const char = "abcdefghjkmnpqrstuvwxyzABCDEFGHJKMNPQRSTUVWXYZ23456789"
	var s bytes.Buffer
	for i := 0; i < size; i++ {
		result, _ := rand.Int(rand.Reader, big.NewInt(1000000))
		s.WriteByte(char[result.Int64()%int64(len(char))])
	}
	return s.String()
}

// RandCodeString 生成随机字数字符串
func RandCodeString(size int) string {
	const char = "0123456789"
	var s bytes.Buffer
	for i := 0; i < size; i++ {
		result, _ := rand.Int(rand.Reader, big.NewInt(1000000))
		s.WriteByte(char[result.Int64()%int64(len(char))])
	}
	return s.String()
}

// RandHexString RandHexString
func RandHexString(size int) string {
	const char = "abcdefghjkmnpqrstuvwxyz23456789"
	var s bytes.Buffer
	for i := 0; i < size; i++ {
		result, _ := rand.Int(rand.Reader, big.NewInt(1000000))
		s.WriteByte(char[result.Int64()%int64(len(char))])
	}
	return s.String()
}

func RandInt(max int64) int64 {
	n, _ := rand.Int(rand.Reader, big.NewInt(max))
	return n.Int64()
}
