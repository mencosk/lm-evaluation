package translator

import "fmt"

type Binary struct{}

func (Binary) TranslateToText(text string) string {
	return "AAAAAAA"
}

func (Binary) TranslateToTarget(text string) string {
	var translation string
	for _, c := range text {
		translation = fmt.Sprintf("%s%.8b", translation, c)
	}
	return translation
}