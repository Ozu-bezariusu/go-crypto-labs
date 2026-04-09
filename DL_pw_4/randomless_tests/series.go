package randomlesstests

func SeriesTest(data []byte) bool {
	const maxSeriesLength = 36

	var zeroSeries, oneSeries int
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

				if zeroSeries > maxSeriesLength || oneSeries > maxSeriesLength {
					return false
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

					if zeroSeries > maxSeriesLength {
						return false
					}

					oneSeries = 0
				} else {
					oneSeries++
					if oneSeries > maxSeriesLength {
						return false
					}
					zeroSeries = 0
				}
			}
		}
	}

	return true
}
