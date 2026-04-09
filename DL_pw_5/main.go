package main

import (
	test "DL_pw_5/test"
	"crypto/rand"
	"fmt"
)

func main() {
	for i := 0; i < 20; i++ {
		token := make([]byte, 2000*i)
		rand.Read(token)

		fmt.Printf("Test #%d with %d input bytes\n", i, 2000*i)
		test.FuncTest(token)
		test.SpeedTest(token)
		fmt.Println()
	}

}
