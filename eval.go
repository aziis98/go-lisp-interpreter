package lisp

import (
	"fmt"
	"os"
	"reflect"
)

type TokenType string

type Token struct {
	Type     TokenType
	Value    string
	Location int
}

func ParseExpression(tokens []Token) (any, error) {
	p := &parser{tokens: tokens, cursor: 0}
	return p.parseExpression()
}

type Scope struct {
	Parent   *Scope
	Bindings map[string]any
}

type Interpreter struct {
	Root *Scope
}

func NewInterpreter() *Interpreter {
	return &Interpreter{
		&Scope{
			Bindings: map[string]any{
				"fmt": map[string]any{
					"Print":   fmt.Print,
					"Printf":  fmt.Printf,
					"Println": fmt.Println,
				},
				"os": map[string]any{
					"Exit": os.Exit,
				},

				"int64": func(v any) int64 { return reflect.ValueOf(v).Int() },
				"int32": func(v any) int32 { return int32(reflect.ValueOf(v).Int()) },
				"int16": func(v any) int16 { return int16(reflect.ValueOf(v).Int()) },
				"int8":  func(v any) int8 { return int8(reflect.ValueOf(v).Int()) },
				"int":   func(v any) int { return int(reflect.ValueOf(v).Int()) },

				"quote": func(item any) any { return Quote{item} },
				"list":  func(items ...any) any { return List{items} },
			},
		},
	}
}

func (i *Interpreter) Eval(program string) (any, error) {
	tokens, err := Tokenize(program)
	if err != nil {
		return nil, err
	}

	p := &parser{tokens: tokens, cursor: 0}
	ast, err := p.parseProgram()
	if err != nil {
		return nil, err
	}

	return i.evaluateNode(i.Root, ast)
}

func (interp *Interpreter) evaluateNode(scope *Scope, v any) (r any, err error) {
	defer func() {
		if rec := recover(); rec != nil {
			r = nil
			err = fmt.Errorf(`internal error: %v`, rec)
		}
	}()

	switch v := v.(type) {
	case List:
		if len(v.Values) == 0 { // () is nil
			return nil, nil
		}

		callee := v.Values[0]
		args := v.Values[1:]

		if id, ok := callee.(Symbol); ok {
			switch id.Name {
			case "do": // (do ...)
				var result any
				for _, stmt := range args {
					var err error
					if result, err = interp.evaluateNode(scope, stmt); err != nil {
						return nil, err
					}
				}
				return result, nil
			case "get": // (get <value> <keys...>)
				value, err := interp.evaluateNode(scope, args[0])
				if err != nil {
					return nil, err
				}

				keys := args[1:]
				for _, key := range keys {
					if l, ok := value.(List); ok {
						value = l.Values
					}

					switch k := key.(type) {
					case Symbol:
						val := reflect.ValueOf(value)
						switch val.Kind() {
						case reflect.Struct:
							value = val.FieldByName(k.Name).Interface()
						case reflect.Map:
							value = val.MapIndex(reflect.ValueOf(k.Name)).Interface()
						default:
							return nil, fmt.Errorf(`cannot get "%s" from %v`, k.Name, value)
						}
					case string:
						value = reflect.ValueOf(value).MapIndex(reflect.ValueOf(k)).Interface()
					case int64:
						value = reflect.ValueOf(value).Index(int(k)).Interface()
					default:
						return nil, fmt.Errorf(`cannot get %#v from %v`, key, value)
					}
				}

				return value, nil
			case "+": // (+ ...)
				var total int64
				for _, arg := range args {
					value, err := interp.evaluateNode(scope, arg)
					if err != nil {
						return nil, err
					}

					total += value.(int64)
				}
				return total, nil

			case "set!":
				variable := args[0].(Symbol).Name

				value, err := interp.evaluateNode(scope, args[1])
				if err != nil {
					return nil, err
				}

				scope.Bindings[variable] = value
				return nil, nil
			}
		}

		evaluatedCallee, err := interp.evaluateNode(scope, callee)
		if err != nil {
			return nil, err
		}

		rArgs := make([]reflect.Value, len(args))
		for i, arg := range args {
			evaluatedArg, err := interp.evaluateNode(scope, arg)
			if err != nil {
				return nil, err
			}

			if evaluatedArg == nil {
				typOfFn := reflect.TypeOf(evaluatedCallee)

				actualArgIndex := i
				if typOfFn.IsVariadic() {
					actualArgIndex = typOfFn.NumIn() - 1
					expectedType := typOfFn.In(actualArgIndex).Elem()
					rArgs[i] = reflect.New(expectedType).Elem()
				} else {
					expectedType := typOfFn.In(actualArgIndex)
					rArgs[i] = reflect.New(expectedType).Elem()
				}
			} else {
				rArgs[i] = reflect.ValueOf(arg)
			}
		}

		rOuts := reflect.ValueOf(evaluatedCallee).Call(rArgs)

		outputs := make([]any, len(rOuts))
		for i, rOut := range rOuts {
			outputs[i] = rOut.Interface()
		}

		if len(outputs) == 1 {
			return outputs[0], nil
		}
		return List{outputs}, nil

	case Symbol:
		return scope.Bindings[v.Name], nil

	default:
		return v, nil
	}
}
