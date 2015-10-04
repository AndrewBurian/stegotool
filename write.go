package main

import (
	"crypto/rc4"
	"fmt"
	"github.com/andrewburian/stegoimg"
	"io"
)

func write_stego(img, data io.Reader, secret string, out io.Writer) {

	// see if encryption is used or not
	var encrypt bool = false
	var crypt *rc4.Cipher

	if secret != "" {
		encrypt = true
		var err error
		crypt, err = rc4.NewCipher([]byte(secret))
		if err != nil {
			panic(err)
		}

		// when the function is finished, zero the keyspace
		defer crypt.Reset()
	}

	// get a buffer to read data from the data
	buf := make([]byte, 128)

	// create the new stego img writer to encode the data
	stegoWriter, err := stegoimg.NewStegoImgWriter(img, out)
	if err != nil {
		panic(err)
	}
	defer stegoWriter.Close()

	// read and encode the data
	for {

		// read the new data block
		n, err := data.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}

		// break if no data was read
		if n == 0 {
			break
		}

		// optionally encrypt
		if encrypt {
			crypt.XORKeyStream(buf[:n], buf[:n])
		}

		// write to the stego writer
		_, writeErr := stegoWriter.Write(buf[:n])

		// check if the image filled up
		if writeErr == stegoimg.ImageFullError && err != io.EOF {
			fmt.Println("Image full before data finished")
			break
		}

		// break if that's the end of the data
		if err == io.EOF {
			break
		}
	}

	// finished creating the image
	return
}
