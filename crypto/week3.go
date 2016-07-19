package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	//f, err := os.Open("6.2.birthday.mp4_download")
	f, err := os.Open("6.1.intro.mp4_download")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		fmt.Println(err)
	}

	// calculate the data size
	var size int64 = info.Size()
	data := make([]byte, size)

	// read into buffer
	buffer := bufio.NewReader(f)
	_, err = buffer.Read(data)

	offset := len(data) % 1024

	end := len(data)

	var hash_bytes []byte
	var t []byte

	for start := len(data) - offset; start >= 0; {

		t = append(t, data[start:end]...)
		t = append(t, hash_bytes...)

		hasher := sha256.New()
		hasher.Write(t)
		hash_bytes = hasher.Sum(nil)

		end = start
		start -= 1024
		t = []byte{}
	}

	fmt.Println("Result:", hex.EncodeToString(hash_bytes))
}
