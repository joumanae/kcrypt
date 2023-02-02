package main

import (
	"ccypher"
	"flag"
	"fmt"
)

func main() {
	var cipherText, decipherText string
	flag.StringVar(&cipherText, "ct", "En Attendant Godot", "What needs to be encrypted")
	flag.StringVar(&decipherText, "dt", "Fo Buubmbof Hpepu", "What needs to be decrypted")
	flag.Parse()
	c := ccypher.Cipher{Key: 1}
	d := ccypher.Cipher{Key: -1}
	fmt.Println(c.Encipher(cipherText))
	fmt.Println(d.Decipher(decipherText))
}
