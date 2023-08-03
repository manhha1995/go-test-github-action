package stringtokenizer

type StringTokenizer struct {
	Text     string
	Position int
	Value    Token // reference to the token interface
}

type Token interface {
	GetWord() string
}

type ErrorToken struct {
	ERROR string
}
