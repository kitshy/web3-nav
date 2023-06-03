package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestAesEny(t *testing.T) {
	plaintext := "我 爱 你"
	fmt.Println("明文", plaintext)
	ciptext := AesEny([]byte(plaintext))
	fmt.Println("加密", ciptext)
	platext1 := AesDec(ciptext)
	fmt.Println("解密", string(platext1))

	str := []byte("shy000")
	by, _ := bcrypt.GenerateFromPassword(str, bcrypt.DefaultCost)
	fmt.Println("====", string(by))
}
