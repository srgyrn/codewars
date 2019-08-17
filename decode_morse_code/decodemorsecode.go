package decode_morse_code

import (
	"strings"
)

func DecodeMorse(morseCode string) string {
	if len(morseCode) == 0 {
		return ""
	}

	morseWords := strings.Split(strings.TrimSpace(morseCode), "   ")
	result := make([]string, len(morseWords))
	var word string

	for _, morseWord := range morseWords {
		word = ""
		for _, letter := range strings.Split(strings.TrimSpace(morseWord), " ") {
			if letter, exists := MORSE_CODE[letter]; exists {
				word += letter
			}
		}
		result = append(result, word)
	}

	return strings.TrimSpace(strings.Join(result, " "))
}
