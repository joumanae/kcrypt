package diffiehellman

import (
	"flag"
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"time"
)

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

func Power(base *big.Int, x int) *big.Int {
	result := big.NewInt(1)
	for i := 0; i < x; i++ {
		result.Mul(result, base)
	}
	return result
}

func ParseBigInt(s string) (*big.Int, bool) {
	n := new(big.Int)
	return n.SetString(s, 10)
}

// A=g^a mod p
func PublicKey(base int, modulus int) *big.Int {
	secret := GenerateRandomNumber()
	p := Power(big.NewInt(int64(base)), secret)
	p.Mod(p, big.NewInt(int64(modulus))) // p = p % modulus
	return p
}

func SharedKey(publicKey *big.Int, secret int, modulus int) *big.Int {

	p := Power(publicKey, secret)
	p = p.Mod(p, big.NewInt(int64(modulus)))
	// write key into file
	return p
}

func Main() int {
	// var b bob
	// var a alice
	mod := flag.Int("m", 0, "mod")   // 0 means use a random prime
	base := flag.Int("b", 0, "base") // 0 means use a random base
	secret := flag.Int("s", 0, "This is a randomly generated secret number")
	pubKey := flag.String("k", "", "This is a public key")

	flag.Parse()

	fmt.Println("Alice enters the chat")
	if len(*pubKey) == 0 {
		// take out the middle parameter
		// encode decode base64 and then calculate shared key
		pn1 := PublicKey(*base, *mod)
		fmt.Printf("This is your public key: %s", pn1)
	} else {
		pk, ok := ParseBigInt(*pubKey)
		if !ok {
			fmt.Println("Your public key is not valid")
			os.Exit(1)
		}
		sk := SharedKey(pk, *secret, *mod)
		fmt.Printf("This is your shared key %s", sk)
	}
	//
	//A dhkeygen -b 18 -m 11
	// A your secret is 25
	//A This is your public key abcdf1234
	//B dhkeygen -b 18 -m 11
	// B your secret is 35
	//B This is your public key gh567890
	//B dhkeygen -pubk= abcdf1234 -s 35 -m 11
	//B This is your shared key X
	//A dhkeygen -pubk= gh567890 -s 25 -m 11
	//A This is your shared key X

	// This is your private key X
	// f1, err := os.Create("alice.priv")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// os.WriteFile("alice.priv", []byte("Here's alice's private key"+" "+pn1.String()), 0644)

	// fmt.Println("Bob enters the chat")
	// pn2 := CalculatePublicKey(keyExchange.base, keyExchange.alice, keyExchange.modulus)
	// CalculatePrivateKey(pn2, keyExchange.bob, keyExchange.modulus)
	// f2, err := os.Create("bob.priv")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// os.WriteFile("bob.priv", []byte("Here's Bob's private key"+" "+pn2.String()), 0644)

	return 0
}
