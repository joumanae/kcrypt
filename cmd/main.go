package main

import (
	"flag"
	"fmt"

	"ccipher"
)

func main() {
	cipherText := flag.Bool("ct", false, "What needs to be encrypted")
	decipherText := flag.Bool("dt", false, "What needs to be decrypted")
	message := flag.String("m", " ", "This is the message to cipher or decipher")
	key := flag.Int("k", 1, "This is the key to cipher or decipher")
	flag.Parse()

	if *cipherText {
		c := ccipher.New(*key)
		ciphertext := c.Encipher(*message)
		fmt.Println(ciphertext)
	} else if *decipherText {
		c := ccipher.New(*key)
		deciphered := c.Encipher(*message)
		fmt.Println(deciphered)
	} else {
		fmt.Println("Please enter a message to cipher or decipher")
	}
}
