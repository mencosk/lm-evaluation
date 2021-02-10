package translator

import "fmt"

type Morse struct{}




func (Morse) TranslateToText(text string) string {
	return ""
}

func (Morse) TranslateToTarget(text string) string {
	var translation string
	for _, c := range text {
		translation = fmt.Sprintf("%s%s", translation, textToMorse(string(c)))
	}
	return translation
}

func textToMorse(text string) string{
	morseTable := map[string]string {
		"A": ".-",
		"B": "-...",
		"C": "-.-.",
		"D": "-..",
		"E": ".",
		"F": "..-.",
		"G": "--.",
		"H": "....",
		"I": "..",
		"J": ".---",
		"K": "-.-",
		"L": ".-..",
		"M": "--",
		"N": "-.",
		"O": "---",
	}
	return morseTable[text]
}