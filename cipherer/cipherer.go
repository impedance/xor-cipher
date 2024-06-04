package cipherer

import (
	"encoding/base64"
	"errors"
)

func Cipher(rawString, secret string) (string, error) {
  if len(secret) == 0 {
    return "", errors.New("secret key cant be empty")
  }
  processed, err := process([]byte(rawString), []byte(secret))
  if err != nil {
    return "", err
  }
  return base64.StdEncoding.EncodeToString(processed), nil
}

func Decipher(cipheredText, secret string) (string, error) {
  if len(secret) == 0 {
    return "", errors.New("secret key cant be empty")
  }
  cipheredBytes, err := base64.StdEncoding.DecodeString(cipheredText)
  if err != nil {
    return "", errors.New("error occured while decoding cipheredtext to base64")
  }

  decryptedBytes, err := process(cipheredBytes, []byte(secret))

  if err != nil {
    return "", errors.New("decryption failed, invalid input, exiting...")
  }

  return string(decryptedBytes), nil
}

func process(input, secret []byte) ([]byte, error) {
  for i, b := range input {
    input[i] = b ^ secret[i % len(secret)]
  }
  return input, nil
}
