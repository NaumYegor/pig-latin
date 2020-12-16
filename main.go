package main

import (
	"fmt"
	"github.com/naumyegor/pig-latin/input"
	piglatin "github.com/naumyegor/pig-latin/pig-latin"
	"github.com/naumyegor/pig-latin/translator"
)

func main() {
	var pigLatin input.Input
	var encdec input.Input

	fmt.Println("Enter your phrase in pig latin:")
	pigLatin.ReadInput()
	fmt.Println(pigLatin.TranslateText(piglatin.TranslateWord))

	fmt.Println("Enter your phrase to encode:")
	encdec.ReadInput()
	fmt.Println(encdec.TranslateText(translator.EncodeWord))

	fmt.Println("Enter your phrase to decode:")
	encdec.ReadInput()
	fmt.Println(encdec.TranslateText(translator.DecodeWord))
}
