package main

import (
	"ccypher"
	"flag"
	"fmt"
)

func main() {
	var cipherText, decipherText bool
	var message string
	flag.BoolVar(&cipherText, "ct", false, "What needs to be encrypted")
	flag.BoolVar(&decipherText, "dt", false, "What needs to be decrypted")
	flag.Parse()
	c := ccypher.Cipher{Key: 1}
	d := ccypher.Cipher{Key: -1}
	fmt.Println(c.Encipher(message))
	fmt.Println(d.Decipher(message))
}
