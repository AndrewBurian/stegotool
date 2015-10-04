package main

import (
	"crypto/rc4"
	"github.com/andrewburian/stegoimg"
	"fmt"
	"io"
)

func read_stego(img io.Reader, secret string, out io.Writer) {

	// see if encryption is used or not
	var encrypt bool = false
	var crypt *rc4.Cipher

	if secret != "" {
		fmt.Println("Encrypting data with AES")
		encrypt = true
		var err error
		crypt, err = rc4.NewCipher([]byte(secret))
		if err != nil {
			panic(err)
		}

		// when the function is finished, zero the keyspace
		defer crypt.Reset()
	}

	// get a buffer to read data from the image
	buf := make([]byte, 128)

	// create the stego img reader to read the data
	stegoReader, err := stegoimg.NewStegoImgReader(img)
	if err != nil {
		panic(err)
	}

	// read and decrypt the data
	for {

		// read the new data block
		n, err := stegoReader.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}

		// break if no data was read
		if n == 0 {
			break
		}

		// optionally decrypt
		if encrypt {
			crypt.XORKeyStream(buf[:n], buf[:n])
		}

		// write to the stego writer
		_, writeErr := out.Write(buf[:n])
		if writeErr != nil {
			panic(writeErr)
		}

		// break if that's the end of the data
		if err == io.EOF {
			break
		}
	}

	return
}
