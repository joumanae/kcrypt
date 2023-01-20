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
	fmt.Println(ccypher.CipherText(cipherText))
}
