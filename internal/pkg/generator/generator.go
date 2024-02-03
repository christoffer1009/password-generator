package generator

import (
	"crypto/rand"
	"fmt"

	"github.com/christoffer1009/password-generator/internal/entity"
	"github.com/christoffer1009/password-generator/internal/pkg/utils"
)

func GeneratePassword(options *entity.Options) (string, error) {
	var upper, lower, digit, special, charset, minCharset string
	var upperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var lowerCase = "abcdefghijklmnopqrstuvwxyz"
	var digits = "0123456789"
	var specialChars = "!@#$%^&*()_+-=[]{};:,<.>?"

	if options.UpperCase {
		charset += upperCase
		upper = utils.GetRandomChar(upperCase)
		minCharset += upper
	}
	if options.LowerCase {
		charset += lowerCase
		lower = utils.GetRandomChar(lowerCase)
		minCharset += lower
	}
	if options.Digits {
		charset += digits
		digit = utils.GetRandomChar(digits)
		minCharset += digit
	}
	if options.SpecialChars {
		charset += specialChars
		special = utils.GetRandomChar(specialChars)
		minCharset += special
	}

	if charset == "" {
		return "", fmt.Errorf("at least on opeion must be true")
	}

	minCharset, err := utils.ShuffleString(minCharset)
	if err != nil {
		return "", err
	}

	buffer := make([]byte, options.Length-len(minCharset))
	_, err = rand.Read(buffer)
	if err != nil {
		return "", err
	}

	for i, v := range buffer {
		buffer[i] = charset[int(v)%len(charset)]
	}

	return string(buffer) + minCharset, nil
}
