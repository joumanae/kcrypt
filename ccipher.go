package ccipher

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Cipher struct {
	Key int
}

func (c *Cipher) Encipher(message string) string {
	return Transform(message, c.Key)
}

func (c *Cipher) Decipher(message string) string {
	return Transform(message, -c.Key)
}

func Transform(message string, key int) string {
	var b bytes.Buffer
	for _, r := range message {
		r = ShiftRune(r, key)
		b.WriteRune(r)
	}
	return b.String()
}

func FrequencyAnalysis(message string) map[string]int {
	freq := make(map[string]int)
	maxN := 0
	mostFrequentLetter := ""
	for _, r := range message {
		freq[string(r)]++
	}
	for n := range freq {

		if freq[n] > maxN && strings.ContainsAny(n, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
			maxN = freq[n]
			mostFrequentLetter = n
			// if 2 or more letters have the same frequency, the first one will be chosen
		}
	}
	fmt.Printf("The highest frequency is %v, the letter is %v.\n", maxN, mostFrequentLetter)
	return freq
}

func ShiftRune(r rune, Shift int) rune {
	if !unicode.IsLetter(r) {
		return r
	}

	r = unicode.ToUpper(r)

	if r >= 'A' && r <= 'Z' {
		r += rune(Shift)
		if r > 'Z' {
			r -= 26
		}
	}
	return r
}

func New(key int) *Cipher {
	return &Cipher{
		Key: key,
	}
}

const DefaultKey = 13

func Main() int {
	decipherMode := flag.Bool("d", false, "decipher mode")
	decipherwithoutkey := flag.Bool("nk", false, "decipher without key mode")
	key := flag.Int("k", DefaultKey, "the key to encipher/decipher with")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s -k KEY PATH\nEnciphers a given file with a given key\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if len(flag.Args()) < 1 {
		flag.Usage()
		return 1
	}
	message, err := os.ReadFile(flag.Args()[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	cipher := New(*key)
	var output string
	if *decipherMode {
		output = cipher.Decipher(string(message))
	} else {
		output = cipher.Encipher(string(message))
	}
	if *decipherwithoutkey {
		FrequencyAnalysis(string(message))
	}
	fmt.Println(output)
	return 0
}
