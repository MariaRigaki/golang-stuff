package main

import (
	"crypto/aes"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func generatePadding(numBytes int) []byte {
	padding := make([]byte, aes.BlockSize)

	for i := 0; i < numBytes; i++ {
		padding[aes.BlockSize-i-1] = byte(numBytes)
	}
	return padding
}

func xorSlices(a, b []byte) ([]byte, error) {
	if len(b) != len(a) {
		return []byte{}, errors.New("Uneven lengths")
	}

	result := make([]byte, len(a))
	for i, _ := range a {
		result[i] = a[i] ^ b[i]
	}

	return result, nil
}

func main() {
	URL := "http://crypto-class.appspot.com/po"
	c := "f20bdba6ff29eed7b046d1df9fb7000058b1ffb4210a580f748b4ac714c001bd4a61044426fb515dad3f21f18aa577c0bdf302936266926ff37dbf7035d5eeb4"

	client := &http.Client{}
	req, _ := http.NewRequest("GET", URL, nil)
	q := req.URL.Query()
	q.Add("er", c)

	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL.String())
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println(resp)
	// fmt.Println(len(c))

	ciphertext, _ := hex.DecodeString(c)
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// fmt.Println(len(ciphertext) / aes.BlockSize)
	// fmt.Println(len(iv))

	//fmt.Println(generatePadding(2))

	c0 := ciphertext[:aes.BlockSize]
	//c1 := ciphertext[aes.BlockSize : 2*aes.BlockSize]
	//c3 := ciphertext[2*aes.BlockSize:]

	newIV := make([]byte, aes.BlockSize)
	cNew := make([]byte, len(ciphertext))

	plaintext := make([]byte, len(ciphertext))
	copy(newIV, iv)

	for i := 1; i <= 16; i++ {

		p := generatePadding(i)
		newIV, _ = xorSlices(newIV, p)

		fmt.Println(hex.EncodeToString(newIV))

		for g := 32; g < 122; g++ {

			cNew = []byte{}
			cNew = append(cNew, newIV[:aes.BlockSize-i]...)
			cNew = append(cNew, []byte{newIV[aes.BlockSize-i] ^ byte(g)}...)
			cNew = append(cNew, newIV[aes.BlockSize-i+1:]...)
			cNew = append(cNew, c0...)

			q.Set("er", hex.EncodeToString(cNew))
			req.URL.RawQuery = q.Encode()
			//fmt.Println(req)
			resp, _ = client.Do(req)
			if resp.StatusCode == 404 {
				fmt.Println("Found g:", g, resp.StatusCode)
				//fmt.Println(hex.EncodeToString(cNew))
				//newIV[aes.BlockSize-i] = byte(g) ^ byte(i)
				plaintext[aes.BlockSize-i] = byte(g)
				break
			}
			time.Sleep(150 * time.Millisecond)
		}
		fmt.Println(string(plaintext))
	}

	fmt.Println(string(plaintext))
}
