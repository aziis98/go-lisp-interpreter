package lisp

import (
	"fmt"
	"strings"
)

type Apply struct{ Values []any }
type List struct{ Values []any }
type Drill struct{ Chain []any }
type Map struct{ Entries []Apply }
type Symbol struct{ Name string }
type Quote struct{ Value any }
type Unquote struct{ Value any }

func (e Apply) String() string   { return Show(e) }
func (e List) String() string    { return Show(e) }
func (e Drill) String() string   { return Show(e) }
func (e Map) String() string     { return Show(e) }
func (e Symbol) String() string  { return Show(e) }
func (e Quote) String() string   { return Show(e) }
func (e Unquote) String() string { return Show(e) }

func Show(v any) string {
	if v == nil {
		return "()"
	}

	switch v := v.(type) {
	case string:
		return fmt.Sprintf(`%q`, v)

	case Symbol:
		return v.Name
	case Drill:
		sb := &strings.Builder{}

		for i, v := range v.Chain {
			if i > 0 {
				sb.WriteString(".")
			}
			sb.WriteString(Show(v))
		}

		return sb.String()
	case Quote:
		return wrap("#", "", Show(v.Value))
	case Unquote:
		return wrap("$", "", Show(v.Value))
	case Apply:
		return wrap("(", ")", concatShowers(v.Values))
	case List:
		return wrap("[", "]", concatShowers(v.Values))
	case Map:
		return wrap("{", "}", concat(v.Entries, func(v Apply) string {
			return Show(v.Values[0]) + " " + Show(v.Values[1])
		}))

	default:
		return fmt.Sprintf(`%#v`, v)
	}
}

func wrap(before, after, middle string) string {
	return before + middle + after
}

func concat[T any](values []T, mapper func(T) string) string {
	sb := &strings.Builder{}

	for i, v := range values {
		if i > 0 {
			sb.WriteString(" ")
		}
		sb.WriteString(mapper(v))
	}

	return sb.String()
}

func concatShowers(values []any) string {
	return concat(values, func(v any) string { return Show(v) })
}
