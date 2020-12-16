package input

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Input struct {
	text           string
	separatedWords []string
}

// ReadInput reads one line from the os.Stdin
func (i *Input) ReadInput() {
	reader := bufio.NewReader(os.Stdin)

	val, err := reader.ReadString('\n')

	if err != nil {
		log.Fatalln("Failed to read input")
	}

	bytesVal := []byte(val)

	if bytesVal[len(bytesVal)-1] == '\n' {
		bytesVal = append(bytesVal[:len(bytesVal)-1], bytesVal[len(bytesVal):]...)
	}

	i.text = string(bytesVal)
}

//TranslateText translates all words in the Input and returns translated text
func (i *Input) TranslateText() string {
	translatedWords := make([]string, 0)

	i.getSeparatedWords()

	for _, word := range i.separatedWords {
		translatedWords = append(translatedWords, translateWord(word))
	}

	return strings.ToLower(strings.Join(translatedWords, " "))
}

//getSeparatedWords updates the Input.separatedWords field of Input object
func (i *Input) getSeparatedWords() {
	separatedWords := strings.Split(i.text, " ")
	i.separatedWords = separatedWords
}
