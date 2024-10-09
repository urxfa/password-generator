package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/urxfa/password-generator/password"
)

func main() {
	argv := os.Args

	if len(argv) != 3 {
		fmt.Println("\nUsage: ./passgen [complexity] [length]")
		fmt.Printf("Ex: ./passgen 5 32\n\n")
		return
	}

	complexity, err := strconv.ParseUint(argv[1], 10, 8)
	if err != nil || complexity < 1 || complexity > 5 {
		fmt.Println("Use a number from 1 to 5 for complexity")
		return
	}

	length, err2 := strconv.ParseUint(argv[2], 10, 8)
	if err2 != nil || length < 4 || length > 64 {
		fmt.Println("Length: Min: 4 | Max: 64")
		return
	}
	pass := password.PasswordBuild{Complexity: uint8(complexity), Length: uint8(length)}
	fmt.Println(pass.Generate())
}
