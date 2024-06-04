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

	if len(*secretKey) == 0 {
		fmt.Fprintln(os.Stderr, "no secret provided")
		os.Exit(1)
	}

	switch *mode {
	case "cipher":
		plainText := getUserInput("enter your text to cipher: ")
		fmt.Println(plainText)

		cipheredText, err := cipherer.Cipher(plainText, *secretKey)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error encrypting text: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(cipheredText)
	case "decipher":
		cipheredText := getUserInput("enter your text to decipher: ")

		decipheredText, err := cipherer.Decipher(cipheredText, *secretKey)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Decription error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(decipheredText)
	default:
		fmt.Println("please specify cipher or decipher")
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
