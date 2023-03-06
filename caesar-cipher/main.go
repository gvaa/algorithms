package main

import (
	"fmt"
	"strings"
)

type cipher struct {
	codedWord string
	step      int
}

func main() {

	codedWords := []cipher{
		{codedWord: "Ncevy", step: 13},
		{codedWord: "gpvsui", step: 25},
		{codedWord: "ugflgkg", step: -18},
		{codedWord: "wjmmf", step: -1},
	}

	decipheredMessage := ""

	for _, cipherToSolve := range codedWords {
		decipheredMessage = decipheredMessage + lassoWord(strings.ToLower(cipherToSolve.codedWord), cipherToSolve.step) + " "
	}

	fmt.Println(decipheredMessage)

}

func lassoLetter(letter rune, shiftAmount int) string {

	alphabet := 26
	aCode := int('a')
	letterCode := int(letter)

	trueLetterCode := aCode + (((letterCode - aCode) + shiftAmount) % alphabet)

	if (((letterCode - aCode) + shiftAmount) % alphabet) < 0 {
		trueLetterCode = aCode + (((letterCode - aCode) + shiftAmount) % alphabet) + alphabet
	}

	decodedLetter := string(rune(trueLetterCode))

	return decodedLetter
}

func lassoWord(word string, shift_amount int) string {

	decodedWord := ""

	for _, char := range word {
		decodedLetter := lassoLetter(char, shift_amount)
		decodedWord = decodedWord + decodedLetter
	}

	return decodedWord
}
