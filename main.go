package main

import (
	"errors"
	"fmt"
)

func Tokenizer(text string, value Token) *StringTokenizer {
	return &StringTokenizer{
		Text:  text,
		Value: value,
	}
}

func (s *StringTokenizer) String(source string) ([]string, error) {
	input := []string{source}
	if len(s.Text) == 0 {
		return input, errors.New("cannot tokenize errors")
	}
	return input, nil
}

func Error(text string) *ErrorToken {
	if text == "" || len(text) == 0 {
		fmt.Println("token not valid")
	}
	return &ErrorToken{
		ERROR: string(text),
	}
}

func main() {
	fmt.Println("starting process...")
}
