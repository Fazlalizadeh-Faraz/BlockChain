package blockchain

import (
		"testing"
	"fmt"
	"encoding/hex"
		"github.com/stretchr/testify/assert"
)




func TestDiffLevel2(t *testing.T) {


	fmt.Println("~~~~~~~~~~~ testing difficulty: 2 ~~~~~~~~~~~~~~~ ")
	fmt.Println("")
	fmt.Println("\t pre set proof value")


	b0 := Initial(2)
	b0.SetProof(242278)
	fmt.Println(b0.Proof, hex.EncodeToString(b0.Hash))
	b1 := b0.Next("this is an interesting message")
	b1.SetProof(41401)
	fmt.Println(b1.Proof, hex.EncodeToString(b1.Hash))
	b2 := b1.Next("this is not interesting")
	b2.SetProof(195955)
	fmt.Println(b2.Proof, hex.EncodeToString(b2.Hash))



	// ~~~~~~~~~~~~~~~ check the prof ~~~~~~~~~~~~~
	assert.Equal(t, 242278 , int(b0.Proof))
	assert.Equal(t, 41401 , int(b1.Proof))
	assert.Equal(t, 195955 , int(b2.Proof))


	// ~~~~~~~~~~~~ check Hash Value ~~~~~~~~~~~~~~~~~~~~~~
	assert.Equal(t, "29528aaf90e167b2dc248587718caab237a81fd25619a5b18be4986f75f30000" , hex.EncodeToString(b0.Hash))
	assert.Equal(t, "d558f4b9a0b5df021a98066efa5779758cd02f721ebcd8112872265799580000" , hex.EncodeToString(b1.Hash))
	assert.Equal(t, "b611d6fc52964c1c89c717731807c5785ca6bf50d0922b0fc2fdf283757c0000" , hex.EncodeToString(b2.Hash))


	fmt.Println("\t using mine")

	b00 := Initial(2)
	b00.Mine(1)

	fmt.Println(b00.Proof, hex.EncodeToString(b00.Hash))

	b11 := b00.Next("this is an interesting message")
	b11.Mine(1)
	fmt.Println(b11.Proof, hex.EncodeToString(b11.Hash))
	b22 := b11.Next("this is not interesting")
	b22.Mine(1)
	fmt.Println(b22.Proof, hex.EncodeToString(b22.Hash))


	// ~~~~~~~~~~~~~~~ check the prof ~~~~~~~~~~~~~
	assert.Equal(t, 242278 , int(b00.Proof))
	assert.Equal(t, 41401 , int(b11.Proof))
	assert.Equal(t, 195955 , int(b22.Proof))


	// ~~~~~~~~~~~~ check Hash Value ~~~~~~~~~~~~~~~~~~~~~~
	assert.Equal(t, "29528aaf90e167b2dc248587718caab237a81fd25619a5b18be4986f75f30000" , hex.EncodeToString(b00.Hash))
	assert.Equal(t, "d558f4b9a0b5df021a98066efa5779758cd02f721ebcd8112872265799580000" , hex.EncodeToString(b11.Hash))
	assert.Equal(t, "b611d6fc52964c1c89c717731807c5785ca6bf50d0922b0fc2fdf283757c0000" , hex.EncodeToString(b22.Hash))










}

