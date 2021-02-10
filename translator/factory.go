package translator

import "errors"

type TranslateFactory interface {
	TranslateToText(text string) string
	TranslateToTarget(text string) string
}

func GetInterpreterSource(source string) (TranslateFactory, error) {
	if source == "BINARY" {
		return Binary{}, nil
	}
	if source == "TEXT" {
		return Text{}, nil
	}
	if source == "MORSE" {
		return Morse{}, nil
	}


	return nil, errors.New("invalid source format")
}

func GetInterpreterTarget(target string) (TranslateFactory, error) {
	if target == "BINARY" {
		return Binary{}, nil
	}
	if target == "TEXT" {
		return Text{}, nil
	}
	if target == "MORSE" {
		return Morse{}, nil
	}
	return nil, errors.New("invalid target format")
}
