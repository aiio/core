package utils

import (
	"fmt"
	"testing"
)

func TestAesEncrypt(t *testing.T) {
	orig := "hello world"
	key := "whats.framework.whats.framework."
	fmt.Println("orig：", orig)
	encryptCode := AesEncrypt(orig, key)
	fmt.Println("encryptCode：", encryptCode)
	decryptCode := AesDecrypt(encryptCode, key)
	fmt.Println("decryptCode：", decryptCode)
}
