package wallet

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func GenerateWallet() {
	pub, priv, _ := ed25519.GenerateKey(rand.Reader)
	fmt.Println("Private:", hex.EncodeToString(priv))
	fmt.Println("Public :", hex.EncodeToString(pub))
}