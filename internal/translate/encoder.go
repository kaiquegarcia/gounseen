package translate

import (
	"strings"
)

var translationMap = map[string]string{
	"ㅏ": "A",
	"Б": "B",
	"C": "C",
	"𑀘": "C",
	"Δ": "D",
	"Э": "E",
	"Ф": "F",
	"Γ": "G",
	"ㅎ": "H",
	"И": "I",
	"𑀛": "J",
	"ㅈ": "J",
	"ㅋ": "K",
	"Λ": "L",
	"ㅁ": "M",
	"N": "N",
	"𑀦": "N",
	"Ω": "O",
	"Π": "P",
	"Q": "Q",
	"ㄹ": "R",
	"Σ": "S",
	"T": "T",
	"U": "U",
	"V": "V",
	"W": "W",
	"X": "X",
	"Y": "Y",
	"З": "Z",
}

var (
	encodedLetters = make([]string, len(translationMap))
	decodedLetters = make([]string, len(translationMap))
)

func InitializeLetters() {
	index := 0
	for encodedLetter, decodedLetter := range translationMap {
		encodedLetters[index] = encodedLetter
		decodedLetters[index] = decodedLetter
		index++
	}
}

func Decode(encodedPhrase string) string {
	encodedPhraseLetters := strings.Split(encodedPhrase, "")
	phraseLen := len(encodedPhraseLetters)
	if phraseLen == 0 {
		return ""
	}

	output := ""
	for charIndex := 0; charIndex < phraseLen; charIndex++ {
		letter := encodedPhraseLetters[charIndex]
		decodedLetter, exists := translationMap[letter]
		if exists {
			output = output + decodedLetter
		} else {
			output = output + letter
		}
	}

	return output
}

func Encode(decodedPhrase string) string {
	decodedPhraseLetters := strings.Split(decodedPhrase, "")
	phraseLen := len(decodedPhraseLetters)
	if phraseLen == 0 {
		return ""
	}

	output := ""
	for charIndex := 0; charIndex < phraseLen; charIndex++ {
		letter := decodedPhraseLetters[charIndex]
		letterIndex := getDecodedLetterIndex(letter)
		if letterIndex >= 0 {
			output = output + encodedLetters[letterIndex]
		} else {
			output = output + letter
		}
	}

	return output
}

func getDecodedLetterIndex(letter string) int {
	for index, decodedLetter := range decodedLetters {
		if decodedLetter == letter {
			return index
		}
	}

	return -1
}
