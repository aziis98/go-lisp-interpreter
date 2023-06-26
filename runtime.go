package lisp

import (
	"fmt"
	"strings"
)

type List struct{ Values []any }

func (e List) String() string { return e.Show() }

func (e List) Show() string {
	ss := make([]string, len(e.Values))
	for i, v := range e.Values {
		ss[i] = Show(v)
	}

	return fmt.Sprintf(`(%s)`, strings.Join(ss, " "))
}

type Symbol struct{ Name string }

func (e Symbol) String() string { return e.Show() }

func (e Symbol) Show() string {
	return e.Name
}

type Quote struct{ Value any }

func (e Quote) String() string { return e.Show() }

func (e Quote) Show() string {
	return fmt.Sprintf(`#%s`, Show(e.Value))
}

type Shower interface {
	Show() string
}

func Show(v any) string {
	if v == nil {
		return "()"
	}

	switch v := v.(type) {
	case string:
		return fmt.Sprintf(`%q`, v)
	case Shower:
		return v.Show()
	default:
		return fmt.Sprintf(`%#v`, v)
	}
}
