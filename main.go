package main

import (
	"fmt"
	"github.com/naumyegor/pig-latin/input"
)

func main() {
	var inp input.Input
	fmt.Println("Enter your phrase in pig latin:")
	inp.ReadInput()
	fmt.Println(inp.TranslateText())
}