func  TestDiffLevel3(t *testing.T) {


	fmt.Println("~~~~~~~~~~~ testing difficulty: 3 ~~~~~~~~~~~~~~~ ")
	fmt.Println("")
	fmt.Println("\t pre set proof value")


	b0 := Initial(3)
	b0.SetProof(8816998)
	//fmt.Println(b0.Hash)
	fmt.Println(b0.Proof, hex.EncodeToString(b0.Hash))
	b1 := b0.Next("this is an interesting message")
	b1.SetProof(16634616)
	fmt.Println(b1.Proof, hex.EncodeToString(b1.Hash))
	b2 := b1.Next("this is not interesting")
	b2.SetProof(8148543)
	fmt.Println(b2.Proof, hex.EncodeToString(b2.Hash))


	// ~~~~~~~~~~~~~~~ check the prof ~~~~~~~~~~~~~
	assert.Equal(t, 8816998 , int(b0.Proof))
	assert.Equal(t, 16634616 , int(b1.Proof))
	assert.Equal(t, 8148543 , int(b2.Proof))


	// ~~~~~~~~~~~~ check Hash Value ~~~~~~~~~~~~~~~~~~~~~~
	assert.Equal(t, "b7eeb8d8af2133911e09e5687c7b8da31807579fa3661e7dff6fa260e8000000" , hex.EncodeToString(b0.Hash))
	assert.Equal(t, "196d83601de2eedd2d2dd6e10ae804c24f7d45297f16c293ed2c50cfcd000000" , hex.EncodeToString(b1.Hash))
	assert.Equal(t, "e91c5fdf5154cdaa0b3063fb4f9189beb43fd7e67d1c32e9f7e0aa2cd0000000" , hex.EncodeToString(b2.Hash))







	fmt.Println("\t using mine")

	b00 := Initial(3)
	b00.Mine(1)

	fmt.Println(b00.Proof, hex.EncodeToString(b00.Hash))

	b11 := b00.Next("this is an interesting message")
	b11.Mine(1)
	fmt.Println(b11.Proof, hex.EncodeToString(b11.Hash))
	b22 := b11.Next("this is not interesting")
	b22.Mine(1)
	fmt.Println(b22.Proof, hex.EncodeToString(b22.Hash))

	// ~~~~~~~~~~~~~~~ check the prof ~~~~~~~~~~~~~
	assert.Equal(t, 8816998 , int(b00.Proof))
	assert.Equal(t, 16634616 , int(b11.Proof))
	assert.Equal(t, 8148543 , int(b22.Proof))


	// ~~~~~~~~~~~~ check Hash Value ~~~~~~~~~~~~~~~~~~~~~~
	assert.Equal(t, "b7eeb8d8af2133911e09e5687c7b8da31807579fa3661e7dff6fa260e8000000" , hex.EncodeToString(b00.Hash))
	assert.Equal(t, "196d83601de2eedd2d2dd6e10ae804c24f7d45297f16c293ed2c50cfcd000000" , hex.EncodeToString(b11.Hash))
	assert.Equal(t, "e91c5fdf5154cdaa0b3063fb4f9189beb43fd7e67d1c32e9f7e0aa2cd0000000" , hex.EncodeToString(b22.Hash))





}


func TestBlockchainAddandIsValid2(t *testing.T) {

	fmt.Println("~~~~~~~~~~~~ block Chain test diff 2 ~~~~~~~~~~~~~~~~")


	b0 := Initial(2)
	b0.Mine(1)

	fmt.Println(b0.Proof, hex.EncodeToString(b0.Hash))

	b1 := b0.Next("this is an interesting message")
	b1.Mine(1)
	fmt.Println(b1.Proof, hex.EncodeToString(b1.Hash))
	b2 := b1.Next("this is not interesting")
	b2.Mine(1)
	fmt.Println(b2.Proof, hex.EncodeToString(b2.Hash))



	chain := new(Blockchain)
	chain.Add(b0)
	chain.Add(b1)
	chain.Add(b2)

	fmt.Println("Chain validation: " , chain.IsValid())

}



func TestBlockchainAddandIsValid3(t *testing.T) {

	fmt.Println("~~~~~~~~~~~~ block Chain test diff 3 ~~~~~~~~~~~~~~~~")



	b0 := Initial(3)
	b0.Mine(1)

	fmt.Println(b0.Proof, hex.EncodeToString(b0.Hash))

	b1 := b0.Next("this is an interesting message")
	b1.Mine(1)
	fmt.Println(b1.Proof, hex.EncodeToString(b1.Hash))
	b2 := b1.Next("this is not interesting")
	b2.Mine(1)
	fmt.Println(b2.Proof, hex.EncodeToString(b2.Hash))



	chain := new(Blockchain)
	chain.Add(b0)
	chain.Add(b1)
	chain.Add(b2)

	fmt.Println("Chain validation: ",chain.IsValid())

}