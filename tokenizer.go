package lisp

import (
	"fmt"
	"regexp"
	"strings"
)

func computeLineColumn(source string, index int) (line, column int) {
	lines := strings.Split(source, "\n")
	totalChars := 0

	for i, line := range lines {
		lineLength := len(line) + 1
		if index < totalChars+lineLength {
			lineIndex := index - totalChars
			return i + 1, lineIndex + 1
		}
		totalChars += lineLength
	}

	panic("character index out of range")
}

type TokenizeError struct {
	Source   *string
	Location int
	Message  string
}

func (e TokenizeError) Error() string {
	line, col := computeLineColumn(*e.Source, e.Location)
	return fmt.Sprintf(`[%d:%d] %s`, line, col, e.Message)
}

type rule struct {
	Type   TokenType
	Regex  *regexp.Regexp
	Ignore bool
}

var (
	CommentToken     TokenType = "Comment"
	FloatToken       TokenType = "Float"
	IntegerToken     TokenType = "Integer"
	StringToken      TokenType = "String"
	PunctuationToken TokenType = "Punctuation"
	WhitespaceToken  TokenType = "Whitespace"
	IdentifierToken  TokenType = "Identifier"
)

var rules = []rule{
	{Type: CommentToken, Ignore: true,
		Regex: regexp.MustCompile(`^//.*`)},
	{Type: FloatToken,
		Regex: regexp.MustCompile(`^[0-9]+\.[0-9]+`)},
	{Type: IntegerToken,
		Regex: regexp.MustCompile(`^[0-9]+`)},
	{Type: StringToken,
		Regex: regexp.MustCompile(`^"(\\.|[^"])*"`)},
	{Type: PunctuationToken,
		Regex: regexp.MustCompile(`^[\#\$\.\(\)\[\]\{\}]`)},
	{Type: WhitespaceToken, Ignore: true,
		Regex: regexp.MustCompile(`^[ \t\n]+`)},
	{Type: IdentifierToken,
		Regex: regexp.MustCompile(`^[^\#\$\.\(\)\[\]\{\}\s]+`)},
}

func matchRules(source string) (*Token, bool) {
	for _, rule := range rules {
		match := rule.Regex.FindString(source)
		if match != "" {
			return &Token{Type: rule.Type, Value: match}, rule.Ignore
		}
	}

	return nil, true
}

func Tokenize(source string) ([]Token, error) {
	cursor := 0
	tokens := []Token{}

	for cursor < len(source) {
		remaining := source[cursor:]

		t, ignore := matchRules(remaining)
		if t == nil {
			return nil, TokenizeError{&source, cursor, "unexpected character"}
		}

		cursor += len(t.Value)
		if !ignore {
			tokens = append(tokens, *t)
		}
	}

	return tokens, nil
}
