package main

import (
	"os"

	dhkeygen "github.com/joumanae/cryptographywithgo/dhkeygen"
)

func main() {
	os.Exit(dhkeygen.Main())
}
