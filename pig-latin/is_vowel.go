package piglatin

var (
	vowels = [...]byte{'a', 'e', 'i', 'o', 'A', 'E', 'I', 'O'}
)

// isVowel checks if this byte is vowel
func isVowel(letter byte) bool {
	for _, vowel := range vowels {
		if letter == vowel {
			return true
		}
	}
	return false
}
