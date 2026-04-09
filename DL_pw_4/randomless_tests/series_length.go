package randomlesstests

//import "fmt"

func SeriesLengthTest(data []byte) bool {
	lower := map[uint8]int{
		1: 2267,
		2: 1079,
		3: 502,
		4: 223,
		5: 90,
		6: 90,
	}

	upper := map[uint8]int{
		1: 2733,
		2: 1421,
		3: 748,
		4: 402,
		5: 223,
		6: 223,
	}

	zeroSeriesCount := make(map[uint8]int)
	oneSeriesCount := make(map[uint8]int)
	var zeroSeries, oneSeries uint8
	var currentSeria bool // true for seria of ones' and false for seria of zeros'
	seriaStarted := false
	var currentBit uint8
	var previousBit uint8

	for key, b := range data {
		for i := 7; i >= 0; i-- {
			if key == 0 && i == 7 {
				continue
			} else {
				previousBit = (b & (1 << (i + 1))) >> (i + 1)
			}

			currentBit = (b & (1 << i)) >> i

			if key == len(data)-1 && i == 0 {
				if currentBit == previousBit {
					continue
				}

				if currentSeria {
					oneSeriesCount[oneSeries]++
				} else {
					zeroSeriesCount[zeroSeries]++
				}
			}

			if !seriaStarted && (currentBit == previousBit) {
				continue
			} else if !seriaStarted && (currentBit != previousBit) {
				seriaStarted = true
			}

			if seriaStarted {
				if currentBit == 0 {
					zeroSeries++
					oneSeriesCount[oneSeries]++
					oneSeries = 0
					currentSeria = false
				} else {
					oneSeries++
					zeroSeriesCount[zeroSeries]++
					zeroSeries = 0
					currentSeria = true
				}
			}
		}
	}

	for key, value := range zeroSeriesCount {

		if key == 0 {
			continue
		}
		if key > 6 {
			key = 6
		}
		//fmt.Printf("Series of zeros' with length %d appears %d times\n", key, value) // you can check how many times series of ones or zeroes apper
		if value < lower[key] || value > upper[key] {
			return false
		}
	}

	for key, value := range oneSeriesCount {
		if key == 0 {
			continue
		}
		if key > 6 {
			key = 6
		}
		//fmt.Printf("Series of ones' with length %d appears %d times\n", key, value) // you can check how many times series of ones or zeroes apper
		if value < lower[key] || value > upper[key] {
			return false
		}
	}
	return true

}
