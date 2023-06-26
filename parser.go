package lisp

import (
	"fmt"
	"strconv"
	"strings"
)

type parser struct {
	tokens []Token
	cursor int
}

func (p *parser) done() bool {
	return p.cursor >= len(p.tokens)
}

func (p *parser) peek() Token {
	return p.tokens[p.cursor]
}

func (p *parser) advance() Token {
	p.cursor++
	return p.tokens[p.cursor-1]
}

func (p *parser) expectValue(value string) error {
	if p.done() {
		return fmt.Errorf(`expected "%s" but got eof`, value)
	}
	if p.peek().Value != value {
		return fmt.Errorf(`expected "%s" but got "%s"`, value, p.peek().Value)
	}
	p.advance()
	return nil
}

func (p *parser) expectType(typ TokenType) (Token, error) {
	if p.done() {
		return Token{}, fmt.Errorf(`expected %v but got eof`, typ)
	}
	if p.peek().Type != typ {
		return Token{}, fmt.Errorf(`expected %v but got %v`, typ, p.peek().Type)
	}
	return p.advance(), nil
}

//
// Parser
//

func (p *parser) parseProgram() (any, error) {
	statements := []any{Symbol{"do"}}
	for !p.done() {
		e, err := p.parseExpression()
		if err != nil {
			return nil, err
		}

		statements = append(statements, e)
	}

	return List{statements}, nil
}

func (p *parser) parseExpression() (any, error) {
	node, err := p.parseValue()
	if err != nil {
		return nil, err
	}

	if p.done() || p.peek().Value != "." {
		return node, nil
	}

	chain := []any{Symbol{"get"}, node}

	for !p.done() && p.peek().Value == "." {
		p.expectValue(".")

		if key, err := p.parseIdentifier(); err == nil {
			chain = append(chain, key)
			continue
		}
		if key, err := p.parseInteger(); err == nil {
			chain = append(chain, key)
			continue
		}
		if key, err := p.parseString(); err == nil {
			chain = append(chain, key)
			continue
		}

		return nil, fmt.Errorf(`expected accessor but got %v`, p.peek().Type)
	}

	return List{chain}, nil
}

func (p *parser) parseValue() (any, error) {
	if n, err := p.parseList(); err == nil {
		return n, nil
	}
	if n, err := p.parseQuoted(); err == nil {
		return n, nil
	}
	if n, err := p.parseUnquoted(); err == nil {
		return n, nil
	}
	if n, err := p.parseIdentifier(); err == nil {
		return n, nil
	}
	if n, err := p.parseInteger(); err == nil {
		return n, nil
	}
	if n, err := p.parseString(); err == nil {
		return n, nil
	}
	if n, err := p.parseFloat(); err == nil {
		return n, nil
	}

	return nil, fmt.Errorf(`expected value but got "%s"`, p.peek().Value)
}

func (p *parser) parseList() (any, error) {
	if err := p.expectValue(`(`); err != nil {
		return nil, err
	}

	args := []any{}
	for !p.done() && p.peek().Value != ")" {
		arg, err := p.parseExpression()
		if err != nil {
			return nil, err
		}

		args = append(args, arg)
	}

	if err := p.expectValue(`)`); err != nil {
		return nil, err
	}

	return List{args}, nil
}

func (p *parser) parseQuoted() (any, error) {
	if err := p.expectValue("#"); err != nil {
		return nil, err
	}

	inner, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	return Quote{inner}, nil
}

func (p *parser) parseUnquoted() (any, error) {
	if err := p.expectValue("$"); err != nil {
		return nil, err
	}

	inner, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	return List{[]any{Symbol{"unquote"}, inner}}, nil
}

func (p *parser) parseIdentifier() (any, error) {
	t, err := p.expectType(IdentifierToken)
	if err != nil {
		return nil, err
	}

	return Symbol{t.Value}, nil
}

func (p *parser) parseInteger() (any, error) {
	t, err := p.expectType(IntegerToken)
	if err != nil {
		return nil, err
	}

	value, err := strconv.ParseInt(t.Value, 10, 64)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (p *parser) parseFloat() (any, error) {
	t, err := p.expectType(FloatToken)
	if err != nil {
		return nil, err
	}

	value, err := strconv.ParseFloat(t.Value, 64)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (p *parser) parseString() (any, error) {
	t, err := p.expectType(StringToken)
	if err != nil {
		return nil, err
	}

	value := t.Value[1 : len(t.Value)-1]
	value = strings.ReplaceAll(value, `\n`, "\n")
	value = strings.ReplaceAll(value, `\t`, "\t")

	return value, nil
}
