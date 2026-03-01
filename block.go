package blockchain

import "time"

type Block struct {
	Index        int
	Timestamp    int64
	Proof        int
	PreviousHash string
}

func NewBlock(index int, proof int, prevHash string) Block {
	return Block{
		Index:        index,
		Timestamp:    time.Now().Unix(),
		Proof:        proof,
		PreviousHash: prevHash,
	}
}