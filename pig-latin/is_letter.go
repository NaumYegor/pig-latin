package piglatin

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
