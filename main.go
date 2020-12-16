package main

import (
	"fmt"
	"github.com/naumyegor/test-translators/input"
	"github.com/naumyegor/test-translators/leet-parody"
	"github.com/naumyegor/test-translators/pig-latin"
)

func main() {
	var pigLatin input.Input
	var encdec input.Input

	fmt.Println("Enter your phrase in pig latin:")
	pigLatin.ReadInput()
	fmt.Println(pigLatin.TranslateText(piglatin.TranslateWord))

	fmt.Println("Enter your phrase to encode:")
	encdec.ReadInput()
	fmt.Println(encdec.TranslateText(leetparody.EncodeWord))

	fmt.Println("Enter your phrase to decode:")
	encdec.ReadInput()
	fmt.Println(encdec.TranslateText(leetparody.DecodeWord))
}
