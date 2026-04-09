package sha1

func completeMessage(input []byte) []byte {
	initLen := uint64(len(input) * 8) //length in bits
	input = append(input, 0x80)       // add "1" bit

	//append with "0"s to last block of length 448 bits
	for len(input)%64 != 56 {
		input = append(input, 0x00)
	}

	//divide uint64 initLen var into two 32 bits words
	high, low := uint32(initLen>>32), uint32(initLen&0xFFFFFFFF)

	//append length of initial message in bits in big-endian format to the resulted message, completed with '1' bit and '0's bits
	input = append(input,
		byte(high>>24), byte(high>>16), byte(high>>8), byte(high),
		byte(low>>24), byte(low>>16), byte(low>>8), byte(low),
	)

	return input
}

func leftRotate(value uint32, shift uint) uint32 {
	return (value << shift) | (value >> (32 - shift))
}

func Sum(input []byte) [20]byte {

	var A uint32 = 0x67452301
	var B uint32 = 0xEFCDAB89
	var C uint32 = 0x98BADCFE
	var D uint32 = 0x10325476
	var E uint32 = 0xC3D2E1F0

	input = completeMessage(input)

	for block := 0; block < len(input); block += 64 {
		var w [80]uint32
		//divide into 32-bit chunks
		for i := 0; i < 16; i++ {
			w[i] = uint32(input[block+i*4])<<24 | uint32(input[block+i*4+1])<<16 | uint32(input[block+i*4+2])<<8 | uint32(input[block+i*4+3])
		}

		for i := 16; i < 80; i++ {
			value := w[i-3] ^ w[i-8] ^ w[i-14] ^ w[i-16]
			w[i] = leftRotate(value, 1)
		}

		a, b, c, d, e := A, B, C, D, E

		for i := 0; i < 80; i++ {
			var f, k uint32

			switch {
			case i < 20:
				f = (b & c) | ((^b) & d)
				k = 0x5A827999
			case i < 40:
				f = b ^ c ^ d
				k = 0x6ED9EBA1
			case i < 60:
				f = (b & c) | (b & d) | (c & d)
				k = 0x8F1BBCDC
			default:
				f = b ^ c ^ d
				k = 0xCA62C1D6
			}

			temp := leftRotate(a, 5) + f + e + k + w[i]
			e = d
			d = c
			c = leftRotate(b, 30)
			b = a
			a = temp
		}

		A += a
		B += b
		C += c
		D += d
		E += e
	}

	//write result as sequnce of bytes in big-endian format
	var hash [20]byte
	for i := uint(0); i < 4; i++ {
		hash[i] = byte(A >> (24 - i*8))
		hash[4+i] = byte(B >> (24 - i*8))
		hash[8+i] = byte(C >> (24 - i*8))
		hash[12+i] = byte(D >> (24 - i*8))
		hash[16+i] = byte(E >> (24 - i*8))
	}

	return hash
}
