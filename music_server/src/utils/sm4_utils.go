package utils

import (
	"bytes"
	"crypto/cipher"
	"encoding/hex"
	"github.com/spf13/viper"
	"github.com/tjfoc/gmsm/sm4"
)

// 明文数据填充
func paddingLastGroup(plainText []byte, blockSize int) []byte {
	//1.计算最后一个分组中明文后需要填充的字节数
	padNum := blockSize - len(plainText)%blockSize
	//2.将字节数转换为byte类型
	char := []byte{byte(padNum)}
	//3.创建切片并初始化
	newPlain := bytes.Repeat(char, padNum)
	//4.将填充数据追加到原始数据后
	newText := append(plainText, newPlain...)

	return newText
}

// 去掉明文后面的填充数据
func unpaddingLastGroup(plainText []byte) []byte {
	//1.拿到切片中的最后一个字节
	length := len(plainText)
	lastChar := plainText[length-1]
	//2.将最后一个数据转换为整数
	number := int(lastChar)
	return plainText[:length-number]
}

func SM4Encrypt(plainText string) string {
	var secret_key = viper.GetString("sm4.secretKey")
	var sm4iv = viper.GetString("sm4.iv")

	block, err := sm4.NewCipher([]byte(secret_key))
	if err != nil {
		panic(err)
	}
	paddData := paddingLastGroup([]byte(plainText), block.BlockSize())
	iv := []byte(sm4iv)
	blokMode := cipher.NewCBCEncrypter(block, iv)
	cipherText := make([]byte, len(paddData))
	blokMode.CryptBlocks(cipherText, paddData)
	// 将加密结果转换为十六进制字符串
	ciphertextstr := hex.EncodeToString(cipherText)
	return ciphertextstr
}

func SM4Dectypt(cipherText string) string {
	bb, _ := hex.DecodeString(cipherText)
	var secret_key = viper.GetString("sm4.secretKey")
	var sm4iv = viper.GetString("sm4.iv")

	block, err := sm4.NewCipher([]byte(secret_key))
	if err != nil {
		panic(err)
	}
	iv := []byte(sm4iv)
	blockMode := cipher.NewCBCDecrypter(block, iv)
	blockMode.CryptBlocks(bb, bb)
	plainText := unpaddingLastGroup(bb)

	return string(plainText)
}
