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
	pn := CalculatePublicNumber(keyExchange.base, keyExchange.bob, keyExchange.modulus)
	http.HandleFunc("/alice", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, " Here's Bob's's public number %s", pn.String())
	})

}

func ClientBob() {
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

	return p
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

	channel := make(chan int)
	go func() {
		fmt.Println("Alice enters the chat")
		pn := CalculatePublicNumber(keyExchange.base, keyExchange.bob, keyExchange.modulus)
		CalculatePrivateKey(pn, keyExchange.alice, keyExchange.modulus)
		_, err := os.Create("alice.txt")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		os.WriteFile("alice.txt", []byte("Here's alice's private key"+" "+pn.String()), 0644)
		channel <- 1
	}()
	go func() {
		fmt.Println("Bob enters the chat")
		pn := CalculatePublicNumber(keyExchange.base, keyExchange.alice, keyExchange.modulus)
		CalculatePrivateKey(pn, keyExchange.bob, keyExchange.modulus)
		_, err := os.Create("bob.txt")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		os.WriteFile("bob.txt", []byte("Here's Bob's private key"+" "+pn.String()), 0644)
		channel <- 2
	}()
	<-channel

	Server()
	return 0
}
