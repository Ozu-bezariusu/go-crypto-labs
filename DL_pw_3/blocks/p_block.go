package blocks

func Pblock(input byte) byte {
	//set permitation table
	permutation := []int{7, 0, 5, 2, 3, 4, 1, 6}
	//init zero byte
	output := byte(0)

	for i, p := range permutation {
		//check whether bit is 1 and if it is we set that to appropriate position defined by permutation table
		//if bit is 0 we just skip round, because output byte is already a sequence of zero bits
		if (input>>uint(p))&1 == 1 {
			output |= 1 << uint(i)
		}
	}

	return output
}

func InversePBlock(input byte) byte {
	permutation := []int{7, 0, 5, 2, 3, 4, 1, 6}
	output := byte(0)

	for i, p := range permutation {
		if (input>>uint(i))&1 == 1 {
			output |= 1 << uint(p)
		}
	}

	return output
}
