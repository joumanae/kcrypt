package main

import (
	"ccypher"
	"flag"
	"fmt"
)

func main() {
	var cipherText, decipherText bool
	var message string
	var key int
	flag.BoolVar(&cipherText, "ct", false, "What needs to be encrypted")
	flag.BoolVar(&decipherText, "dt", false, "What needs to be decrypted")
	flag.StringVar(&message, "m", " ", "This is the message to cipher or decipher")
	flag.IntVar(&key, "k", 1, "This is the key to cipher or decipher")
	flag.Parse()

	if cipherText {
		c := ccypher.Cipher{Key: key}
		fmt.Println(c.Encipher(message))
	}
	if decipherText {
		d := ccypher.Cipher{Key: -key}
		fmt.Println(d.Decipher(message))
	}
}
