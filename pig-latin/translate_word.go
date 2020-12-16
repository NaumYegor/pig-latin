package piglatin

// TranslateWord translates word to pig latin
func TranslateWord(word string) string {
	if len(word) == 1 {
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
