package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"log"
)

const (
	aesKey = "123das88d1mn2fgh"
	aesIv  = "ab123ab13e345678"
)

// 加密
func AesEny(plaintext []byte) string {
	var (
		block cipher.Block
		err   error
	)
	//创建aes
	if block, err = aes.NewCipher([]byte(aesKey)); err != nil {
		log.Fatal(err)
	}
	//创建ctr
	stream := cipher.NewCTR(block, []byte(aesIv))
	//加密, src,dst 可以为同一个内存地址
	stream.XORKeyStream(plaintext, plaintext)

	return hex.EncodeToString(plaintext)
}

// 解密
func AesDec(str string) string {
	var (
		block cipher.Block
		err   error
	)
	ciptext, err := hex.DecodeString(str)
	if err != nil {
		println("===")
	}
	//创建aes
	if block, err = aes.NewCipher([]byte(aesKey)); err != nil {
		log.Fatal(err)
	}
	//创建ctr
	stream := cipher.NewCTR(block, []byte(aesIv))
	stream.XORKeyStream(ciptext, ciptext)
	return string(ciptext)
}

// 解密
func AesDec2(ciptext []byte) string {
	//对密文再进行一次按位异或就可以得到明文
	//例如：3的二进制是0011和8的二进制1000按位异或(相同为0,不同为1)后得到1011，
	//对1011和8的二进制1000再进行按位异或得到0011即是3
	return AesEny(ciptext)
}
