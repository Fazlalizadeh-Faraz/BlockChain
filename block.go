package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"bytes"
)

type Block struct {
	Generation uint64
	Difficulty uint8
	Data       string
	PrevHash   []byte
	Hash       []byte
	Proof      uint64
}

// Create new initial (generation 0) block.
func Initial(difficulty uint8) Block {
	// TODO
	newBlcok := Block {
			0,
			difficulty,
			"",
			make([]byte, 32),
			nil,
			0	}
	return newBlcok
}

// Create new block to follow this block, with provided data.
func (prev_block Block) Next(data string) Block {
	// TODO
	followingBlock := Block {
		prev_block.Generation + 1 ,
		prev_block.Difficulty,
		data,
		prev_block.Hash,
		nil,
		0	}
	return followingBlock
}

// Calculate the block's hash.
func (blk Block) CalcHash() []byte {
	// TODO
	data := fmt.Sprintf("%v:%v:%v:%v:%v",
		hex.EncodeToString(blk.PrevHash),
		blk.Generation,
		blk.Difficulty,
		blk.Data,
		blk.Proof )


	hashVal := sha256.New()
	hashVal.Write([]byte(data))
	hashed:= hashVal.Sum(nil)

	//testing
	//mdStr := hex.EncodeToString(hashed)
	//fmt.Println(mdStr)



	return hashed
}

// Is this block's hash valid?
func (blk Block) ValidHash() bool {
	// TODO


	testDiff := make([]byte,blk.Difficulty)

	last :=blk.Hash[len(blk.Hash)-int(blk.Difficulty):]


	if bytes.Equal(testDiff,last){
		return true
	}
	return false


}

// Set the proof-of-work and calculate the block's "true" hash.
func (blk *Block) SetProof(proof uint64) {
	blk.Proof = proof
	blk.Hash = blk.CalcHash()

}
