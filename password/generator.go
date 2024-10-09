package password

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var r *rand.Rand

/*
Complexity is a number from 1 to 5:

	1 - only numbers password
	2 - only uppercase alphanumeric password
	3 - only lowercase alphanumeric password
	4 - lower and upper alphanumeric password
	5 - lower and upper alphanumeric and must have 4 or more special characters
*/
type PasswordBuild struct {
	Complexity uint8 // From 1 to 5
	Length     uint8 // Max 64
}

func init() {
	src := rand.NewSource(time.Now().Unix())
	r = rand.New(src)
}

func (p PasswordBuild) Generate() string {
	if !isValidLength(p.Length) {
		return "Length should be from 4 to 32"
	}

	fmt.Printf("Generating password - Length: %d | Complexity: %d\n", p.Length, p.Complexity)

	switch p.Complexity {
	case 1:
		return onlyNumbers(p.Length)
	case 2:
		return alphanumericStraight(p.Length, true)
	case 3:
		return alphanumericStraight(p.Length, false)
	case 4:
		return alphanumericMixed(p.Length)
	case 5:
		return mixed(p.Length)
	default:
		panic("Invalid Option")
	}

}

func isValidLength(length uint8) bool {
	return length >= 4 && length <= 64
}

func getRandomSpecialChar() byte {
	var special []byte = []byte{
		'!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '_', '=', '+', '[', ']', '{', '}', ':', ';', ',', '.', '?', '/',
	}

	random := special[r.Intn(len(special))]
	return random
}

func onlyNumbers(length uint8) string {
	var p string

	for i := 0; i < int(length); i++ {
		p = p + string(r.Intn(57-48+1)+48) // ascii code for 0 to 9
	}
	return p
}

func alphaChar(lowercase bool) byte {
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	character := letters[r.Intn(len(letters))]
	if lowercase {
		character = strings.ToLower(string(character))[0]
	}

	return character
}

func alphanumericStraight(length uint8, lowercase bool) string {
	var p strings.Builder

	for i := 0; i < int(length); i++ {
		p.WriteByte(alphaChar(lowercase))
	}

	return p.String()
}

func alphanumericMixed(length uint8) string {
	var p strings.Builder

	for i := 0; i < int(length); i++ {
		var character byte
		percentil := r.Float32() * 100

		if percentil >= 50 {
			character = alphaChar(true) // lowercase
		} else {
			character = alphaChar(false) // uppercase
		}

		p.WriteByte(character)
	}
	return p.String()
}

func mixed(length uint8) string {
	var p strings.Builder

	for i := 0; i < int(length); i++ {
		var character byte
		var percentil = r.Float32() * 100

		if percentil > 33 && percentil <= 66 {
			character = getRandomSpecialChar()
		} else if percentil <= 33 {
			character = alphaChar(false)
		} else {
			character = alphaChar(true)
		}

		p.WriteByte(character)
	}
	return p.String()
}
