package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

// 1. choose cipher - des
// 2. choose padding
// 3. choose iv
// 4. choose mode-cbc mode
// 5. encrypt

func desCBCEecrypt(plaintext []byte, key []byte) ([]byte, error) {

	block, err := des.NewCipher(key)

	if err != nil {
		fmt.Println("new block", block, "error:", err)
		return nil, err
	}
	iv := bytes.Repeat([]byte("1"), block.BlockSize())

	blockmode := cipher.NewCBCEncrypter(block, iv)
	paddings := padding(plaintext, block.BlockSize())
	fmt.Println("paddings are :", string(paddings))

	plaintext = append(plaintext, paddings...)
	fmt.Println("plaintext with paddings are :", string(plaintext))

	ciphertext := make([]byte, len(plaintext))
	blockmode.CryptBlocks(ciphertext, []byte(plaintext))

	fmt.Printf("the ciphertext is:%x\n", ciphertext)
	return ciphertext, nil
}

func desCBCDecrypt(ciphertext, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)

	if err != nil {
		fmt.Println("new block", block, "error:", err)
		return nil, err
	}

	iv := bytes.Repeat([]byte("1"), block.BlockSize())
	blockmode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))

	blockmode.CryptBlocks(plaintext, ciphertext)
	fmt.Printf("plain text is : %s\n", string(plaintext))

	return plaintext, nil
}

func padding(plaintext []byte, blocksize int) []byte {
	fmt.Println("length of plaintext is : ", len(string(plaintext)))

	if len(string(plaintext))%blocksize == 0 {
		str := bytes.Repeat([]byte{byte(blocksize)}, blocksize)
		fmt.Println("str is : ", string(str))
		return str
	}
	str2 := bytes.Repeat([]byte{byte(len(string(plaintext)) % blocksize)}, len(string(plaintext))%blocksize)
	fmt.Println("str 1 is :", string(str2))
	return str2
}

func main() {
	plaintext := []byte("123456789123123123")
	key := []byte("12345678")

	ciphertext, _ := desCBCEecrypt(plaintext, key)
	// str := hex.EncodeToString(ciphertext)

	plaintext, _ = desCBCDecrypt(ciphertext, key)

}
