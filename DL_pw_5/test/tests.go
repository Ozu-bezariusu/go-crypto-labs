package test

import (
	mysha1 "DL_pw_5/sha-1"
	"crypto/sha1"
	"fmt"
	"time"
)

func FuncTest(input []byte) {
	customSum := mysha1.Sum(input)
	sum := sha1.Sum(input)

	if sum == customSum {
		fmt.Println("Custom algorithm works correct \u2714")
	} else {
		fmt.Println("Custom algorithm works incorrect \u2716")
	}
}

func SpeedTest(input []byte) {

	customStart := time.Now()
	mysha1.Sum(input)
	customDuration := time.Since(customStart).Microseconds()

	start := time.Now()
	sha1.Sum(input)
	duration := time.Since(start).Microseconds()

	if customDuration > duration {
		fmt.Printf("Custom implementation works slower by %d  microseconds\n", customDuration-duration)
	} else {
		fmt.Printf("Custom implementation works faster by %d microseconds\n", duration-customDuration)
	}
}
