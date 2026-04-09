package randomlesstests

import "math"

func PokerTest(data []byte) bool {
	const m = 4
	const lower = 1.03
	const upper = 57.4
	n := float64((len(data) * 8) / m) //blocks quantity

	blockCount := make(map[uint8]int)
	mask := byte((1 << m) - 1) // mask to choose m bits

	for _, b := range data {
		for i := 0; i < 8; i += m {
			block := (b >> i) & mask
			blockCount[block]++
		}
	}

	sum := 0.0

	for _, blockFreq := range blockCount {
		sum += float64(blockFreq * blockFreq)
	}

	x3 := (math.Pow(2, m)/n)*sum - n

	return lower <= x3 && x3 <= upper
}
