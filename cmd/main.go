package main

import (
	"ccypher"
	"flag"
	"fmt"
)

func main() {
	var cipherText string
	flag.StringVar(&cipherText, "ct", "En Attendant Godot", "What needs to be encrypted")
	flag.Parse()
	c := ccypher.Cipher{Key: 1, PlainText: cipherText}
	fmt.Println(c.CipherText(cipherText))
}
