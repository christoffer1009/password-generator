package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func GetRandomChar(inputString string) string {
	if len(inputString) == 0 {
		return ""
	}
	randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(inputString))))
	if err != nil {
		return ""
	}
	return string(inputString[randomIndex.Int64()])
}

func ShuffleString(inputString string) (string, error) {
	runes := []rune(inputString)
	n := len(runes)
	for i := n - 1; i > 0; i-- {
		j, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			return "", fmt.Errorf("error generating random number")
		}
		runes[i], runes[j.Int64()] = runes[j.Int64()], runes[i]
	}

	return string(runes), nil
}
