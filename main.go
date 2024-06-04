package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"xor/cipherer"
)

var mode = flag.String("mode", "cipher", "Set to 'cipher' or 'decipher'. Default is 'cipher'.")
var secretKey = flag.String("secret", "", "Your secret key. Must contain at least 1 character")

func main() {
	flag.Parse()

	switch *mode {
	case "cipher":
		plainText := getUserInput("enter your text to cipher: ")
		fmt.Println(plainText)
		fmt.Println(cipherer.Cipher(plainText, *secretKey))
	// ..
	case "decipher":
		cipheredText := getUserInput("enter your text to decipher: ")
		fmt.Println(cipheredText)
		fmt.Println(cipherer.Decipher(cipheredText, *secretKey))
	// ..
	default:
		fmt.Println("please specify cipher or decipher")
		os.Exit(1)
	}

	if len(*secretKey) == 0 {
		fmt.Println("no secret provided")
		os.Exit(1)
	}

}

func getUserInput(msg string) string {
	fmt.Print(msg)

	reader := bufio.NewReader(os.Stdin)

	for {
		result, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("an error occured while reading text, try again")
			continue
		}

		return strings.TrimRight(result, "\n")
	}
}
