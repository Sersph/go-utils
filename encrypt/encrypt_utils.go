package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

const BASE64Table = "IJjkKLMNO567PQX12RVW3YZaDEFGbcdefghiABCHlSTUmnopqrxyz04stuvw89+/"

// 对称加密解密操作
// 原文进去密文出来
// 密文进原文出
func SymmetricEncrypt(source string, key string) []byte {
	if len(key) != 16 {
		panic("key必须是16位密钥！")
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	// 分组
	iv := bytes.Repeat([]byte("1"), block.BlockSize())

	stream := cipher.NewCTR(block, iv)

	sourceByte := []byte(source)

	resultBuffer := make([]byte, len(sourceByte))

	for i, b := range resultBuffer {
		fmt.Println(i, string(b))
	}

	stream.XORKeyStream(resultBuffer, sourceByte)
	return resultBuffer
}
