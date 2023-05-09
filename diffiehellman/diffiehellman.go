package diffiehellman

import (
	"bufio"
	"flag"
	"fmt"
	"math/big"
	"math/rand"
	"net"
	"net/http"
	"time"

	"github.com/go-errors/errors"
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

func ClientAlice() {
	http.Post("http://localhost:8080", "json", nil)

}

func ClientBob() {
	http.Post("http://localhost:8080", "json", nil)
}

func Server(w http.ResponseWriter, r *http.Request) {
	// get the public key from the client
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		errors.Wrap(err, 1)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			errors.Wrap(err, 1)
		}
		go handleConnection(conn)

	}

}

func handleConnection(conn net.Conn) {
	// handle the connection
	for {
		userInput, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			errors.Wrap(err, 1)
		}
		fmt.Println(userInput)
	}
}

func CalculatePrivateKey(publicNumber *big.Int, secret int, modulus int) *big.Int {
	b := keyExchange.base
	pn := CalculatePublicNumber(b, secret, modulus)
	p := Power(int(pn.Int64()), secret)
	p.Mod(p, big.NewInt(int64(modulus)))
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
	PNumberA := CalculatePublicNumber(keyExchange.base, keyExchange.alice, keyExchange.modulus)
	fmt.Println(CalculatePrivateKey(PNumberA, keyExchange.alice, keyExchange.modulus))
	PNumberB := CalculatePublicNumber(keyExchange.base, keyExchange.bob, keyExchange.modulus)
	fmt.Println(CalculatePrivateKey(PNumberB, keyExchange.bob, keyExchange.modulus))

	return 0
}
