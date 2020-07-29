package aes

import (
	"crypto/aes"
	"crypto/cipher"
)

func padCcu(i, j int) int {
	result := 16 + (-i)%j
	for result > 16 {
		result -= 16
	}
	return result
}

func Encode(data, key,iv []byte) ([]byte, error) {
	padL := padCcu(len(data), 16)
	pad := make([]byte, padL)
	for i := 0; i < padL; i++ {
		pad[i] = byte(padL)
	}
	data = append(data, pad...)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil,err
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	en := make([]byte, len(data))
	mode.CryptBlocks(en, data)
	return en,nil
}

func Decode(en, key,iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	data := make([]byte, len(en))
	mode.CryptBlocks(data, en)
	pad := data[len(data)-1]
	if pad <= 16 && data[len(data)-int(pad)] == pad {
		data = data[:len(data)-int(pad)]
	}
	return data,nil
}