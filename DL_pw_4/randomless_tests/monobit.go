package randomlesstests

func MonobitTest(input []byte) bool {
	const (
		n     = 20000
		lower = 9654
		upper = 10346
	)

	var count int

	for _, b := range input {
		for i := 0; i < 8; i++ {
			if b&(1<<i) != 0 {
				count++
			}
		}
		if count > upper {
			return false
		}
	}

	return lower <= count
}
