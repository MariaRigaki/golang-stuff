package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"fmt"
)

func decryptCBC(ciphertext, key []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte("None"), err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		return []byte("None"), errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// CBC mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		return []byte("None"), errors.New("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(ciphertext, ciphertext)

	return ciphertext, nil
}

func decryptCTR(ciphertext, key []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte("None"), err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		return []byte("None"), errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	plaintext := make([]byte, len(ciphertext))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(plaintext, ciphertext)

	return plaintext, nil
}

func main() {
	// Decode CBC
	c1 := "4ca00ff4c898d61e1edbf1800618fb2828a226d160dad07883d04e008a7897ee2e4b7465d5290d0c0e6c6822236e1daafb94ffe0c5da05d9476be028ad7c1d81"
	ciphertext1, _ := hex.DecodeString(c1)

	c2 := "5b68629feb8606f9a6667670b75b38a5b4832d0f26e1ab7da33249de7d4afc48e713ac646ace36e872ad5fb8a512428a6e21364b0c374df45503473c5242a253"
	ciphertext2, _ := hex.DecodeString(c2)
	k := "140b41b22a29beb4061bda66b6747e14"
	key, _ := hex.DecodeString(k)

	result, err := decryptCBC(ciphertext1, key)
	if err != nil {
		fmt.Println(err)
	}
	// Remove padding bytes (PKCS#7) before printing the result
	n := int(result[len(result)-1])
	fmt.Println(string(result[:len(result)-n]))

	result, err = decryptCBC(ciphertext2, key)
	if err != nil {
		fmt.Println(err)
	}
	// Remove padding bytes (PKCS#7) before printing the result
	n = int(result[len(result)-1])
	fmt.Println(string(result[:len(result)-n]))

	c3 := "69dda8455c7dd4254bf353b773304eec0ec7702330098ce7f7520d1cbbb20fc388d1b0adb5054dbd7370849dbf0b88d393f252e764f1f5f7ad97ef79d59ce29f5f51eeca32eabedd9afa9329"
	ciphertext3, _ := hex.DecodeString(c3)
	c4 := "770b80259ec33beb2561358a9f2dc617e46218c0a53cbeca695ae45faa8952aa0e311bde9d4e01726d3184c34451"
	ciphertext4, _ := hex.DecodeString(c4)
	k2 := "36f18357be4dbd77f050515c73fcf9f2"
	key2, _ := hex.DecodeString(k2)

	result, err = decryptCTR(ciphertext3, key2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(result))

	result, err = decryptCTR(ciphertext4, key2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(result))
}
