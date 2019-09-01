package decodemorsecode

import (
	"regexp"
	"strings"
)

const (
	pauseBetweenCharacters = 3
	dashLength             = 3
	pauseBetweenWords      = 7
)

var frequency = 1

// DecodeBits converts bits to morse code
func DecodeBits(bits string) string {
	bits = strings.Trim(bits, "0")

	if len(bits) > 1 {
		frequency = findFrequency(bits)
	}

	bits = strings.ReplaceAll(bits, strings.Repeat("0", pauseBetweenWords*frequency), strings.Repeat(" ", pauseBetweenWords))
	bits = strings.ReplaceAll(bits, strings.Repeat("0", pauseBetweenCharacters*frequency), strings.Repeat(" ", pauseBetweenCharacters))
	bits = strings.ReplaceAll(bits, strings.Repeat("1", frequency*dashLength), "-")
	bits = strings.ReplaceAll(bits, strings.Repeat("1", frequency), ".")
	return strings.ReplaceAll(bits, strings.Repeat("0", frequency), "")
}

// DecodeMorse converts morse code to Latin Alphabet
func DecodeMorse(morseCode string) string {
	defer func() {
		frequency = 1
	}()

	if len(morseCode) == 0 {
		return ""
	}

	morseWords := strings.Split(strings.TrimSpace(morseCode), strings.Repeat(" ", pauseBetweenWords))

	result := make([]string, len(morseWords))
	var word string

	for _, morseWord := range morseWords {
		word = ""
		for _, letter := range strings.Split(strings.TrimSpace(morseWord), strings.Repeat(" ", pauseBetweenCharacters)) {
			if letter, exists := MORSE_CODE[letter]; exists {
				word += letter
			}
		}

		if len(word) > 0 {
			result = append(result, strings.TrimSpace(word))
		}
	}

	return strings.TrimSpace(strings.Join(result, " "))
}

func findFrequency(bits string) int {
	if strings.Count(bits, "0") == 0 {
		return len(bits)
	}

	findFq := func(sslice []string) int {
		maxPause := 1
		for _, pause := range sslice {
			if len(pause) > maxPause {
				maxPause = len(pause)
			}
		}

		for _, divider := range [2]int{pauseBetweenWords, pauseBetweenCharacters} {
			if maxPause%divider == 0 {
				if maxPause == divider {
					return maxPause
				}

				return maxPause / divider
			}
		}

		return maxPause
	}

	fq := findFq(regexp.MustCompile("1+").FindAllString(bits, -1))

	if fq != dashLength {
		return fq
	}

	if fq == dashLength && findFq(regexp.MustCompile("0+").FindAllString(bits, -1)) == dashLength {
		return 3
	}

	return 1
}
