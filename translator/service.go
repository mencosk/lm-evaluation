package translator

type Service interface {
	Translate(text string, source string, target string) (string, error)
}

type Interpreter struct{}

func (Interpreter) Translate(text string, source string, target string) (string, error) {
	sourceInterpreter, err := GetInterpreterSource(source)
	if err != nil {
		return "", err
	}
	targetInterpreter, err := GetInterpreterTarget(target)
	if err != nil {
		return "", err
	}

	translateToText := sourceInterpreter.TranslateToText(text)
	translateToTarget := targetInterpreter.TranslateToTarget(translateToText)

	return translateToTarget, nil
}
