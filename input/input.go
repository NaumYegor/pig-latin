package input

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var (
	vowels =  [...]byte {'a', 'e', 'i', 'o', 'A', 'E', 'I', 'O'}
)

const (
	lowerLetterIndexCap   = 65
	upperLetterIndexCap   = 90
	lowerLetterIndexSmall = 97
	upperLetterIndexSmall = 122
)

func isLetter (letter byte) bool {
	isCapital := letter <= upperLetterIndexCap && letter >= lowerLetterIndexCap
	isSmall := letter <= upperLetterIndexSmall && letter >= lowerLetterIndexSmall

	if !isCapital && !isSmall {
		return false
	}
	return true
}

func isVowel(letter byte) bool {
	for _, vowel := range vowels {
		if letter == vowel {
			return true
		}
	}
	return false
}

type Input struct {
	text 	  string
	separatedWords []string
}

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

func (i *Input) getSeparatedWords()  {
	separatedWords := strings.Split(i.text, " ")
	i.separatedWords = separatedWords
}

func (i *Input) TranslateText() string {
	translatedWords := make([]string, 0)

	i.getSeparatedWords()

	for _, word := range i.separatedWords {
		translatedWords = append(translatedWords, translateWord(word))
	}


	return strings.ToLower(strings.Join(translatedWords, " "))
}

func translateWord(word string) string {
	if len(word) == 1 && !isLetter(word[0]) {
		return word
	}

	lastIsLetter := isLetter(word[len(word)-1])

	translated := make([]byte, 0)
	for j, letter := range []byte(word) {
		if isVowel(letter) {
			if lastIsLetter {
				translated = append(translated, word[j:]...)
				translated = append(translated, word[:j]...)
				if translated[len(translated)-1] == 'a' {
					translated = append(translated, 'y')
				} else {
					translated = append(translated, "ay"...)
				}
			} else {
				translated = append(translated, word[j:len(word)-1]...)
				translated = append(translated, word[:j]...)
				if translated[len(translated)-1] == 'a' {
					translated = append(translated, 'y')
				} else {
					translated = append(translated, "ay"...)
				}
				translated = append(translated, word[len(word)-1])
			}
			break
		}
	}

	return string(translated)
}
