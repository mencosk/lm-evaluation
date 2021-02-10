package main

import (
	"fmt"
	"github.com/mencosk/lm-evaluation/translator"
)

func main() {
	var interpreter translator.Service
	interpreter = translator.Interpreter{}
	fmt.Println(interpreter.Translate("AAAA","TEXT", "BINARY"))
	fmt.Println(interpreter.Translate("AAAA","TEXT", "MORSE"))
}
