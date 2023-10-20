package main

import (
	"fmt"
	"gounseen/internal/translate"
	"os"
	"strings"
)

func main() {
	translate.InitializeLetters()

	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("please, inform an operation (encode/decode) and a phrase to process")
		os.Exit(1)
	}

	cmd := args[0]
	phrase := strings.Join(args[1:], " ")
	switch cmd {
	case "encode":
		fmt.Printf("encoding '%s':\n", phrase)
		fmt.Printf("result: '%s'\n", translate.Encode(phrase))
		os.Exit(0)
	case "decode":
		fmt.Printf("decoding '%s':\n", phrase)
		fmt.Printf("result: '%s'\n", translate.Decode(phrase))
		os.Exit(0)
	}

	fmt.Println("invalid command. the first argument of this program should be 'encode' or 'decode'")
	os.Exit(2)
}
