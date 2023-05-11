package diffiehellman

import (
	"flag"
	"fmt"
	"math/big"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var keyExchange struct {
	modulus int
	base    int
	alice   int
	bob     int
}

// TODO: add clients and server

func GenerateRandomNumber() int {
	rand.Seed(time.Now().Unix())
	prime := rand.Intn(1000)
	return prime
}

// generate random secret key
func GenerateSecretKey() int {
	rand.Seed(time.Now().Unix())
	secret := rand.Intn(1000)
	return secret
}

func Power(base int, x int) *big.Int {
	result := big.NewInt(1)
	for i := 0; i < x; i++ {
		result.Mul(result, big.NewInt(int64(base)))
	}
	return result
}

// A=g^a mod p
func CalculatePublicNumber(base int, secret int, modulus int) *big.Int {
	p := Power(base, secret)
	p.Mod(p, big.NewInt(int64(modulus))) // p = p % modulus
	return p
}

func Server() {

	ClientAlice()
	ClientBob()
	http.ListenAndServe("localhost:8080", nil)
}

func ClientAlice() {
	fmt.Println("Alice enters the chat")

	pn := CalculatePublicNumber(keyExchange.base, keyExchange.bob, keyExchange.modulus)
	http.HandleFunc("/alice", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, " Here's Bob's's public number %s", pn.String())
	})

}

func ClientBob() {
	fmt.Println("Bob enters the chat")
	pn := CalculatePublicNumber(keyExchange.base, keyExchange.alice, keyExchange.modulus)
	http.HandleFunc("/bob", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, " Here's Alice's public number %s", pn.String())
	})
}

func CalculatePrivateKey(publicNumber *big.Int, secret int, modulus int) *big.Int {
	b := keyExchange.base
	pn := CalculatePublicNumber(b, secret, modulus)
	p := Power(int(pn.Int64()), secret)
	p = p.Mod(p, big.NewInt(int64(modulus)))
	// write key into file
	writeKeytoFile(p)
	return p
}

func writeKeytoFile(key *big.Int) {
	file, err := os.OpenFile("key.txt", os.O_APPEND, 0660)
	if err != nil {
		fmt.Println("Error opening file")
	}
	defer file.Close()

}

func Main() int {
	// var b bob
	// var a alice
	mod := flag.Int("mod", GenerateRandomNumber(), "mod")    // 0 means use a random prime
	base := flag.Int("base", GenerateRandomNumber(), "base") // 0 means use a random base

	flag.Parse()

	keyExchange.modulus = *mod
	keyExchange.base = *base
	keyExchange.alice = GenerateSecretKey()
	keyExchange.bob = GenerateSecretKey()
	PNumberA := CalculatePublicNumber(keyExchange.base, keyExchange.alice, keyExchange.modulus)
	CalculatePrivateKey(PNumberA, keyExchange.alice, keyExchange.modulus)
	PNumberB := CalculatePublicNumber(keyExchange.base, keyExchange.bob, keyExchange.modulus)
	CalculatePrivateKey(PNumberB, keyExchange.bob, keyExchange.modulus)

	Server()
	return 0
}
