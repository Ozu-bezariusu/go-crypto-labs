package main

import (
	elg "DL_pw_9/elg"
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter message to sign: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading from console:", err)
		return
	}

	//mt.Println("You entered:", text)

	p, g := elg.PrepareParams(2048)
	a, b := elg.GenerateKeys(p, g)

	message := []byte(text)

	r, s := elg.Sign(p, g, a, b, message)

	verification := elg.CheckSignature(p, g, a, b, r, s, message)

	if verification {
		fmt.Print("Message signed correct")
	} else {
		fmt.Print("Message signed incorrect")
	}
}
