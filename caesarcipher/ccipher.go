package caesarcipher

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"sort"
	"unicode"
)

type Cipher struct {
	Key rune
}

func (c *Cipher) Encipher(message string) string {
	return Transform(message, c.Key)
}

func (c *Cipher) Decipher(message string) string {
	return Transform(message, -c.Key)
}

func Transform(message string, key rune) string {
	var b bytes.Buffer
	for _, r := range message {
		r = ShiftRune(r, key)
		b.WriteRune(r)
	}
	return b.String()
}

func FrequencyAnalysis(message string) (map[rune]int, int) {
	fmt.Println("Running the frequency analysis")
	freq := make(map[rune]int)
	maxN := 0
	for _, r := range message {
		freq[r]++
	}
	for n := range freq {

		if freq[n] > maxN && unicode.IsLetter(n) {
			maxN = freq[n]
		}
	}
	return freq, maxN
}

func FindKey(mostFrequentLetters []string, message string) int {
	fmt.Println("finding the key")
	freq, maxN := FrequencyAnalysis(message)
	for n := range freq {
		if freq[n] == maxN {
			mostFrequentLetters = append(mostFrequentLetters, string(n))
		}
	}
	sort.Strings(mostFrequentLetters)
	mostFrequentLetter := mostFrequentLetters[0]
	key := int(unicode.ToUpper(rune(mostFrequentLetter[0]))) - int('E')
	fmt.Println("Most frequent letter is", mostFrequentLetter, "and the key is", key)
	return key
}

func (c *Cipher) DecipherWithoutKey(message string) string {
	fmt.Println("Running Decipher without key")
	key := FindKey([]string{}, message)
	return Transform(message, rune(-key))
}

func ShiftRune(r rune, Shift rune) rune {
	if !unicode.IsLetter(r) {
		return r
	}

	r = unicode.ToUpper(r)

	if r >= 'A' && r <= 'Z' {
		r += (Shift - 64) % 64
		fmt.Println(Shift, r)
		if r > 'Z' {
			r -= 26
		}
		if r < 'A' {
			r += 26
		}
	}
	return r
}

func New(key rune) *Cipher {
	return &Cipher{
		Key: key,
	}
}

const DefaultKey = "13"

func Main() int {
	decipherMode := flag.Bool("d", false, "decipher mode")
	decipherwithoutkey := flag.Bool("nk", false, "decipher without key mode")
	key := flag.String("k", DefaultKey, "the key to encipher/decipher with")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s -k KEY PATH\nEnciphers a given file with a given key\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	fmt.Println(*decipherMode, *decipherwithoutkey, *key)

	if len(flag.Args()) < 1 {
		flag.Usage()
		return 1
	}
	message, err := os.ReadFile(flag.Args()[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	cipher := New(rune((*key)[0]))

	var output string
	if *decipherMode {
		output = cipher.Decipher(string(message))
	} else {
		output = cipher.Encipher(string(message))
	}
	if *decipherwithoutkey {
		output = cipher.DecipherWithoutKey(string(message))
	}
	fmt.Println(output)
	return 0
}
