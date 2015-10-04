package main

import (
	"os"
	"fmt"
	"flag"
	"io"
)

func main() {

	// Set up program aruments
	write := flag.Bool("write", false, "Set mode to creating a new stego image")
	read := flag.Bool("read", false, "Set mode to reading a stego image")
	imgfile := flag.String("img", "", "The image file to use as the source. Can be png, jpg, gif, or bmp")
	outfile := flag.String("output", "stdout", "Output of the operation. Either a stego img, or extracted data")
	datafile := flag.String("data", "stdin", "Data file to embedd")
	secret := flag.String("secret", "", "RC4 encryption key for stored data")
	help := flag.Bool("help", false, "Print usage")
	flag.Parse()

	// if help, print usage and exit
	if *help {
		flag.Usage()
		return
	}

	// check argument sanity
	if !*read && !*write {
		fmt.Println("No mode selected")
		flag.Usage()
		return
	}

	if *read && *write {
		fmt.Println("Conflicting modes selected")
		flag.Usage()
		return
	}

	if *imgfile == "" {
		fmt.Println("Image file required")
		flag.Usage()
		return
	}

	// variables
	var img, data io.Reader
	var out io.Writer

	// open the source image
	img, err := os.Open(*imgfile)
	if err != nil {
		panic(err)
	}

	// open the output
	if *outfile == "stdout" {
		out = os.Stdout
	} else {
		out, err := os.Open(*outfile)
		if err != nil {
			panic(err)
		}
		defer out.Close()
	}

	// open the datafile
	if *datafile == "stdin" {
		data = os.Stdin
	} else {
		data, err := os.Open(*datafile)
		if err != nil {
			panic(err)
		}
		defer data.Close()
	}

	// run the correct mode
	if *write {
		write_stego(img, data, *secret, out)
	}
	if *read {
		read_stego(img, *secret, out)
	}

	// done
	return

}


func read_stego(img io.Reader, secret string, out io.Writer) {

	return
}
