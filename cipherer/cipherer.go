package cipherer

import (
	"encoding/base64"
	"fmt"
	"os"
)

func Cipher(rawString, secret string) string {
  processed := process([]byte(rawString), []byte(secret))
  return base64.StdEncoding.EncodeToString(processed)
}

func Decipher(cipheredText, secret string) string {
  cipheredBytes, err := base64.StdEncoding.DecodeString(cipheredText)
  if err != nil {
    fmt.Print("error occured while decoding cipheredtext to base64")
    os.Exit(1)
  }

  decryptedBytes := process(cipheredBytes, []byte(secret))
  if len(decryptedBytes) == 0 {
    fmt.Println("decryption failed, invalid input, exiting...")
    os.Exit(1)
  }

  return string(decryptedBytes)
}

func process(input, secret []byte) []byte {
  if len(secret) == 0 {
    fmt.Println("secret key cannot be empty. Fix it")
    os.Exit(1)
  }
  for i, b := range input {
    input[i] = b ^ secret[i % len(secret)]
  }
  return input
}
