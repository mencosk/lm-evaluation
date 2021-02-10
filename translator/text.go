package translator

type Text struct{}

func (Text) TranslateToText(text string) string {
	return text
}

func (Text) TranslateToTarget(text string) string {
	return text
}
