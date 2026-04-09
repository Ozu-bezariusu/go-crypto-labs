package main

import (
	tests "DL_pw_7/tests"
	"fmt"
)

func main() {
	if tests.ECTest() {
		fmt.Print("Correct!")
	} else {
		fmt.Print("Incorrect")
	}
}
