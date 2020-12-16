package input

var (
	vowels = [...]byte{'a', 'e', 'i', 'o', 'A', 'E', 'I', 'O'}
)

const (
	lowerLetterIndexCap   = 65
	upperLetterIndexCap   = 90
	lowerLetterIndexSmall = 97
	upperLetterIndexSmall = 122
)

// isLetter checks if this byte is letter in ascii or another symbol
func isLetter(letter byte) bool {
	isCapital := letter <= upperLetterIndexCap && letter >= lowerLetterIndexCap
	isSmall := letter <= upperLetterIndexSmall && letter >= lowerLetterIndexSmall

	if !isCapital && !isSmall {
		return false
	}
	return true
}

// isVowel checks if this byte is vowel
func isVowel(letter byte) bool {
	for _, vowel := range vowels {
		if letter == vowel {
			return true
		}
	}
	return false
}

// translateWord translates word to pig latin
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
