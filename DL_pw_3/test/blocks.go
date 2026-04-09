package test

import (
	blocks "DL_pw_3/blocks"
	"fmt"
	"math/rand"
)

func TestSBlock() {
	fmt.Print("Starting S-block testing...\n")
	sbox := blocks.GetSBox()
	invsbox := blocks.InverseSBox(sbox)

	for i := 0; i < 20; i++ {
		input := uint8(rand.Intn(256))
		sOut := blocks.Sblock(input, sbox)
		sIn := blocks.InvSblock(sOut, invsbox)

		if sIn == input {
			fmt.Printf("Input:  %X passed!\n", input)
		} else {
			fmt.Printf("Input:  %X failed!\n", input)
		}
	}

	fmt.Print("S-block testing finished!\n\n")
}

func TestPBlock() {
	fmt.Print("Starting P-block testing...\n")

	for i := 0; i < 20; i++ {
		input := uint8(rand.Intn(256))
		pOut := blocks.Pblock(input)
		pIn := blocks.InversePBlock(pOut)

		if pIn == input {
			fmt.Printf("Input:  %X passed!\n", input)
		} else {
			fmt.Printf("Input:  %X failed!\n", input)
		}
	}

	fmt.Print("P-block testing finished!\n\n")
}

func TestPnSblocks() {
	fmt.Print("Starting combined S- and P- blocks testing...\n")
	sbox := blocks.GetSBox()
	invsbox := blocks.InverseSBox(sbox)

	for i := 0; i < 20; i++ {
		input := uint8(rand.Intn(256))
		sOut := blocks.Sblock(input, sbox)
		pOut := blocks.Pblock(sOut)
		pIn := blocks.InversePBlock(pOut)
		sIn := blocks.InvSblock(pIn, invsbox)

		if sIn == input {
			fmt.Printf("Input:  %X passed!\n", input)
		} else {
			fmt.Printf("Input:  %X failed!\n", input)
		}
	}

	fmt.Print("Combined S- and P- blocks testing finished!\n\n")
}
