package leetparody

import "strings"

var (
	encodeList = map[byte]byte{
		'a': '1',
		'e': '2',
		'i': '3',
		'o': '4',
		'u': '5',
	}
	decodeList = map[byte]byte{
		'1': 'a',
		'2': 'e',
		'3': 'i',
		'4': 'o',
		'5': 'u',
	}
)

func translateWord(text string, dictionary map[byte]byte) string {
	textBytes := []byte(strings.ToLower(text))

	for i, _ := range textBytes {
		if val, ok := dictionary[textBytes[i]]; ok {
			textBytes[i] = val
		}
	}

	return string(textBytes)
}

func EncodeWord(text string) string {
	return translateWord(text, encodeList)
}

func DecodeWord(text string) string {
	return translateWord(text, decodeList)
}
