package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"fmt"
	"time"
)

type Block struct {
	Timestamp		int64
	Data 			[]byte
	PrevBlockHash 	[]byte
	Hash 			[]byte
	Nonce         	int
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		log.Fatalf("\n Error in Searialize. Error: %s", err)
	}
	return result.Bytes()
}

func DeserializeBlock(d []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		fmt.Printf("\n Error occurred on decode: %s", err)
	}
	return &block

}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
