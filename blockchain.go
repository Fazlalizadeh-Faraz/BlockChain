package blockchain

import (
		"fmt"
	"bytes"
)

type Blockchain struct {
	Chain []Block
}

func (chain *Blockchain) Add(blk Block) {
	if !blk.ValidHash() {
		panic("adding block with invalid hash")
	}
	chain.Chain = append(chain.Chain, blk)
	// TODO
}

func (chain Blockchain) IsValid() bool {
	// TODO
	tempBlock := Initial(chain.Chain[0].Difficulty)
	counter := 0

	for i := 0; i < len(chain.Chain[0].PrevHash); i++ {
		if chain.Chain[0].PrevHash[i] != 0 {
			return false
		}
	}

	for i := 1; i < len(chain.Chain); i++ {

		//Check the difficulties to see if they change
		if tempBlock.Difficulty != chain.Chain[i].Difficulty {
			fmt.Println("different difficulty values")
			return false
		}

		//Check the generation for each increment in the chain
		if chain.Chain[counter].Generation+1 != chain.Chain[i].Generation {
			fmt.Println("generation value failed")
			return false
		}

		////Check each value in the hash
		checkValidHash := chain.Chain[counter].ValidHash()
		if !checkValidHash {
			fmt.Println("valid hash failed")
			return false
		}


		if !bytes.Equal(chain.Chain[i].CalcHash() ,chain.Chain[i].Hash) {
			fmt.Println("hash and cal hash are not equal")
			return false
		}


		//Check if the hash values of a block are equal to the previous hash value of the next one
		if !bytes.Equal (chain.Chain[counter].Hash , chain.Chain[i].PrevHash) {
			fmt.Println("prev hash not right")
			return false
		}


		//Increment counter with each i
		counter++
	}

	return true
}