package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
)

func ProofOfWork(previousProof int, difficulty int) int {
	newProof := 1
	target := strings.Repeat("0", difficulty)

	for {
		hash := sha256.Sum256([]byte(strconv.Itoa(newProof*newProof - previousProof*previousProof)))
		hashString := hex.EncodeToString(hash[:])

		if hashString[:difficulty] == target {
			return newProof
		}
		newProof++
	}
}